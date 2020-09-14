pragma solidity >=0.4.22 <0.7.0;
import "../RLPReader.sol";
import "../MerklePatriciaProof.sol";

/**
 * @title Storage
 * @dev Store & retreive value in a variable
 */
contract Deposit {

    string constant issuingBlockchainName = "ETH";
    string constant networkName = "private";
    uint constant gwei = 1000000000;
    uint public lockedBalances = 0;
    mapping(address => uint256) public lockedBalancesHistory;
    // address ibcServerPublicKeyAddress = 0x72ba7d8e73fe8eb666ea66babc8116a41bfb10e2;
    mapping(address => bool) public participatingIssuingContract;
    mapping(string => address) public issuingContractAddresses;
    mapping(bytes32 => bool) public redeemedTransactions;

    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;
    using RLPReader for bytes;

    constructor(address _depositContractAddress) public {
        ibcServerPublicKeyAddress = msg.sender;
        depositContractAddress = _depositContractAddress;
    }

    struct RedeemData {
        address issueContractAddress;   // the contract which has burnt the tokens on the other blockchian
        address payable redeemerAddress;
        address claimContractAddress;
        uint value;        // the value to create on this chain
        bool isBurnValid;    // indicates whether the burning of tokens has taken place (didn't abort, e.g., due to require statement)
    }

    struct IssueRequest{
        uint issueRequestId;
        address otherContractAddress;
        address issuerAddress;
        address counterpartAddress;
        uint amount;
    }

    uint256 issueRequestId = 0;
    IssueRequest[] public issueRequests;

    function registerCBAContract(string memory _blockchainName, address _issuingContractAddress) public {
        require(_issuingContractAddress != address(0), "contract address must not be zero address");
        issuingContractAddresses[_blockchainName] = _issuingContractAddress;
        participatingIssuingContract[_issuingContractAddress] = true;
    }


    function issue(address _issueAddress, string memory blockchainName) public payable {
        require(msg.value > 1 * gwei, "The value must higher than 1 gwei");
        address addr = issuingContractAddresses[blockchainName];
        require(participatingIssuingContract[addr] == true, "There is no participatingIssuing Contract");

        issueRequests.push(IssueRequest(issueRequestId, issuingContractAddresses[blockchainName], msg.sender, _issueAddress, msg.value));
        issueRequestId += 1;
        lockedBalances += msg.value;
        lockedBalancesHistory[msg.sender] = msg.value;
        emit IssueRequestEvent(issueRequestId, address(this), issuingContractAddresses[blockchainName], msg.sender, _issueAddress, msg.value);
    }

    function totalLockedBalance() public view returns (uint){
        return lockedBalances;
    }

    // IBC -Server call
    function handleRedeem(bytes memory rlpHeader, bytes memory rlpEncodedTx, bytes memory rlpEncodedReceipt, bytes memory rlpMerkleProofTx, bytes memory rlpMerkleProofReceipt, bytes memory path) public payable{
        // tx data includes [contract address, redeemerAddress, ]
        RedeemData memory redeemData = parsingRedeemTransaction(rlpEncodedTx, rlpEncodedReceipt);
        bytes32 txHash = keccak256(rlpEncodedTx);
        // Check if tx is already claimed.
        require(redeemedTransactions[txHash] == false, "The transaction is already submitted");
        // Check submitted tx contract address is equal in otherContractAddress
        require(participatingIssuingContract[redeemData.issueContractAddress] == true, "burn contract address is not registered");
        // Destination Check
        require(redeemData.claimContractAddress == address(this), "Different targetAddress please check the transaction");

        _transfer(redeemData.value, redeemData.redeemerAddress);
        emit RedeemEvent(redeemData.redeemerAddress, redeemData.value);
    }

    // Call at handle Redeem
    function _transfer(uint _amount, address payable redeemerAddress) public payable{
        lockedBalances -= _amount;
        //
        redeemerAddress.transfer(_amount);
    }

    function parsingRedeemTransaction(bytes memory rlpTransaction, bytes memory rlpReceipt) private pure returns (RedeemData memory) {
        RedeemData memory redeemData;
        // parse transaction
        RLPReader.RLPItem[] memory transaction = rlpTransaction.toRlpItem().toList();
        redeemData.issueContractAddress = transaction[3].toAddress();

        // parse receipt
        RLPReader.RLPItem[] memory receipt = rlpReceipt.toRlpItem().toList();

        // read logs
        RLPReader.RLPItem[] memory logs = receipt[4].toList();
        RLPReader.RLPItem[] memory redeemEventTuple = logs[1].toList();  // logs[0] contains the transfer event emitted by the ECR20 method _burn
        // logs[1] contains the burn event emitted by the method burn (this contract)
        RLPReader.RLPItem[] memory redeemEventTopics = redeemEventTuple[1].toList();  // topics contain all indexed event fields

        // read value and recipient from issue event
        //emit RedeemRequestEvent(redeemRequestId, depositContractAddress, address(this), msg.sender, _counterpartAddress, _amount);
        redeemData.claimContractAddress = address(redeemEventTopics[2].toUint());
        redeemData.issueContractAddress = address(redeemEventTopics[3].toUint());
        redeemData.redeemerAddress = address(redeemEventTopics[5].toUint());  // indices of indexed fields start at 1 (0 is reserved for the hash of the event signature)
        redeemData.value = redeemEventTopics[6].toUint();

        return redeemData;
    }

    event IssueRequestEvent(uint issueRequestId, address depositContractAddress, address otherContractAddress, address issuerAddress, address counterpartAddress, uint amount);
    event RedeemEvent(address redeemerAddress, uint amount);

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
            if ( idx == 4 ) return bytes32(it.next().toUint());
            else it.next();

            idx++;
        }

        return 0;
    }

    function getReceiptsRoot(bytes memory rlpHeader) internal pure returns (bytes32) {
        RLPReader.Iterator memory it = rlpHeader.toRlpItem().iterator();
        uint idx;
        while(it.hasNext()) {
            if ( idx == 5 ) return bytes32(it.next().toUint());
            else it.next();

            idx++;
        }

        return 0;
    }
}

