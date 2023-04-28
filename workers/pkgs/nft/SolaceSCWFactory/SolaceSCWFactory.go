// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SolaceSCWFactory

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

// SolaceSCWFactoryMetaData contains all meta data concerning the SolaceSCWFactory contract.
var SolaceSCWFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"scw\",\"type\":\"address\"}],\"name\":\"SCWCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createSCW\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getSCW\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userToSCW\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SolaceSCWFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SolaceSCWFactoryMetaData.ABI instead.
var SolaceSCWFactoryABI = SolaceSCWFactoryMetaData.ABI

// SolaceSCWFactory is an auto generated Go binding around an Ethereum contract.
type SolaceSCWFactory struct {
	SolaceSCWFactoryCaller     // Read-only binding to the contract
	SolaceSCWFactoryTransactor // Write-only binding to the contract
	SolaceSCWFactoryFilterer   // Log filterer for contract events
}

// SolaceSCWFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolaceSCWFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolaceSCWFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolaceSCWFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolaceSCWFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolaceSCWFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolaceSCWFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolaceSCWFactorySession struct {
	Contract     *SolaceSCWFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolaceSCWFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolaceSCWFactoryCallerSession struct {
	Contract *SolaceSCWFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SolaceSCWFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolaceSCWFactoryTransactorSession struct {
	Contract     *SolaceSCWFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SolaceSCWFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolaceSCWFactoryRaw struct {
	Contract *SolaceSCWFactory // Generic contract binding to access the raw methods on
}

// SolaceSCWFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolaceSCWFactoryCallerRaw struct {
	Contract *SolaceSCWFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SolaceSCWFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolaceSCWFactoryTransactorRaw struct {
	Contract *SolaceSCWFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolaceSCWFactory creates a new instance of SolaceSCWFactory, bound to a specific deployed contract.
func NewSolaceSCWFactory(address common.Address, backend bind.ContractBackend) (*SolaceSCWFactory, error) {
	contract, err := bindSolaceSCWFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWFactory{SolaceSCWFactoryCaller: SolaceSCWFactoryCaller{contract: contract}, SolaceSCWFactoryTransactor: SolaceSCWFactoryTransactor{contract: contract}, SolaceSCWFactoryFilterer: SolaceSCWFactoryFilterer{contract: contract}}, nil
}

// NewSolaceSCWFactoryCaller creates a new read-only instance of SolaceSCWFactory, bound to a specific deployed contract.
func NewSolaceSCWFactoryCaller(address common.Address, caller bind.ContractCaller) (*SolaceSCWFactoryCaller, error) {
	contract, err := bindSolaceSCWFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWFactoryCaller{contract: contract}, nil
}

// NewSolaceSCWFactoryTransactor creates a new write-only instance of SolaceSCWFactory, bound to a specific deployed contract.
func NewSolaceSCWFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SolaceSCWFactoryTransactor, error) {
	contract, err := bindSolaceSCWFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWFactoryTransactor{contract: contract}, nil
}

// NewSolaceSCWFactoryFilterer creates a new log filterer instance of SolaceSCWFactory, bound to a specific deployed contract.
func NewSolaceSCWFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SolaceSCWFactoryFilterer, error) {
	contract, err := bindSolaceSCWFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWFactoryFilterer{contract: contract}, nil
}

// bindSolaceSCWFactory binds a generic wrapper to an already deployed contract.
func bindSolaceSCWFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolaceSCWFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolaceSCWFactory *SolaceSCWFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolaceSCWFactory.Contract.SolaceSCWFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolaceSCWFactory *SolaceSCWFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolaceSCWFactory.Contract.SolaceSCWFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolaceSCWFactory *SolaceSCWFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolaceSCWFactory.Contract.SolaceSCWFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolaceSCWFactory *SolaceSCWFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolaceSCWFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolaceSCWFactory *SolaceSCWFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolaceSCWFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolaceSCWFactory *SolaceSCWFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolaceSCWFactory.Contract.contract.Transact(opts, method, params...)
}

// GetSCW is a free data retrieval call binding the contract method 0x71c47320.
//
// Solidity: function getSCW(address user) view returns(address)
func (_SolaceSCWFactory *SolaceSCWFactoryCaller) GetSCW(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _SolaceSCWFactory.contract.Call(opts, &out, "getSCW", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSCW is a free data retrieval call binding the contract method 0x71c47320.
//
// Solidity: function getSCW(address user) view returns(address)
func (_SolaceSCWFactory *SolaceSCWFactorySession) GetSCW(user common.Address) (common.Address, error) {
	return _SolaceSCWFactory.Contract.GetSCW(&_SolaceSCWFactory.CallOpts, user)
}

// GetSCW is a free data retrieval call binding the contract method 0x71c47320.
//
// Solidity: function getSCW(address user) view returns(address)
func (_SolaceSCWFactory *SolaceSCWFactoryCallerSession) GetSCW(user common.Address) (common.Address, error) {
	return _SolaceSCWFactory.Contract.GetSCW(&_SolaceSCWFactory.CallOpts, user)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SolaceSCWFactory *SolaceSCWFactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolaceSCWFactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SolaceSCWFactory *SolaceSCWFactorySession) Implementation() (common.Address, error) {
	return _SolaceSCWFactory.Contract.Implementation(&_SolaceSCWFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SolaceSCWFactory *SolaceSCWFactoryCallerSession) Implementation() (common.Address, error) {
	return _SolaceSCWFactory.Contract.Implementation(&_SolaceSCWFactory.CallOpts)
}

// UserToSCW is a free data retrieval call binding the contract method 0x5b47dc9b.
//
// Solidity: function userToSCW(address ) view returns(address)
func (_SolaceSCWFactory *SolaceSCWFactoryCaller) UserToSCW(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SolaceSCWFactory.contract.Call(opts, &out, "userToSCW", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserToSCW is a free data retrieval call binding the contract method 0x5b47dc9b.
//
// Solidity: function userToSCW(address ) view returns(address)
func (_SolaceSCWFactory *SolaceSCWFactorySession) UserToSCW(arg0 common.Address) (common.Address, error) {
	return _SolaceSCWFactory.Contract.UserToSCW(&_SolaceSCWFactory.CallOpts, arg0)
}

// UserToSCW is a free data retrieval call binding the contract method 0x5b47dc9b.
//
// Solidity: function userToSCW(address ) view returns(address)
func (_SolaceSCWFactory *SolaceSCWFactoryCallerSession) UserToSCW(arg0 common.Address) (common.Address, error) {
	return _SolaceSCWFactory.Contract.UserToSCW(&_SolaceSCWFactory.CallOpts, arg0)
}

// CreateSCW is a paid mutator transaction binding the contract method 0x02f1a798.
//
// Solidity: function createSCW(address owner) returns(address)
func (_SolaceSCWFactory *SolaceSCWFactoryTransactor) CreateSCW(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _SolaceSCWFactory.contract.Transact(opts, "createSCW", owner)
}

// CreateSCW is a paid mutator transaction binding the contract method 0x02f1a798.
//
// Solidity: function createSCW(address owner) returns(address)
func (_SolaceSCWFactory *SolaceSCWFactorySession) CreateSCW(owner common.Address) (*types.Transaction, error) {
	return _SolaceSCWFactory.Contract.CreateSCW(&_SolaceSCWFactory.TransactOpts, owner)
}

// CreateSCW is a paid mutator transaction binding the contract method 0x02f1a798.
//
// Solidity: function createSCW(address owner) returns(address)
func (_SolaceSCWFactory *SolaceSCWFactoryTransactorSession) CreateSCW(owner common.Address) (*types.Transaction, error) {
	return _SolaceSCWFactory.Contract.CreateSCW(&_SolaceSCWFactory.TransactOpts, owner)
}

// SolaceSCWFactorySCWCreatedIterator is returned from FilterSCWCreated and is used to iterate over the raw logs and unpacked data for SCWCreated events raised by the SolaceSCWFactory contract.
type SolaceSCWFactorySCWCreatedIterator struct {
	Event *SolaceSCWFactorySCWCreated // Event containing the contract specifics and raw log

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
func (it *SolaceSCWFactorySCWCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolaceSCWFactorySCWCreated)
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
		it.Event = new(SolaceSCWFactorySCWCreated)
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
func (it *SolaceSCWFactorySCWCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolaceSCWFactorySCWCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolaceSCWFactorySCWCreated represents a SCWCreated event raised by the SolaceSCWFactory contract.
type SolaceSCWFactorySCWCreated struct {
	Scw common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSCWCreated is a free log retrieval operation binding the contract event 0xa4570aa2cf110984d7d51e6607ec55a9f731e239a3a6eb692b40280a58d791cb.
//
// Solidity: event SCWCreated(address indexed scw)
func (_SolaceSCWFactory *SolaceSCWFactoryFilterer) FilterSCWCreated(opts *bind.FilterOpts, scw []common.Address) (*SolaceSCWFactorySCWCreatedIterator, error) {

	var scwRule []interface{}
	for _, scwItem := range scw {
		scwRule = append(scwRule, scwItem)
	}

	logs, sub, err := _SolaceSCWFactory.contract.FilterLogs(opts, "SCWCreated", scwRule)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWFactorySCWCreatedIterator{contract: _SolaceSCWFactory.contract, event: "SCWCreated", logs: logs, sub: sub}, nil
}

// WatchSCWCreated is a free log subscription operation binding the contract event 0xa4570aa2cf110984d7d51e6607ec55a9f731e239a3a6eb692b40280a58d791cb.
//
// Solidity: event SCWCreated(address indexed scw)
func (_SolaceSCWFactory *SolaceSCWFactoryFilterer) WatchSCWCreated(opts *bind.WatchOpts, sink chan<- *SolaceSCWFactorySCWCreated, scw []common.Address) (event.Subscription, error) {

	var scwRule []interface{}
	for _, scwItem := range scw {
		scwRule = append(scwRule, scwItem)
	}

	logs, sub, err := _SolaceSCWFactory.contract.WatchLogs(opts, "SCWCreated", scwRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolaceSCWFactorySCWCreated)
				if err := _SolaceSCWFactory.contract.UnpackLog(event, "SCWCreated", log); err != nil {
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
func (_SolaceSCWFactory *SolaceSCWFactoryFilterer) ParseSCWCreated(log types.Log) (*SolaceSCWFactorySCWCreated, error) {
	event := new(SolaceSCWFactorySCWCreated)
	if err := _SolaceSCWFactory.contract.UnpackLog(event, "SCWCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
