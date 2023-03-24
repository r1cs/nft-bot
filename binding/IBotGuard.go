// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BotGuardMetaData contains all meta data concerning the BotGuard contract.
var BotGuardMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"call894739472894\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mintContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_expectedTotalSupply\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_maxNum\",\"type\":\"uint64\"}],\"name\":\"expectMintTotalSupply\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_targetDomain\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_faucetToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"staticCall897984374\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BotGuardABI is the input ABI used to generate the binding from.
// Deprecated: Use BotGuardMetaData.ABI instead.
var BotGuardABI = BotGuardMetaData.ABI

// BotGuard is an auto generated Go binding around an Ethereum contract.
type BotGuard struct {
	BotGuardCaller     // Read-only binding to the contract
	BotGuardTransactor // Write-only binding to the contract
	BotGuardFilterer   // Log filterer for contract events
}

// BotGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type BotGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BotGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BotGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BotGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BotGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BotGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BotGuardSession struct {
	Contract     *BotGuard         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BotGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BotGuardCallerSession struct {
	Contract *BotGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BotGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BotGuardTransactorSession struct {
	Contract     *BotGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BotGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type BotGuardRaw struct {
	Contract *BotGuard // Generic contract binding to access the raw methods on
}

// BotGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BotGuardCallerRaw struct {
	Contract *BotGuardCaller // Generic read-only contract binding to access the raw methods on
}

// BotGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BotGuardTransactorRaw struct {
	Contract *BotGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBotGuard creates a new instance of BotGuard, bound to a specific deployed contract.
func NewBotGuard(address common.Address, backend bind.ContractBackend) (*BotGuard, error) {
	contract, err := bindBotGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BotGuard{BotGuardCaller: BotGuardCaller{contract: contract}, BotGuardTransactor: BotGuardTransactor{contract: contract}, BotGuardFilterer: BotGuardFilterer{contract: contract}}, nil
}

// NewBotGuardCaller creates a new read-only instance of BotGuard, bound to a specific deployed contract.
func NewBotGuardCaller(address common.Address, caller bind.ContractCaller) (*BotGuardCaller, error) {
	contract, err := bindBotGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BotGuardCaller{contract: contract}, nil
}

// NewBotGuardTransactor creates a new write-only instance of BotGuard, bound to a specific deployed contract.
func NewBotGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*BotGuardTransactor, error) {
	contract, err := bindBotGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BotGuardTransactor{contract: contract}, nil
}

// NewBotGuardFilterer creates a new log filterer instance of BotGuard, bound to a specific deployed contract.
func NewBotGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*BotGuardFilterer, error) {
	contract, err := bindBotGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BotGuardFilterer{contract: contract}, nil
}

// bindBotGuard binds a generic wrapper to an already deployed contract.
func bindBotGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BotGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BotGuard *BotGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BotGuard.Contract.BotGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BotGuard *BotGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BotGuard.Contract.BotGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BotGuard *BotGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BotGuard.Contract.BotGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BotGuard *BotGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BotGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BotGuard *BotGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BotGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BotGuard *BotGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BotGuard.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BotGuard *BotGuardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BotGuard.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BotGuard *BotGuardSession) Owner() (common.Address, error) {
	return _BotGuard.Contract.Owner(&_BotGuard.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BotGuard *BotGuardCallerSession) Owner() (common.Address, error) {
	return _BotGuard.Contract.Owner(&_BotGuard.CallOpts)
}

// StaticCall897984374 is a free data retrieval call binding the contract method 0xac1f49ea.
//
// Solidity: function staticCall897984374(address _target, bytes _params) view returns(bytes)
func (_BotGuard *BotGuardCaller) StaticCall897984374(opts *bind.CallOpts, _target common.Address, _params []byte) ([]byte, error) {
	var out []interface{}
	err := _BotGuard.contract.Call(opts, &out, "staticCall897984374", _target, _params)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// StaticCall897984374 is a free data retrieval call binding the contract method 0xac1f49ea.
//
// Solidity: function staticCall897984374(address _target, bytes _params) view returns(bytes)
func (_BotGuard *BotGuardSession) StaticCall897984374(_target common.Address, _params []byte) ([]byte, error) {
	return _BotGuard.Contract.StaticCall897984374(&_BotGuard.CallOpts, _target, _params)
}

// StaticCall897984374 is a free data retrieval call binding the contract method 0xac1f49ea.
//
// Solidity: function staticCall897984374(address _target, bytes _params) view returns(bytes)
func (_BotGuard *BotGuardCallerSession) StaticCall897984374(_target common.Address, _params []byte) ([]byte, error) {
	return _BotGuard.Contract.StaticCall897984374(&_BotGuard.CallOpts, _target, _params)
}

// Call894739472894 is a paid mutator transaction binding the contract method 0x6ab2a3b9.
//
// Solidity: function call894739472894(address _target, bytes params) payable returns()
func (_BotGuard *BotGuardTransactor) Call894739472894(opts *bind.TransactOpts, _target common.Address, params []byte) (*types.Transaction, error) {
	return _BotGuard.contract.Transact(opts, "call894739472894", _target, params)
}

// Call894739472894 is a paid mutator transaction binding the contract method 0x6ab2a3b9.
//
// Solidity: function call894739472894(address _target, bytes params) payable returns()
func (_BotGuard *BotGuardSession) Call894739472894(_target common.Address, params []byte) (*types.Transaction, error) {
	return _BotGuard.Contract.Call894739472894(&_BotGuard.TransactOpts, _target, params)
}

// Call894739472894 is a paid mutator transaction binding the contract method 0x6ab2a3b9.
//
// Solidity: function call894739472894(address _target, bytes params) payable returns()
func (_BotGuard *BotGuardTransactorSession) Call894739472894(_target common.Address, params []byte) (*types.Transaction, error) {
	return _BotGuard.Contract.Call894739472894(&_BotGuard.TransactOpts, _target, params)
}

// ExpectMintTotalSupply is a paid mutator transaction binding the contract method 0x68d4889e.
//
// Solidity: function expectMintTotalSupply(address _mintContract, bytes _params, address _nftContract, uint64 _expectedTotalSupply, uint256 _mintPrice, uint64 _maxNum) payable returns()
func (_BotGuard *BotGuardTransactor) ExpectMintTotalSupply(opts *bind.TransactOpts, _mintContract common.Address, _params []byte, _nftContract common.Address, _expectedTotalSupply uint64, _mintPrice *big.Int, _maxNum uint64) (*types.Transaction, error) {
	return _BotGuard.contract.Transact(opts, "expectMintTotalSupply", _mintContract, _params, _nftContract, _expectedTotalSupply, _mintPrice, _maxNum)
}

// ExpectMintTotalSupply is a paid mutator transaction binding the contract method 0x68d4889e.
//
// Solidity: function expectMintTotalSupply(address _mintContract, bytes _params, address _nftContract, uint64 _expectedTotalSupply, uint256 _mintPrice, uint64 _maxNum) payable returns()
func (_BotGuard *BotGuardSession) ExpectMintTotalSupply(_mintContract common.Address, _params []byte, _nftContract common.Address, _expectedTotalSupply uint64, _mintPrice *big.Int, _maxNum uint64) (*types.Transaction, error) {
	return _BotGuard.Contract.ExpectMintTotalSupply(&_BotGuard.TransactOpts, _mintContract, _params, _nftContract, _expectedTotalSupply, _mintPrice, _maxNum)
}

// ExpectMintTotalSupply is a paid mutator transaction binding the contract method 0x68d4889e.
//
// Solidity: function expectMintTotalSupply(address _mintContract, bytes _params, address _nftContract, uint64 _expectedTotalSupply, uint256 _mintPrice, uint64 _maxNum) payable returns()
func (_BotGuard *BotGuardTransactorSession) ExpectMintTotalSupply(_mintContract common.Address, _params []byte, _nftContract common.Address, _expectedTotalSupply uint64, _mintPrice *big.Int, _maxNum uint64) (*types.Transaction, error) {
	return _BotGuard.Contract.ExpectMintTotalSupply(&_BotGuard.TransactOpts, _mintContract, _params, _nftContract, _expectedTotalSupply, _mintPrice, _maxNum)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _domain, string _targetDomain, address _faucetToken) returns()
func (_BotGuard *BotGuardTransactor) Initialize(opts *bind.TransactOpts, _domain string, _targetDomain string, _faucetToken common.Address) (*types.Transaction, error) {
	return _BotGuard.contract.Transact(opts, "initialize", _domain, _targetDomain, _faucetToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _domain, string _targetDomain, address _faucetToken) returns()
func (_BotGuard *BotGuardSession) Initialize(_domain string, _targetDomain string, _faucetToken common.Address) (*types.Transaction, error) {
	return _BotGuard.Contract.Initialize(&_BotGuard.TransactOpts, _domain, _targetDomain, _faucetToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _domain, string _targetDomain, address _faucetToken) returns()
func (_BotGuard *BotGuardTransactorSession) Initialize(_domain string, _targetDomain string, _faucetToken common.Address) (*types.Transaction, error) {
	return _BotGuard.Contract.Initialize(&_BotGuard.TransactOpts, _domain, _targetDomain, _faucetToken)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BotGuard *BotGuardTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _BotGuard.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BotGuard *BotGuardSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _BotGuard.Contract.Multicall(&_BotGuard.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BotGuard *BotGuardTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _BotGuard.Contract.Multicall(&_BotGuard.TransactOpts, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BotGuard *BotGuardTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BotGuard.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BotGuard *BotGuardSession) RenounceOwnership() (*types.Transaction, error) {
	return _BotGuard.Contract.RenounceOwnership(&_BotGuard.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BotGuard *BotGuardTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BotGuard.Contract.RenounceOwnership(&_BotGuard.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BotGuard *BotGuardTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BotGuard.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BotGuard *BotGuardSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BotGuard.Contract.TransferOwnership(&_BotGuard.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BotGuard *BotGuardTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BotGuard.Contract.TransferOwnership(&_BotGuard.TransactOpts, newOwner)
}

// BotGuardInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BotGuard contract.
type BotGuardInitializedIterator struct {
	Event *BotGuardInitialized // Event containing the contract specifics and raw log

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
func (it *BotGuardInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BotGuardInitialized)
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
		it.Event = new(BotGuardInitialized)
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
func (it *BotGuardInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BotGuardInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BotGuardInitialized represents a Initialized event raised by the BotGuard contract.
type BotGuardInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BotGuard *BotGuardFilterer) FilterInitialized(opts *bind.FilterOpts) (*BotGuardInitializedIterator, error) {

	logs, sub, err := _BotGuard.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BotGuardInitializedIterator{contract: _BotGuard.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BotGuard *BotGuardFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BotGuardInitialized) (event.Subscription, error) {

	logs, sub, err := _BotGuard.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BotGuardInitialized)
				if err := _BotGuard.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BotGuard *BotGuardFilterer) ParseInitialized(log types.Log) (*BotGuardInitialized, error) {
	event := new(BotGuardInitialized)
	if err := _BotGuard.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BotGuardOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BotGuard contract.
type BotGuardOwnershipTransferredIterator struct {
	Event *BotGuardOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BotGuardOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BotGuardOwnershipTransferred)
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
		it.Event = new(BotGuardOwnershipTransferred)
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
func (it *BotGuardOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BotGuardOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BotGuardOwnershipTransferred represents a OwnershipTransferred event raised by the BotGuard contract.
type BotGuardOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BotGuard *BotGuardFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BotGuardOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BotGuard.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BotGuardOwnershipTransferredIterator{contract: _BotGuard.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BotGuard *BotGuardFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BotGuardOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BotGuard.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BotGuardOwnershipTransferred)
				if err := _BotGuard.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BotGuard *BotGuardFilterer) ParseOwnershipTransferred(log types.Log) (*BotGuardOwnershipTransferred, error) {
	event := new(BotGuardOwnershipTransferred)
	if err := _BotGuard.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
