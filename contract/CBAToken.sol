pragma solidity >=0.4.22 <0.7.0;
import "../RLPReader.sol";

/**
 * @title Storage
 * @dev Store & retreive value in a variable
 */
contract CBAToken {

    string constant backingBlockchainName = "ETH";
    string constant networkName = "private";
    uint256 totalSupply_;
    mapping(address => uint256) balances;
    // address ibcServerPublicKeyAddress = 0x72ba7d8e73fe8eb666ea66babc8116a41bfb10e2;
    uint issueRequestId = 0;
    uint redeemRequestId = 0;
    address otherContractAddress;
    mapping(bytes32 => bool) claimedTransactions;
    
    struct IssueData {
        address burnContract;   // the contract which has burnt the tokens on the other blockchian
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
    mapping(uint => address) requestOwners;
    
    // For Contract Administrator
    function registerDepositContract(address tokenContract) public {
        require(tokenContract != address(0), "contract address must not be zero address");
        otherContractAddress = tokenContract;
    }
    
    // Issue Reqeust From User
    
    function handleIssue(bytes memory rlpHeader, bytes memory rlpEncodedTx, bytes memory rlpEncodedReceipt, bytes memory rlpMerkleProofTx, bytes memory rlpMerkleProofReceipt, bytes memory path) public returns (uint){
        IssueData memory issueData = parsingIssueTransaction(rlpEncodedTx, rlpEncodedReceipt);
        bytes32 txHash = keccak256(rlpEncodedTx);
        require(claimedTransactions[txHash] == false, "tokens have already been claimed");
        require(otherContractAddress == issueData.otherContractAddress, "burn contract address is not registered");
        require(issueData.claimContract == address(this), "this contract has not been specified as destination token contract");
        require(issueData.isBurnValid == true, "burn transaction was not successful (e.g., require statement was violated)");

        // verify inclusion of burn transaction
        uint txExists = txInclusionVerifier.verifyTransaction(0, rlpHeader, REQUIRED_TX_CONFIRMATIONS, rlpEncodedTx, path, rlpMerkleProofTx);
        require(txExists == 0, "burn transaction does not exist or has not enough confirmations");

        // verify inclusion of receipt
        uint receiptExists = txInclusionVerifier.verifyReceipt(0, rlpHeader, REQUIRED_TX_CONFIRMATIONS, rlpEncodedReceipt, path, rlpMerkleProofReceipt);
        require(receiptExists == 0, "burn receipt does not exist or has not enough confirmations");
        // bytes memory rlpHeader,             // rlp-encoded header of the block containing burn tx along with its receipt
        // bytes memory rlpEncodedTx,          // rlp-encoded burn tx
        // bytes memory rlpEncodedReceipt,     // rlp-encoded receipt of burn tx ('burn receipt)
        // bytes memory rlpMerkleProofTx,      // rlp-encoded Merkle proof of Membership for burn tx (later passed to relay)
        // bytes memory rlpMerkleProofReceipt, // rlp-encoded Merkle proof of Membership for burn receipt (later passed to relay)
        // bytes memory path  
        // issueLogs.push(IssueLog(msg.sender, msg.sender, txData, txProof, 3));
        // TODO Tx Hash identifier
        // Verify Event
        balances[msg.sender] += 100000;
    }
    
    function redeem(address _counterpartAddress, uint _amount) public payable{
        // require(msg.sender == ibcServerPublicKeyAddress);
        require(balances[msg.sender] >= _amount);
        burn(msg.sender, _amount);
        emit RedeemRequestEvent(redeemRequestId, msg.sender, _counterpartAddress, _amount);
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
        IssueData memory c;
        // parse transaction
        RLPReader.RLPItem[] memory transaction = rlpTransaction.toRlpItem().toList();
        c.burnContract = transaction[3].toAddress();

        // parse receipt
        RLPReader.RLPItem[] memory receipt = rlpReceipt.toRlpItem().toList();
        c.isBurnValid = receipt[0].toBoolean();

        // read logs
        RLPReader.RLPItem[] memory logs = receipt[3].toList();
        RLPReader.RLPItem[] memory burnEventTuple = logs[1].toList();  // logs[0] contains the transfer event emitted by the ECR20 method _burn
                                                                       // logs[1] contains the burn event emitted by the method burn (this contract)
        RLPReader.RLPItem[] memory burnEventTopics = burnEventTuple[1].toList();  // topics contain all indexed event fields

        // read value and recipient from burn event
        c.recipient = address(burnEventTopics[1].toUint());  // indices of indexed fields start at 1 (0 is reserved for the hash of the event signature)
        c.claimContract = address(burnEventTopics[2].toUint());
        c.value = burnEventTopics[3].toUint();

        return c;
    }

    // event Transfer(address indexed from, address indexed to, uint tokens);
    event RedeemRequestEvent(uint redeemRequestId, address redeemerAddress, address counterpartAddress, uint amount);

}


contract Relay {
    
    
}

