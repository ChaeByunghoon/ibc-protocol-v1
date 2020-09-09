pragma solidity >=0.4.22 <0.7.0;
import "../RLPReader.sol";

/**
 * @title Storage
 * @dev Store & retreive value in a variable
 */
contract CBAToken {

    string public constant backingBlockchainName = "ETH";
    string public constant networkName = "private";
    uint256 public totalSupply_;
    mapping(address => uint256) public balances;
    // address ibcServerPublicKeyAddress = 0x72ba7d8e73fe8eb666ea66babc8116a41bfb10e2;
    uint public issueRequestId = 0;
    uint public redeemRequestId = 0;
    address otherContractAddress;
    mapping(bytes32 => bool) public claimedTransactions;
    
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;
    using RLPReader for bytes;
    
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;
    using RLPReader for bytes;
    
    struct IssueData {
        address depositContractAddress;   // the contract which has burnt the tokens on the other blockchian
        address recipient;
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
        otherContractAddress = tokenContract;
    }
    
    // Issue Reqeust From User
    // TODO bytes memory rlpHeader, bytes memory rlpMerkleProofTx, bytes memory rlpMerkleProofReceipt, bytes memory path
    function handleIssue(bytes memory rlpEncodedTx, bytes memory rlpEncodedReceipt) public returns (uint){
        IssueData memory issueData = parsingIssueTransaction(rlpEncodedTx, rlpEncodedReceipt);
        bytes32 txHash = keccak256(rlpEncodedTx);
        // Check if tx is already claimed.
        require(claimedTransactions[txHash] == false, "The transaction is already submitted");
        // Check submitted tx contract address is equal in otherContractAddress
        require(otherContractAddress == issueData.depositContractAddress, "burn contract address is not registered");
        // Destination Check
        require(issueData.claimContract == address(this), "Different targetAddress please check the transaction");

        // verify inclusion of burn transaction
        // uint txExists = txInclusionVerifier.verifyTransaction(0, rlpHeader, REQUIRED_TX_CONFIRMATIONS, rlpEncodedTx, path, rlpMerkleProofTx);
        // require(txExists == 0, "burn transaction does not exist or has not enough confirmations");
        // // verify inclusion of receipt
        // uint receiptExists = txInclusionVerifier.verifyReceipt(0, rlpHeader, REQUIRED_TX_CONFIRMATIONS, rlpEncodedReceipt, path, rlpMerkleProofReceipt);
        claimedTransactions[keccak256(rlpEncodedTx)] = true;
        balances[msg.sender] += issueData.value;
    }
    
    function redeem(address _counterpartAddress, uint _amount) public payable{
        // require(msg.sender == ibcServerPublicKeyAddress);
        require(balances[msg.sender] >= _amount);
        burn(msg.sender, _amount);
        emit RedeemRequestEvent(redeemRequestId, msg.sender, _counterpartAddress, _amount);
    }
    
    function balanceOf(address owner) public view returns (uint balance){
        return balances[owner];
    }
    
    function transfer(address to, uint value) public returns (bool success){
        require(to != address(0));
        require(value <= balances[msg.sender]);
        balances[msg.sender] = balances[msg.sender] - value;
        balances[to] = balances[to] + value;
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
        issueData.claimContract = address(issueEventTopics[0].toUint());
        issueData.recipient = address(issueEventTopics[3].toUint());  // indices of indexed fields start at 1 (0 is reserved for the hash of the event signature)
        issueData.value = issueEventTopics[4].toUint();
        
        return issueData;
    }

    // event Transfer(address indexed from, address indexed to, uint tokens);
    event RedeemRequestEvent(uint redeemRequestId, address redeemerAddress, address counterpartAddress, uint amount);

}
