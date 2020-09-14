pragma solidity >=0.4.22 <0.7.0;
import "../RLPReader.sol";
import "../MerklePatriciaProof.sol";

/**
 * @title Storage
 * @dev Store & retreive value in a variable
 */
contract CBAToken {

    string public constant backingBlockchainName = "ETH";
    string public constant networkName = "private";
    uint256 public totalSupply_;
    mapping(address => uint256) public balances;
    address ibcServerPublicKeyAddress;
    uint public issueRequestId = 0;
    uint public redeemRequestId = 0;
    address depositContractAddress;
    mapping(bytes32 => bool) public claimedTransactions;

    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;
    using RLPReader for bytes;

    constructor(address _depositContractAddress) public {
        ibcServerPublicKeyAddress = msg.sender;
        depositContractAddress = _depositContractAddress;
    }

    struct IssueData {
        address depositContractAddress;   // backing deposit contract Address
        address issuerAddress;
        address claimContract;
        uint value;        // the value to create on this chain
        bool isBurnValid;    // indicates whether the burning of tokens has taken place (didn't abort, e.g., due to require statement)
    }

    struct IssueLog{
        address counterpartAddress;
        address issueAddress;
        // TODO Proof
        bytes32 txData;
        bytes32 txProof;
        uint amount;
    }

    struct RedeemRequest{
        uint id;
        uint amount;
        address ownerAddress;
        string backingAddress;
        bool completed;
    }

    IssueLog[] public issueLogs;
    RedeemRequest[] public redeemRequests;
    mapping(uint => address) public requestOwners;

    // For Contract Administrator
    function registerDepositContract(address tokenContract) public {
        require(tokenContract != address(0), "contract address must not be zero address");
        depositContractAddress = tokenContract;
    }


    // Issue Reqeust From User
    // TODO bytes memory rlpHeader, bytes memory rlpMerkleProofTx, bytes memory rlpMerkleProofReceipt, bytes memory path
    function handleIssue(bytes memory rlpHeader, bytes memory rlpEncodedTx, bytes memory rlpEncodedReceipt, bytes memory rlpMerkleProofTx, bytes memory rlpMerkleProofReceipt, bytes memory path) public returns (uint){
        IssueData memory issueData = parsingIssueTransaction(rlpEncodedTx, rlpEncodedReceipt);
        bytes32 txHash = keccak256(rlpEncodedTx);
        // Check if tx is already claimed.
        require(claimedTransactions[txHash] == false, "The transaction is already submitted");
        // Check submitted tx contract address is equal in otherContractAddress
        require(depositContractAddress == issueData.depositContractAddress, "burn contract address is not registered");
        // Destination Check
        require(issueData.claimContract == address(this), "Different targetAddress please check the transaction");

        // verify inclusion of burn transaction
        // uint txExists = txInclusionVerifier.verifyTransaction(0, rlpHeader, REQUIRED_TX_CONFIRMATIONS, rlpEncodedTx, path, rlpMerkleProofTx);
        // require(txExists == 0, "burn transaction does not exist or has not enough confirmations");
        // // verify inclusion of receipt
        // uint receiptExists = txInclusionVerifier.verifyReceipt(0, rlpHeader, REQUIRED_TX_CONFIRMATIONS, rlpEncodedReceipt, path, rlpMerkleProofReceipt);
        claimedTransactions[keccak256(rlpEncodedTx)] = true;
        balances[msg.sender] += issueData.value;

        emit IssueEvent(issueData.issuerAddress, issueData.value);
    }

    function redeem(address _counterpartAddress, uint _amount) public payable{
        // require(msg.sender == ibcServerPublicKeyAddress);
        require(balances[msg.sender] >= _amount);
        burn(msg.sender, _amount);
        emit RedeemRequestEvent(redeemRequestId, depositContractAddress, address(this), msg.sender, _counterpartAddress, _amount);
    }


    function totalSupply() public view returns (uint){
        return totalSupply_;
    }

    function balanceOf(address owner) public view returns (uint balance){
        return balances[owner];
    }

    function transfer(address to, uint value) public returns (bool success){
        require(to != address(0));
        require(value <= balances[msg.sender]);
        balances[msg.sender] = balances[msg.sender] - value;
        balances[to] = balances[to] + value;
        emit TransferEvent(msg.sender, to, value);
        return true;
    }

    function burn(address _targetAddress, uint _amount) private {
        balances[_targetAddress] -= _amount;
        totalSupply_ -= _amount;
    }

    function compareStrings (string memory a, string memory b) public pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))) );
    }

    function parsingIssueTransaction(bytes memory rlpTransaction, bytes memory rlpReceipt) private pure returns (IssueData memory) {
        IssueData memory issueData;
        // parse transaction
        RLPReader.RLPItem[] memory transaction = rlpTransaction.toRlpItem().toList();
        issueData.depositContractAddress = transaction[3].toAddress();

        // parse receipt
        RLPReader.RLPItem[] memory receipt = rlpReceipt.toRlpItem().toList();

        // read logs
        RLPReader.RLPItem[] memory logs = receipt[4].toList();
        RLPReader.RLPItem[] memory issueEventTuple = logs[1].toList();  // logs[0] contains the transfer event emitted by the ECR20 method _burn
        // logs[1] contains the burn event emitted by the method burn (this contract)
        RLPReader.RLPItem[] memory issueEventTopics = issueEventTuple[1].toList();  // topics contain all indexed event fields

        // read value and recipient from issue event
        issueData.depositContractAddress = address(issueEventTopics[2].toUint());
        issueData.claimContract = address(issueEventTopics[3].toUint());
        // counterpartAddress
        issueData.issuerAddress = address(issueEventTopics[5].toUint());  // indices of indexed fields start at 1 (0 is reserved for the hash of the event signature)
        issueData.value = issueEventTopics[6].toUint();

        return issueData;
    }

    event TransferEvent(address indexed from, address indexed to, uint amount);
    event IssueEvent(address indexed issuerAddress, uint amount);
    event RedeemRequestEvent(uint redeemRequestId, address otherContractAddress, address issuingContractAddress,address redeemerAddress, address counterpartAddress, uint amount);

}

contract Relay{

    using RLPReader for *;
    uint8 constant VERIFICATION_TYPE_TX = 1;
    uint8 constant VERIFICATION_TYPE_RECEIPT = 2;

    function verifyTransaction(uint feeInWei, bytes memory rlpHeader, uint8 noOfConfirmations, bytes memory rlpEncodedTx,
        bytes memory path, bytes memory rlpEncodedNodes) payable public returns (uint8) {
        uint8 result = verify(VERIFICATION_TYPE_TX, feeInWei, rlpHeader, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes);
        return result;
    }

    function verifyReceipt(uint feeInWei, bytes memory rlpHeader, uint8 noOfConfirmations, bytes memory rlpEncodedReceipt,
        bytes memory path, bytes memory rlpEncodedNodes) payable public returns (uint8) {
        uint8 result = verify(VERIFICATION_TYPE_RECEIPT, feeInWei, rlpHeader, noOfConfirmations, rlpEncodedReceipt, path, rlpEncodedNodes);
        return result;
    }

    function verify(uint8 verificationType, uint feeInWei, bytes memory rlpHeader, uint8 noOfConfirmations, bytes memory rlpEncodedValue,
        bytes memory path, bytes memory rlpEncodedNodes) private returns (uint8) {

        bytes32 blockHash = keccak256(rlpHeader);
        uint8 result;

        if (verificationType == VERIFICATION_TYPE_TX) {
            result = verifyMerkleProof(blockHash, noOfConfirmations, rlpEncodedValue, path, rlpEncodedNodes, getTxRoot(rlpHeader));
        }
        else if (verificationType == VERIFICATION_TYPE_RECEIPT) {
            result = verifyMerkleProof(blockHash, noOfConfirmations, rlpEncodedValue, path, rlpEncodedNodes, getReceiptsRoot(rlpHeader));
        }
        else {
            revert("Unknown verification type");
        }

        return result;
    }

    function verifyMerkleProof(bytes32 blockHash, uint8 noOfConfirmations, bytes memory rlpEncodedValue,
        bytes memory path, bytes memory rlpEncodedNodes, bytes32 merkleRootHash) internal view returns (uint8) {

        if (MerklePatriciaProof.verify(rlpEncodedValue, path, rlpEncodedNodes, merkleRootHash) > 0) {
            return 1;
        }

        return 0;
    }

    function getTxRoot(bytes memory rlpHeader) internal pure returns (bytes32) {
        RLPReader.Iterator memory it = rlpHeader.toRlpItem().iterator();
        uint idx;
        while(it.hasNext()) {
            if (idx == 4) return bytes32(it.next().toUint());
            else it.next();
            idx++;
        }

        return 0;
    }

    function getReceiptsRoot(bytes memory rlpHeader) internal pure returns (bytes32) {
        RLPReader.Iterator memory it = rlpHeader.toRlpItem().iterator();
        uint idx;
        while(it.hasNext()) {
            if (idx == 5) return bytes32(it.next().toUint());
            else it.next();

            idx++;
        }

        return 0;
    }
}
