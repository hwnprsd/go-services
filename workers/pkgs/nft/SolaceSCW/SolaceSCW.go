// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SolaceSCW

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

// SolaceSCWMetaData contains all meta data concerning the SolaceSCW contract.
var SolaceSCWMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"scw\",\"type\":\"address\"}],\"name\":\"SCWCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createSCW\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getSCW\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userToSCW\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SolaceSCWABI is the input ABI used to generate the binding from.
// Deprecated: Use SolaceSCWMetaData.ABI instead.
var SolaceSCWABI = SolaceSCWMetaData.ABI

// SolaceSCW is an auto generated Go binding around an Ethereum contract.
type SolaceSCW struct {
	SolaceSCWCaller     // Read-only binding to the contract
	SolaceSCWTransactor // Write-only binding to the contract
	SolaceSCWFilterer   // Log filterer for contract events
}

// SolaceSCWCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolaceSCWCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolaceSCWTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolaceSCWTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolaceSCWFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolaceSCWFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolaceSCWSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolaceSCWSession struct {
	Contract     *SolaceSCW        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolaceSCWCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolaceSCWCallerSession struct {
	Contract *SolaceSCWCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SolaceSCWTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolaceSCWTransactorSession struct {
	Contract     *SolaceSCWTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SolaceSCWRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolaceSCWRaw struct {
	Contract *SolaceSCW // Generic contract binding to access the raw methods on
}

// SolaceSCWCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolaceSCWCallerRaw struct {
	Contract *SolaceSCWCaller // Generic read-only contract binding to access the raw methods on
}

// SolaceSCWTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolaceSCWTransactorRaw struct {
	Contract *SolaceSCWTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolaceSCW creates a new instance of SolaceSCW, bound to a specific deployed contract.
func NewSolaceSCW(address common.Address, backend bind.ContractBackend) (*SolaceSCW, error) {
	contract, err := bindSolaceSCW(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolaceSCW{SolaceSCWCaller: SolaceSCWCaller{contract: contract}, SolaceSCWTransactor: SolaceSCWTransactor{contract: contract}, SolaceSCWFilterer: SolaceSCWFilterer{contract: contract}}, nil
}

// NewSolaceSCWCaller creates a new read-only instance of SolaceSCW, bound to a specific deployed contract.
func NewSolaceSCWCaller(address common.Address, caller bind.ContractCaller) (*SolaceSCWCaller, error) {
	contract, err := bindSolaceSCW(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWCaller{contract: contract}, nil
}

// NewSolaceSCWTransactor creates a new write-only instance of SolaceSCW, bound to a specific deployed contract.
func NewSolaceSCWTransactor(address common.Address, transactor bind.ContractTransactor) (*SolaceSCWTransactor, error) {
	contract, err := bindSolaceSCW(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWTransactor{contract: contract}, nil
}

// NewSolaceSCWFilterer creates a new log filterer instance of SolaceSCW, bound to a specific deployed contract.
func NewSolaceSCWFilterer(address common.Address, filterer bind.ContractFilterer) (*SolaceSCWFilterer, error) {
	contract, err := bindSolaceSCW(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWFilterer{contract: contract}, nil
}

// bindSolaceSCW binds a generic wrapper to an already deployed contract.
func bindSolaceSCW(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolaceSCWABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolaceSCW *SolaceSCWRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolaceSCW.Contract.SolaceSCWCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolaceSCW *SolaceSCWRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolaceSCW.Contract.SolaceSCWTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolaceSCW *SolaceSCWRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolaceSCW.Contract.SolaceSCWTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolaceSCW *SolaceSCWCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolaceSCW.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolaceSCW *SolaceSCWTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolaceSCW.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolaceSCW *SolaceSCWTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolaceSCW.Contract.contract.Transact(opts, method, params...)
}

// GetSCW is a free data retrieval call binding the contract method 0x71c47320.
//
// Solidity: function getSCW(address user) view returns(address)
func (_SolaceSCW *SolaceSCWCaller) GetSCW(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "getSCW", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSCW is a free data retrieval call binding the contract method 0x71c47320.
//
// Solidity: function getSCW(address user) view returns(address)
func (_SolaceSCW *SolaceSCWSession) GetSCW(user common.Address) (common.Address, error) {
	return _SolaceSCW.Contract.GetSCW(&_SolaceSCW.CallOpts, user)
}

// GetSCW is a free data retrieval call binding the contract method 0x71c47320.
//
// Solidity: function getSCW(address user) view returns(address)
func (_SolaceSCW *SolaceSCWCallerSession) GetSCW(user common.Address) (common.Address, error) {
	return _SolaceSCW.Contract.GetSCW(&_SolaceSCW.CallOpts, user)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SolaceSCW *SolaceSCWCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SolaceSCW *SolaceSCWSession) Implementation() (common.Address, error) {
	return _SolaceSCW.Contract.Implementation(&_SolaceSCW.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SolaceSCW *SolaceSCWCallerSession) Implementation() (common.Address, error) {
	return _SolaceSCW.Contract.Implementation(&_SolaceSCW.CallOpts)
}

// UserToSCW is a free data retrieval call binding the contract method 0x5b47dc9b.
//
// Solidity: function userToSCW(address ) view returns(address)
func (_SolaceSCW *SolaceSCWCaller) UserToSCW(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "userToSCW", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserToSCW is a free data retrieval call binding the contract method 0x5b47dc9b.
//
// Solidity: function userToSCW(address ) view returns(address)
func (_SolaceSCW *SolaceSCWSession) UserToSCW(arg0 common.Address) (common.Address, error) {
	return _SolaceSCW.Contract.UserToSCW(&_SolaceSCW.CallOpts, arg0)
}

// UserToSCW is a free data retrieval call binding the contract method 0x5b47dc9b.
//
// Solidity: function userToSCW(address ) view returns(address)
func (_SolaceSCW *SolaceSCWCallerSession) UserToSCW(arg0 common.Address) (common.Address, error) {
	return _SolaceSCW.Contract.UserToSCW(&_SolaceSCW.CallOpts, arg0)
}

// CreateSCW is a paid mutator transaction binding the contract method 0x02f1a798.
//
// Solidity: function createSCW(address owner) returns(address)
func (_SolaceSCW *SolaceSCWTransactor) CreateSCW(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _SolaceSCW.contract.Transact(opts, "createSCW", owner)
}

// CreateSCW is a paid mutator transaction binding the contract method 0x02f1a798.
//
// Solidity: function createSCW(address owner) returns(address)
func (_SolaceSCW *SolaceSCWSession) CreateSCW(owner common.Address) (*types.Transaction, error) {
	return _SolaceSCW.Contract.CreateSCW(&_SolaceSCW.TransactOpts, owner)
}

// CreateSCW is a paid mutator transaction binding the contract method 0x02f1a798.
//
// Solidity: function createSCW(address owner) returns(address)
func (_SolaceSCW *SolaceSCWTransactorSession) CreateSCW(owner common.Address) (*types.Transaction, error) {
	return _SolaceSCW.Contract.CreateSCW(&_SolaceSCW.TransactOpts, owner)
}

// SolaceSCWSCWCreatedIterator is returned from FilterSCWCreated and is used to iterate over the raw logs and unpacked data for SCWCreated events raised by the SolaceSCW contract.
type SolaceSCWSCWCreatedIterator struct {
	Event *SolaceSCWSCWCreated // Event containing the contract specifics and raw log

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
func (it *SolaceSCWSCWCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolaceSCWSCWCreated)
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
		it.Event = new(SolaceSCWSCWCreated)
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
func (it *SolaceSCWSCWCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolaceSCWSCWCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolaceSCWSCWCreated represents a SCWCreated event raised by the SolaceSCW contract.
type SolaceSCWSCWCreated struct {
	Scw common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSCWCreated is a free log retrieval operation binding the contract event 0xa4570aa2cf110984d7d51e6607ec55a9f731e239a3a6eb692b40280a58d791cb.
//
// Solidity: event SCWCreated(address indexed scw)
func (_SolaceSCW *SolaceSCWFilterer) FilterSCWCreated(opts *bind.FilterOpts, scw []common.Address) (*SolaceSCWSCWCreatedIterator, error) {

	var scwRule []interface{}
	for _, scwItem := range scw {
		scwRule = append(scwRule, scwItem)
	}

	logs, sub, err := _SolaceSCW.contract.FilterLogs(opts, "SCWCreated", scwRule)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWSCWCreatedIterator{contract: _SolaceSCW.contract, event: "SCWCreated", logs: logs, sub: sub}, nil
}

// WatchSCWCreated is a free log subscription operation binding the contract event 0xa4570aa2cf110984d7d51e6607ec55a9f731e239a3a6eb692b40280a58d791cb.
//
// Solidity: event SCWCreated(address indexed scw)
func (_SolaceSCW *SolaceSCWFilterer) WatchSCWCreated(opts *bind.WatchOpts, sink chan<- *SolaceSCWSCWCreated, scw []common.Address) (event.Subscription, error) {

	var scwRule []interface{}
	for _, scwItem := range scw {
		scwRule = append(scwRule, scwItem)
	}

	logs, sub, err := _SolaceSCW.contract.WatchLogs(opts, "SCWCreated", scwRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolaceSCWSCWCreated)
				if err := _SolaceSCW.contract.UnpackLog(event, "SCWCreated", log); err != nil {
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

// ParseSCWCreated is a log parse operation binding the contract event 0xa4570aa2cf110984d7d51e6607ec55a9f731e239a3a6eb692b40280a58d791cb.
//
// Solidity: event SCWCreated(address indexed scw)
func (_SolaceSCW *SolaceSCWFilterer) ParseSCWCreated(log types.Log) (*SolaceSCWSCWCreated, error) {
	event := new(SolaceSCWSCWCreated)
	if err := _SolaceSCW.contract.UnpackLog(event, "SCWCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
