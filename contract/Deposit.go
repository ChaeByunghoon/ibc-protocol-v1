// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deposit

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DepositABI is the input ABI used to generate the binding from.
const DepositABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"otherContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"issueRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"counterpartAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IssueRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"backingBlockchainAddress\",\"type\":\"string\"}],\"name\":\"RedeemRequestEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"redeemerAddress\",\"type\":\"address\"}],\"name\":\"_transfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gwei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rlpEncodedTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedReceipt\",\"type\":\"bytes\"}],\"name\":\"handleRedeem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issueAddress\",\"type\":\"address\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"issueRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"issueRequestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"otherContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"issuerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"counterpartAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuingBlockchainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"registerCBAContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLockedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DepositFuncSigs maps the 4-byte function signature to its string representation.
var DepositFuncSigs = map[string]string{
	"1942c3ce": "_transfer(uint256,address)",
	"24de3c0e": "gwei()",
	"b5e964af": "handleRedeem(bytes,bytes)",
	"71e928af": "issue(address)",
	"4f918973": "issueRequests(uint256)",
	"decad4e1": "issuingBlockchainName()",
	"53e4a89a": "lockedBalances()",
	"107bf28c": "networkName()",
	"d65cc099": "registerCBAContract(address)",
	"2d98ac5e": "totalLockedBalance()",
}

// DepositBin is the compiled bytecode used for deploying new contracts.
var DepositBin = "0x608060405260008055600060045534801561001957600080fd5b50610d71806100296000396000f3fe6080604052600436106100915760003560e01c806353e4a89a1161005957806353e4a89a146101ec57806371e928af14610201578063b5e964af14610227578063d65cc09914610354578063decad4e11461038757610091565b8063107bf28c146100965780631942c3ce1461012057806324de3c0e1461014e5780632d98ac5e146101755780634f9189731461018a575b600080fd5b3480156100a257600080fd5b506100ab61039c565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e55781810151838201526020016100cd565b50505050905090810190601f1680156101125780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61014c6004803603604081101561013657600080fd5b50803590602001356001600160a01b03166103bf565b005b34801561015a57600080fd5b50610163610401565b60408051918252519081900360200190f35b34801561018157600080fd5b50610163610409565b34801561019657600080fd5b506101b4600480360360208110156101ad57600080fd5b503561040f565b604080519586526001600160a01b03948516602087015292841685840152921660608401526080830191909152519081900360a00190f35b3480156101f857600080fd5b5061016361045c565b61014c6004803603602081101561021757600080fd5b50356001600160a01b0316610462565b61014c6004803603604081101561023d57600080fd5b81019060208101813564010000000081111561025857600080fd5b82018360208201111561026a57600080fd5b8035906020019184600183028401116401000000008311171561028c57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156102df57600080fd5b8201836020820111156102f157600080fd5b8035906020019184600183028401116401000000008311171561031357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610684945050505050565b34801561036057600080fd5b5061014c6004803603602081101561037757600080fd5b50356001600160a01b03166107a5565b34801561039357600080fd5b506100ab61080c565b604051806040016040528060078152602001667072697661746560c81b81525081565b6000805483900381556040516001600160a01b0383169184156108fc02918591818181858888f193505050501580156103fc573d6000803e3d6000fd5b505050565b633b9aca0081565b60005490565b6005818154811061041c57fe5b6000918252602090912060059091020180546001820154600283015460038401546004909401549294506001600160a01b03918216939082169291169085565b60005481565b633b9aca00341161047257600080fd5b60056040518060a001604052806004548152602001600260009054906101000a90046001600160a01b03166001600160a01b03168152602001336001600160a01b03168152602001836001600160a01b031681526020013481525090806001815401808255809150506001900390600052602060002090600502016000909190919091506000820151816000015560208201518160010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160020160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060608201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060808201518160040155505060016004600082825401925050819055503460008082825401925050819055503460016000336001600160a01b03166001600160a01b03168152602001908152602001600020819055507fd7b307d0d6c9f7b0ef8f05eb6e748c8ccf32995a5d9a634d4b4419fc0136c974600260009054906101000a90046001600160a01b031660045433843460405180866001600160a01b03166001600160a01b03168152602001858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b031681526020018281526020019550505050505060405180910390a150565b61068c610c4b565b610696838361082b565b8351602080860191909120600081815260039092526040909120549192509060ff16156106f45760405162461bcd60e51b8152600401808060200182810382526024815260200180610d186024913960400191505060405180910390fd5b81516002546001600160a01b039081169116146107425760405162461bcd60e51b8152600401808060200182810382526027815260200180610cf16027913960400191505060405180910390fd5b60408201516001600160a01b0316301461078d5760405162461bcd60e51b8152600401808060200182810382526034815260200180610cbd6034913960400191505060405180910390fd5b61079f826060015183602001516103bf565b50505050565b6001600160a01b0381166107ea5760405162461bcd60e51b8152600401808060200182810382526029815260200180610c946029913960400191505060405180910390fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6040518060400160405280600381526020016208aa8960eb1b81525081565b610833610c4b565b61083b610c4b565b606061084e61084986610942565b610968565b905061086d8160038151811061086057fe5b6020026020010151610a39565b6001600160a01b03168252606061088661084986610942565b905060606108a78260048151811061089a57fe5b6020026020010151610968565b905060606108bb8260018151811061089a57fe5b905060606108cf8260018151811061089a57fe5b90506108ee816000815181106108e157fe5b6020026020010151610a59565b6001600160a01b03166040870152805161090f90829060039081106108e157fe5b6001600160a01b03166020870152805161093090829060049081106108e157fe5b60608701525093979650505050505050565b61094a610c79565b5060408051808201909152815181526020828101908201525b919050565b606061097382610ab9565b61097c57600080fd5b600061098783610af3565b90506060816040519080825280602002602001820160405280156109c557816020015b6109b2610c79565b8152602001906001900390816109aa5790505b50905060006109d78560200151610b4f565b60208601510190506000805b84811015610a2e576109f483610bb2565b9150604051806040016040528083815260200184815250848281518110610a1757fe5b6020908102919091010152918101916001016109e3565b509195945050505050565b8051600090601514610a4a57600080fd5b610a5382610a59565b92915050565b805160009015801590610a6e57508151602110155b610a7757600080fd5b6000610a868360200151610b4f565b83516020808601518301805193945091849003929190831015610ab057826020036101000a820491505b50949350505050565b8051600090610aca57506000610963565b6020820151805160001a9060c0821015610ae957600092505050610963565b5060019392505050565b8051600090610b0457506000610963565b60008090506000610b188460200151610b4f565b602085015185519181019250015b80821015610b4657610b3782610bb2565b60019093019290910190610b26565b50909392505050565b8051600090811a6080811015610b69576000915050610963565b60b8811080610b84575060c08110801590610b84575060f881105b15610b93576001915050610963565b60c0811015610ba75760b519019050610963565b60f519019050610963565b80516000908190811a6080811015610bcd5760019150610c44565b60b8811015610be257607e1981019150610c44565b60c0811015610c0f5760b78103600185019450806020036101000a85510460018201810193505050610c44565b60f8811015610c245760be1981019150610c44565b60f78103600185019450806020036101000a855104600182018101935050505b5092915050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b60405180604001604052806000815260200160008152509056fe636f6e74726163742061646472657373206d757374206e6f74206265207a65726f2061646472657373446966666572656e74207461726765744164647265737320706c6561736520636865636b20746865207472616e73616374696f6e6275726e20636f6e74726163742061646472657373206973206e6f742072656769737465726564546865207472616e73616374696f6e20697320616c7265616479207375626d6974746564a2646970667358221220d7bdf5132549764d19a69bb4a7fa16731473dfcdf2b45ac5ab8a4dbc4dd128cb64736f6c63430006000033"

// DeployDeposit deploys a new Ethereum contract, binding an instance of Deposit to it.
func DeployDeposit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Deposit, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DepositBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Deposit{DepositCaller: DepositCaller{contract: contract}, DepositTransactor: DepositTransactor{contract: contract}, DepositFilterer: DepositFilterer{contract: contract}}, nil
}

// Deposit is an auto generated Go binding around an Ethereum contract.
type Deposit struct {
	DepositCaller     // Read-only binding to the contract
	DepositTransactor // Write-only binding to the contract
	DepositFilterer   // Log filterer for contract events
}

// DepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositSession struct {
	Contract     *Deposit          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositCallerSession struct {
	Contract *DepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositTransactorSession struct {
	Contract     *DepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositRaw struct {
	Contract *Deposit // Generic contract binding to access the raw methods on
}

// DepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositCallerRaw struct {
	Contract *DepositCaller // Generic read-only contract binding to access the raw methods on
}

// DepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositTransactorRaw struct {
	Contract *DepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeposit creates a new instance of Deposit, bound to a specific deployed contract.
func NewDeposit(address common.Address, backend bind.ContractBackend) (*Deposit, error) {
	contract, err := bindDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deposit{DepositCaller: DepositCaller{contract: contract}, DepositTransactor: DepositTransactor{contract: contract}, DepositFilterer: DepositFilterer{contract: contract}}, nil
}

// NewDepositCaller creates a new read-only instance of Deposit, bound to a specific deployed contract.
func NewDepositCaller(address common.Address, caller bind.ContractCaller) (*DepositCaller, error) {
	contract, err := bindDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositCaller{contract: contract}, nil
}

// NewDepositTransactor creates a new write-only instance of Deposit, bound to a specific deployed contract.
func NewDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositTransactor, error) {
	contract, err := bindDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositTransactor{contract: contract}, nil
}

// NewDepositFilterer creates a new log filterer instance of Deposit, bound to a specific deployed contract.
func NewDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositFilterer, error) {
	contract, err := bindDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositFilterer{contract: contract}, nil
}

// bindDeposit binds a generic wrapper to an already deployed contract.
func bindDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deposit *DepositRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Deposit.Contract.DepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deposit *DepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.Contract.DepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deposit *DepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deposit.Contract.DepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deposit *DepositCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Deposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deposit *DepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deposit *DepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deposit.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0x1942c3ce.
//
// Solidity: function _transfer(uint256 _amount, address redeemerAddress) returns()
func (_Deposit *DepositTransactor) Transfer(opts *bind.TransactOpts, _amount *big.Int, redeemerAddress common.Address) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "_transfer", _amount, redeemerAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0x1942c3ce.
//
// Solidity: function _transfer(uint256 _amount, address redeemerAddress) returns()
func (_Deposit *DepositSession) Transfer(_amount *big.Int, redeemerAddress common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.Transfer(&_Deposit.TransactOpts, _amount, redeemerAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0x1942c3ce.
//
// Solidity: function _transfer(uint256 _amount, address redeemerAddress) returns()
func (_Deposit *DepositTransactorSession) Transfer(_amount *big.Int, redeemerAddress common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.Transfer(&_Deposit.TransactOpts, _amount, redeemerAddress)
}

// Gwei is a paid mutator transaction binding the contract method 0x24de3c0e.
//
// Solidity: function gwei() returns(uint256)
func (_Deposit *DepositTransactor) Gwei(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "gwei")
}

// Gwei is a paid mutator transaction binding the contract method 0x24de3c0e.
//
// Solidity: function gwei() returns(uint256)
func (_Deposit *DepositSession) Gwei() (*types.Transaction, error) {
	return _Deposit.Contract.Gwei(&_Deposit.TransactOpts)
}

// Gwei is a paid mutator transaction binding the contract method 0x24de3c0e.
//
// Solidity: function gwei() returns(uint256)
func (_Deposit *DepositTransactorSession) Gwei() (*types.Transaction, error) {
	return _Deposit.Contract.Gwei(&_Deposit.TransactOpts)
}

// HandleRedeem is a paid mutator transaction binding the contract method 0xb5e964af.
//
// Solidity: function handleRedeem(bytes rlpEncodedTx, bytes rlpEncodedReceipt) returns()
func (_Deposit *DepositTransactor) HandleRedeem(opts *bind.TransactOpts, rlpEncodedTx []byte, rlpEncodedReceipt []byte) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "handleRedeem", rlpEncodedTx, rlpEncodedReceipt)
}

// HandleRedeem is a paid mutator transaction binding the contract method 0xb5e964af.
//
// Solidity: function handleRedeem(bytes rlpEncodedTx, bytes rlpEncodedReceipt) returns()
func (_Deposit *DepositSession) HandleRedeem(rlpEncodedTx []byte, rlpEncodedReceipt []byte) (*types.Transaction, error) {
	return _Deposit.Contract.HandleRedeem(&_Deposit.TransactOpts, rlpEncodedTx, rlpEncodedReceipt)
}

// HandleRedeem is a paid mutator transaction binding the contract method 0xb5e964af.
//
// Solidity: function handleRedeem(bytes rlpEncodedTx, bytes rlpEncodedReceipt) returns()
func (_Deposit *DepositTransactorSession) HandleRedeem(rlpEncodedTx []byte, rlpEncodedReceipt []byte) (*types.Transaction, error) {
	return _Deposit.Contract.HandleRedeem(&_Deposit.TransactOpts, rlpEncodedTx, rlpEncodedReceipt)
}

// Issue is a paid mutator transaction binding the contract method 0x71e928af.
//
// Solidity: function issue(address _issueAddress) returns()
func (_Deposit *DepositTransactor) Issue(opts *bind.TransactOpts, _issueAddress common.Address) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "issue", _issueAddress)
}

// Issue is a paid mutator transaction binding the contract method 0x71e928af.
//
// Solidity: function issue(address _issueAddress) returns()
func (_Deposit *DepositSession) Issue(_issueAddress common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.Issue(&_Deposit.TransactOpts, _issueAddress)
}

// Issue is a paid mutator transaction binding the contract method 0x71e928af.
//
// Solidity: function issue(address _issueAddress) returns()
func (_Deposit *DepositTransactorSession) Issue(_issueAddress common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.Issue(&_Deposit.TransactOpts, _issueAddress)
}

// IssueRequests is a paid mutator transaction binding the contract method 0x4f918973.
//
// Solidity: function issueRequests(uint256 ) returns(uint256 issueRequestId, address otherContractAddress, address issuerAddress, address counterpartAddress, uint256 amount)
func (_Deposit *DepositTransactor) IssueRequests(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "issueRequests", arg0)
}

// IssueRequests is a paid mutator transaction binding the contract method 0x4f918973.
//
// Solidity: function issueRequests(uint256 ) returns(uint256 issueRequestId, address otherContractAddress, address issuerAddress, address counterpartAddress, uint256 amount)
func (_Deposit *DepositSession) IssueRequests(arg0 *big.Int) (*types.Transaction, error) {
	return _Deposit.Contract.IssueRequests(&_Deposit.TransactOpts, arg0)
}

// IssueRequests is a paid mutator transaction binding the contract method 0x4f918973.
//
// Solidity: function issueRequests(uint256 ) returns(uint256 issueRequestId, address otherContractAddress, address issuerAddress, address counterpartAddress, uint256 amount)
func (_Deposit *DepositTransactorSession) IssueRequests(arg0 *big.Int) (*types.Transaction, error) {
	return _Deposit.Contract.IssueRequests(&_Deposit.TransactOpts, arg0)
}

// IssuingBlockchainName is a paid mutator transaction binding the contract method 0xdecad4e1.
//
// Solidity: function issuingBlockchainName() returns(string)
func (_Deposit *DepositTransactor) IssuingBlockchainName(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "issuingBlockchainName")
}

// IssuingBlockchainName is a paid mutator transaction binding the contract method 0xdecad4e1.
//
// Solidity: function issuingBlockchainName() returns(string)
func (_Deposit *DepositSession) IssuingBlockchainName() (*types.Transaction, error) {
	return _Deposit.Contract.IssuingBlockchainName(&_Deposit.TransactOpts)
}

// IssuingBlockchainName is a paid mutator transaction binding the contract method 0xdecad4e1.
//
// Solidity: function issuingBlockchainName() returns(string)
func (_Deposit *DepositTransactorSession) IssuingBlockchainName() (*types.Transaction, error) {
	return _Deposit.Contract.IssuingBlockchainName(&_Deposit.TransactOpts)
}

// LockedBalances is a paid mutator transaction binding the contract method 0x53e4a89a.
//
// Solidity: function lockedBalances() returns(uint256)
func (_Deposit *DepositTransactor) LockedBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "lockedBalances")
}

// LockedBalances is a paid mutator transaction binding the contract method 0x53e4a89a.
//
// Solidity: function lockedBalances() returns(uint256)
func (_Deposit *DepositSession) LockedBalances() (*types.Transaction, error) {
	return _Deposit.Contract.LockedBalances(&_Deposit.TransactOpts)
}

// LockedBalances is a paid mutator transaction binding the contract method 0x53e4a89a.
//
// Solidity: function lockedBalances() returns(uint256)
func (_Deposit *DepositTransactorSession) LockedBalances() (*types.Transaction, error) {
	return _Deposit.Contract.LockedBalances(&_Deposit.TransactOpts)
}

// NetworkName is a paid mutator transaction binding the contract method 0x107bf28c.
//
// Solidity: function networkName() returns(string)
func (_Deposit *DepositTransactor) NetworkName(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "networkName")
}

// NetworkName is a paid mutator transaction binding the contract method 0x107bf28c.
//
// Solidity: function networkName() returns(string)
func (_Deposit *DepositSession) NetworkName() (*types.Transaction, error) {
	return _Deposit.Contract.NetworkName(&_Deposit.TransactOpts)
}

// NetworkName is a paid mutator transaction binding the contract method 0x107bf28c.
//
// Solidity: function networkName() returns(string)
func (_Deposit *DepositTransactorSession) NetworkName() (*types.Transaction, error) {
	return _Deposit.Contract.NetworkName(&_Deposit.TransactOpts)
}

// RegisterCBAContract is a paid mutator transaction binding the contract method 0xd65cc099.
//
// Solidity: function registerCBAContract(address tokenContract) returns()
func (_Deposit *DepositTransactor) RegisterCBAContract(opts *bind.TransactOpts, tokenContract common.Address) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "registerCBAContract", tokenContract)
}

// RegisterCBAContract is a paid mutator transaction binding the contract method 0xd65cc099.
//
// Solidity: function registerCBAContract(address tokenContract) returns()
func (_Deposit *DepositSession) RegisterCBAContract(tokenContract common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.RegisterCBAContract(&_Deposit.TransactOpts, tokenContract)
}

// RegisterCBAContract is a paid mutator transaction binding the contract method 0xd65cc099.
//
// Solidity: function registerCBAContract(address tokenContract) returns()
func (_Deposit *DepositTransactorSession) RegisterCBAContract(tokenContract common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.RegisterCBAContract(&_Deposit.TransactOpts, tokenContract)
}

// TotalLockedBalance is a paid mutator transaction binding the contract method 0x2d98ac5e.
//
// Solidity: function totalLockedBalance() returns(uint256)
func (_Deposit *DepositTransactor) TotalLockedBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "totalLockedBalance")
}

// TotalLockedBalance is a paid mutator transaction binding the contract method 0x2d98ac5e.
//
// Solidity: function totalLockedBalance() returns(uint256)
func (_Deposit *DepositSession) TotalLockedBalance() (*types.Transaction, error) {
	return _Deposit.Contract.TotalLockedBalance(&_Deposit.TransactOpts)
}

// TotalLockedBalance is a paid mutator transaction binding the contract method 0x2d98ac5e.
//
// Solidity: function totalLockedBalance() returns(uint256)
func (_Deposit *DepositTransactorSession) TotalLockedBalance() (*types.Transaction, error) {
	return _Deposit.Contract.TotalLockedBalance(&_Deposit.TransactOpts)
}

// DepositIssueRequestEventIterator is returned from FilterIssueRequestEvent and is used to iterate over the raw logs and unpacked data for IssueRequestEvent events raised by the Deposit contract.
type DepositIssueRequestEventIterator struct {
	Event *DepositIssueRequestEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DepositIssueRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositIssueRequestEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DepositIssueRequestEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DepositIssueRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositIssueRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositIssueRequestEvent represents a IssueRequestEvent event raised by the Deposit contract.
type DepositIssueRequestEvent struct {
	OtherContractAddress common.Address
	IssueRequestId       *big.Int
	IssuerAddress        common.Address
	CounterpartAddress   common.Address
	Amount               *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterIssueRequestEvent is a free log retrieval operation binding the contract event 0xd7b307d0d6c9f7b0ef8f05eb6e748c8ccf32995a5d9a634d4b4419fc0136c974.
//
// Solidity: event IssueRequestEvent(address otherContractAddress, uint256 issueRequestId, address issuerAddress, address counterpartAddress, uint256 amount)
func (_Deposit *DepositFilterer) FilterIssueRequestEvent(opts *bind.FilterOpts) (*DepositIssueRequestEventIterator, error) {

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "IssueRequestEvent")
	if err != nil {
		return nil, err
	}
	return &DepositIssueRequestEventIterator{contract: _Deposit.contract, event: "IssueRequestEvent", logs: logs, sub: sub}, nil
}

// WatchIssueRequestEvent is a free log subscription operation binding the contract event 0xd7b307d0d6c9f7b0ef8f05eb6e748c8ccf32995a5d9a634d4b4419fc0136c974.
//
// Solidity: event IssueRequestEvent(address otherContractAddress, uint256 issueRequestId, address issuerAddress, address counterpartAddress, uint256 amount)
func (_Deposit *DepositFilterer) WatchIssueRequestEvent(opts *bind.WatchOpts, sink chan<- *DepositIssueRequestEvent) (event.Subscription, error) {

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "IssueRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositIssueRequestEvent)
				if err := _Deposit.contract.UnpackLog(event, "IssueRequestEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIssueRequestEvent is a log parse operation binding the contract event 0xd7b307d0d6c9f7b0ef8f05eb6e748c8ccf32995a5d9a634d4b4419fc0136c974.
//
// Solidity: event IssueRequestEvent(address otherContractAddress, uint256 issueRequestId, address issuerAddress, address counterpartAddress, uint256 amount)
func (_Deposit *DepositFilterer) ParseIssueRequestEvent(log types.Log) (*DepositIssueRequestEvent, error) {
	event := new(DepositIssueRequestEvent)
	if err := _Deposit.contract.UnpackLog(event, "IssueRequestEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositRedeemRequestEventIterator is returned from FilterRedeemRequestEvent and is used to iterate over the raw logs and unpacked data for RedeemRequestEvent events raised by the Deposit contract.
type DepositRedeemRequestEventIterator struct {
	Event *DepositRedeemRequestEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DepositRedeemRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositRedeemRequestEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DepositRedeemRequestEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DepositRedeemRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositRedeemRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositRedeemRequestEvent represents a RedeemRequestEvent event raised by the Deposit contract.
type DepositRedeemRequestEvent struct {
	RedeemRequestId          *big.Int
	Issuer                   common.Address
	Amount                   *big.Int
	BackingBlockchainAddress string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRedeemRequestEvent is a free log retrieval operation binding the contract event 0x6d2ba3b4530992ccf459c61579904630daf5a9d4c1c343bac0848b73c38a9aa9.
//
// Solidity: event RedeemRequestEvent(uint256 redeemRequestId, address issuer, uint256 amount, string backingBlockchainAddress)
func (_Deposit *DepositFilterer) FilterRedeemRequestEvent(opts *bind.FilterOpts) (*DepositRedeemRequestEventIterator, error) {

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "RedeemRequestEvent")
	if err != nil {
		return nil, err
	}
	return &DepositRedeemRequestEventIterator{contract: _Deposit.contract, event: "RedeemRequestEvent", logs: logs, sub: sub}, nil
}

// WatchRedeemRequestEvent is a free log subscription operation binding the contract event 0x6d2ba3b4530992ccf459c61579904630daf5a9d4c1c343bac0848b73c38a9aa9.
//
// Solidity: event RedeemRequestEvent(uint256 redeemRequestId, address issuer, uint256 amount, string backingBlockchainAddress)
func (_Deposit *DepositFilterer) WatchRedeemRequestEvent(opts *bind.WatchOpts, sink chan<- *DepositRedeemRequestEvent) (event.Subscription, error) {

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "RedeemRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositRedeemRequestEvent)
				if err := _Deposit.contract.UnpackLog(event, "RedeemRequestEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemRequestEvent is a log parse operation binding the contract event 0x6d2ba3b4530992ccf459c61579904630daf5a9d4c1c343bac0848b73c38a9aa9.
//
// Solidity: event RedeemRequestEvent(uint256 redeemRequestId, address issuer, uint256 amount, string backingBlockchainAddress)
func (_Deposit *DepositFilterer) ParseRedeemRequestEvent(log types.Log) (*DepositRedeemRequestEvent, error) {
	event := new(DepositRedeemRequestEvent)
	if err := _Deposit.contract.UnpackLog(event, "RedeemRequestEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RLPReaderABI is the input ABI used to generate the binding from.
const RLPReaderABI = "[]"

// RLPReaderBin is the compiled bytecode used for deploying new contracts.
var RLPReaderBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204aec344eb323fef04a6bd2e2e699b3c3e045a887b5bc43dc9064bb2a780a31a764736f6c63430006000033"

// DeployRLPReader deploys a new Ethereum contract, binding an instance of RLPReader to it.
func DeployRLPReader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPReader, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPReaderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPReaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPReader{RLPReaderCaller: RLPReaderCaller{contract: contract}, RLPReaderTransactor: RLPReaderTransactor{contract: contract}, RLPReaderFilterer: RLPReaderFilterer{contract: contract}}, nil
}

// RLPReader is an auto generated Go binding around an Ethereum contract.
type RLPReader struct {
	RLPReaderCaller     // Read-only binding to the contract
	RLPReaderTransactor // Write-only binding to the contract
	RLPReaderFilterer   // Log filterer for contract events
}

// RLPReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPReaderSession struct {
	Contract     *RLPReader        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPReaderCallerSession struct {
	Contract *RLPReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPReaderTransactorSession struct {
	Contract     *RLPReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPReaderRaw struct {
	Contract *RLPReader // Generic contract binding to access the raw methods on
}

// RLPReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPReaderCallerRaw struct {
	Contract *RLPReaderCaller // Generic read-only contract binding to access the raw methods on
}

// RLPReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPReaderTransactorRaw struct {
	Contract *RLPReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPReader creates a new instance of RLPReader, bound to a specific deployed contract.
func NewRLPReader(address common.Address, backend bind.ContractBackend) (*RLPReader, error) {
	contract, err := bindRLPReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPReader{RLPReaderCaller: RLPReaderCaller{contract: contract}, RLPReaderTransactor: RLPReaderTransactor{contract: contract}, RLPReaderFilterer: RLPReaderFilterer{contract: contract}}, nil
}

// NewRLPReaderCaller creates a new read-only instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderCaller(address common.Address, caller bind.ContractCaller) (*RLPReaderCaller, error) {
	contract, err := bindRLPReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPReaderCaller{contract: contract}, nil
}

// NewRLPReaderTransactor creates a new write-only instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPReaderTransactor, error) {
	contract, err := bindRLPReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPReaderTransactor{contract: contract}, nil
}

// NewRLPReaderFilterer creates a new log filterer instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPReaderFilterer, error) {
	contract, err := bindRLPReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPReaderFilterer{contract: contract}, nil
}

// bindRLPReader binds a generic wrapper to an already deployed contract.
func bindRLPReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPReaderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPReader *RLPReaderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLPReader.Contract.RLPReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPReader *RLPReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPReader.Contract.RLPReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPReader *RLPReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPReader.Contract.RLPReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPReader *RLPReaderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLPReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPReader *RLPReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPReader *RLPReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPReader.Contract.contract.Transact(opts, method, params...)
}
