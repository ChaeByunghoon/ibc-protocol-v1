pragma solidity >=0.4.22 <0.7.0;

/**
 * @title Storage
 * @dev Store & retreive value in a variable
 */
contract Deposit {

    string constant public issuingBlockchainName = "ETH";
    string constant public networkName = "private";
    uint public constant gwei = 1000000000;
    uint public lockedBalances = 0;
    mapping(address => uint256) public lockedBalancesHistory;
    // address ibcServerPublicKeyAddress = 0x72ba7d8e73fe8eb666ea66babc8116a41bfb10e2;
    address public otherContractAddress;
    
    struct IssueRequest{
        uint issueRequestId;
        address otherContractAddress;
        address issuerAddress;
        address counterpartAddress;
        uint amount;
    }
    
    uint256 issueRequestId = 0;
    IssueRequest[] public issueRequests;
    
    function registerCBAContract(address tokenContract) public {
        require(tokenContract != address(0), "contract address must not be zero address");
        otherContractAddress = tokenContract;
    }
    
   
    function issue(address _issueAddress) public payable {
        require(msg.value > 1 * gwei);
        
        issueRequests.push(IssueRequest(issueRequestId, otherContractAddress, msg.sender, _issueAddress, msg.value));
        issueRequestId += 1;
        lockedBalances += msg.value;
        lockedBalancesHistory[msg.sender] = msg.value;
        emit IssueRequestEvent(otherContractAddress, issueRequestId, msg.sender, _issueAddress, msg.value);
    }
    
    function totalLockedBalance() public view returns (uint){
        return lockedBalances;
    }
    
    // IBC -Server call 
    function handleRedeem(address payable sample, bytes memory txData, bytes memory txProof) public payable{
        // tx data includes [contract address, redeemerAddress, ]
        _transfer(1 * gwei, sample);
    }
    
    // Call at handle Redeem
    function _transfer(uint _amount, address payable redeemerAddress) public payable{
        lockedBalances -= _amount;
        // 
        redeemerAddress.transfer(_amount);
    }


    event IssueRequestEvent(address otherContractAddress, uint issueRequestId, address issuerAddress, address counterpartAddress, uint amount);
    event RedeemRequestEvent(uint redeemRequestId, address issuer, uint amount, string backingBlockchainAddress);

}

