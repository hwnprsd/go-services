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

// MinimalForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type MinimalForwarderForwardRequest struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Gas   *big.Int
	Nonce *big.Int
	Data  []byte
}

// SolaceSCWMetaData contains all meta data concerning the SolaceSCW contract.
var SolaceSCWMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTokenAmount\",\"type\":\"uint256\"}],\"name\":\"EphemeralSignerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"EphemeralSignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTokenAmount\",\"type\":\"uint256\"}],\"name\":\"addEphemeralSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ephemeralSigners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"removeEphemeralSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMinimalForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// EphemeralSigners is a free data retrieval call binding the contract method 0xde3eb583.
//
// Solidity: function ephemeralSigners(address , address ) view returns(uint256 expiryTime, uint256 maxTokenAmount)
func (_SolaceSCW *SolaceSCWCaller) EphemeralSigners(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	ExpiryTime     *big.Int
	MaxTokenAmount *big.Int
}, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "ephemeralSigners", arg0, arg1)

	outstruct := new(struct {
		ExpiryTime     *big.Int
		MaxTokenAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExpiryTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxTokenAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EphemeralSigners is a free data retrieval call binding the contract method 0xde3eb583.
//
// Solidity: function ephemeralSigners(address , address ) view returns(uint256 expiryTime, uint256 maxTokenAmount)
func (_SolaceSCW *SolaceSCWSession) EphemeralSigners(arg0 common.Address, arg1 common.Address) (struct {
	ExpiryTime     *big.Int
	MaxTokenAmount *big.Int
}, error) {
	return _SolaceSCW.Contract.EphemeralSigners(&_SolaceSCW.CallOpts, arg0, arg1)
}

// EphemeralSigners is a free data retrieval call binding the contract method 0xde3eb583.
//
// Solidity: function ephemeralSigners(address , address ) view returns(uint256 expiryTime, uint256 maxTokenAmount)
func (_SolaceSCW *SolaceSCWCallerSession) EphemeralSigners(arg0 common.Address, arg1 common.Address) (struct {
	ExpiryTime     *big.Int
	MaxTokenAmount *big.Int
}, error) {
	return _SolaceSCW.Contract.EphemeralSigners(&_SolaceSCW.CallOpts, arg0, arg1)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_SolaceSCW *SolaceSCWCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_SolaceSCW *SolaceSCWSession) GetNonce(from common.Address) (*big.Int, error) {
	return _SolaceSCW.Contract.GetNonce(&_SolaceSCW.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_SolaceSCW *SolaceSCWCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _SolaceSCW.Contract.GetNonce(&_SolaceSCW.CallOpts, from)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_SolaceSCW *SolaceSCWCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_SolaceSCW *SolaceSCWSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _SolaceSCW.Contract.OnERC1155BatchReceived(&_SolaceSCW.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_SolaceSCW *SolaceSCWCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _SolaceSCW.Contract.OnERC1155BatchReceived(&_SolaceSCW.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_SolaceSCW *SolaceSCWCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_SolaceSCW *SolaceSCWSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _SolaceSCW.Contract.OnERC1155Received(&_SolaceSCW.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_SolaceSCW *SolaceSCWCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _SolaceSCW.Contract.OnERC1155Received(&_SolaceSCW.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_SolaceSCW *SolaceSCWCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_SolaceSCW *SolaceSCWSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _SolaceSCW.Contract.OnERC721Received(&_SolaceSCW.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_SolaceSCW *SolaceSCWCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _SolaceSCW.Contract.OnERC721Received(&_SolaceSCW.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SolaceSCW *SolaceSCWCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SolaceSCW *SolaceSCWSession) Owner() (common.Address, error) {
	return _SolaceSCW.Contract.Owner(&_SolaceSCW.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SolaceSCW *SolaceSCWCallerSession) Owner() (common.Address, error) {
	return _SolaceSCW.Contract.Owner(&_SolaceSCW.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SolaceSCW *SolaceSCWCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SolaceSCW *SolaceSCWSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SolaceSCW.Contract.SupportsInterface(&_SolaceSCW.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SolaceSCW *SolaceSCWCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SolaceSCW.Contract.SupportsInterface(&_SolaceSCW.CallOpts, interfaceId)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_SolaceSCW *SolaceSCWCaller) Verify(opts *bind.CallOpts, req MinimalForwarderForwardRequest, signature []byte) (bool, error) {
	var out []interface{}
	err := _SolaceSCW.contract.Call(opts, &out, "verify", req, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_SolaceSCW *SolaceSCWSession) Verify(req MinimalForwarderForwardRequest, signature []byte) (bool, error) {
	return _SolaceSCW.Contract.Verify(&_SolaceSCW.CallOpts, req, signature)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_SolaceSCW *SolaceSCWCallerSession) Verify(req MinimalForwarderForwardRequest, signature []byte) (bool, error) {
	return _SolaceSCW.Contract.Verify(&_SolaceSCW.CallOpts, req, signature)
}

// AddEphemeralSigner is a paid mutator transaction binding the contract method 0x7a4c169e.
//
// Solidity: function addEphemeralSigner(address _signer, address _contractAddress, uint256 _expiryTime, uint256 _maxTokenAmount) returns()
func (_SolaceSCW *SolaceSCWTransactor) AddEphemeralSigner(opts *bind.TransactOpts, _signer common.Address, _contractAddress common.Address, _expiryTime *big.Int, _maxTokenAmount *big.Int) (*types.Transaction, error) {
	return _SolaceSCW.contract.Transact(opts, "addEphemeralSigner", _signer, _contractAddress, _expiryTime, _maxTokenAmount)
}

// AddEphemeralSigner is a paid mutator transaction binding the contract method 0x7a4c169e.
//
// Solidity: function addEphemeralSigner(address _signer, address _contractAddress, uint256 _expiryTime, uint256 _maxTokenAmount) returns()
func (_SolaceSCW *SolaceSCWSession) AddEphemeralSigner(_signer common.Address, _contractAddress common.Address, _expiryTime *big.Int, _maxTokenAmount *big.Int) (*types.Transaction, error) {
	return _SolaceSCW.Contract.AddEphemeralSigner(&_SolaceSCW.TransactOpts, _signer, _contractAddress, _expiryTime, _maxTokenAmount)
}

// AddEphemeralSigner is a paid mutator transaction binding the contract method 0x7a4c169e.
//
// Solidity: function addEphemeralSigner(address _signer, address _contractAddress, uint256 _expiryTime, uint256 _maxTokenAmount) returns()
func (_SolaceSCW *SolaceSCWTransactorSession) AddEphemeralSigner(_signer common.Address, _contractAddress common.Address, _expiryTime *big.Int, _maxTokenAmount *big.Int) (*types.Transaction, error) {
	return _SolaceSCW.Contract.AddEphemeralSigner(&_SolaceSCW.TransactOpts, _signer, _contractAddress, _expiryTime, _maxTokenAmount)
}

// BatchTransferERC721 is a paid mutator transaction binding the contract method 0x081d536f.
//
// Solidity: function batchTransferERC721(address _nft, address _recipient, uint256[] _tokenIds) returns()
func (_SolaceSCW *SolaceSCWTransactor) BatchTransferERC721(opts *bind.TransactOpts, _nft common.Address, _recipient common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _SolaceSCW.contract.Transact(opts, "batchTransferERC721", _nft, _recipient, _tokenIds)
}

// BatchTransferERC721 is a paid mutator transaction binding the contract method 0x081d536f.
//
// Solidity: function batchTransferERC721(address _nft, address _recipient, uint256[] _tokenIds) returns()
func (_SolaceSCW *SolaceSCWSession) BatchTransferERC721(_nft common.Address, _recipient common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _SolaceSCW.Contract.BatchTransferERC721(&_SolaceSCW.TransactOpts, _nft, _recipient, _tokenIds)
}

// BatchTransferERC721 is a paid mutator transaction binding the contract method 0x081d536f.
//
// Solidity: function batchTransferERC721(address _nft, address _recipient, uint256[] _tokenIds) returns()
func (_SolaceSCW *SolaceSCWTransactorSession) BatchTransferERC721(_nft common.Address, _recipient common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _SolaceSCW.Contract.BatchTransferERC721(&_SolaceSCW.TransactOpts, _nft, _recipient, _tokenIds)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_SolaceSCW *SolaceSCWTransactor) Execute(opts *bind.TransactOpts, req MinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _SolaceSCW.contract.Transact(opts, "execute", req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_SolaceSCW *SolaceSCWSession) Execute(req MinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _SolaceSCW.Contract.Execute(&_SolaceSCW.TransactOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_SolaceSCW *SolaceSCWTransactorSession) Execute(req MinimalForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _SolaceSCW.Contract.Execute(&_SolaceSCW.TransactOpts, req, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_SolaceSCW *SolaceSCWTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _SolaceSCW.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_SolaceSCW *SolaceSCWSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _SolaceSCW.Contract.Initialize(&_SolaceSCW.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_SolaceSCW *SolaceSCWTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _SolaceSCW.Contract.Initialize(&_SolaceSCW.TransactOpts, _owner)
}

// RemoveEphemeralSigner is a paid mutator transaction binding the contract method 0x93cd995f.
//
// Solidity: function removeEphemeralSigner(address _signer, address _contractAddress) returns()
func (_SolaceSCW *SolaceSCWTransactor) RemoveEphemeralSigner(opts *bind.TransactOpts, _signer common.Address, _contractAddress common.Address) (*types.Transaction, error) {
	return _SolaceSCW.contract.Transact(opts, "removeEphemeralSigner", _signer, _contractAddress)
}

// RemoveEphemeralSigner is a paid mutator transaction binding the contract method 0x93cd995f.
//
// Solidity: function removeEphemeralSigner(address _signer, address _contractAddress) returns()
func (_SolaceSCW *SolaceSCWSession) RemoveEphemeralSigner(_signer common.Address, _contractAddress common.Address) (*types.Transaction, error) {
	return _SolaceSCW.Contract.RemoveEphemeralSigner(&_SolaceSCW.TransactOpts, _signer, _contractAddress)
}

// RemoveEphemeralSigner is a paid mutator transaction binding the contract method 0x93cd995f.
//
// Solidity: function removeEphemeralSigner(address _signer, address _contractAddress) returns()
func (_SolaceSCW *SolaceSCWTransactorSession) RemoveEphemeralSigner(_signer common.Address, _contractAddress common.Address) (*types.Transaction, error) {
	return _SolaceSCW.Contract.RemoveEphemeralSigner(&_SolaceSCW.TransactOpts, _signer, _contractAddress)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0xdbecc616.
//
// Solidity: function transferERC1155(address _token, address _recipient, uint256 _id, uint256 _amount, bytes _data) returns()
func (_SolaceSCW *SolaceSCWTransactor) TransferERC1155(opts *bind.TransactOpts, _token common.Address, _recipient common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _SolaceSCW.contract.Transact(opts, "transferERC1155", _token, _recipient, _id, _amount, _data)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0xdbecc616.
//
// Solidity: function transferERC1155(address _token, address _recipient, uint256 _id, uint256 _amount, bytes _data) returns()
func (_SolaceSCW *SolaceSCWSession) TransferERC1155(_token common.Address, _recipient common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _SolaceSCW.Contract.TransferERC1155(&_SolaceSCW.TransactOpts, _token, _recipient, _id, _amount, _data)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0xdbecc616.
//
// Solidity: function transferERC1155(address _token, address _recipient, uint256 _id, uint256 _amount, bytes _data) returns()
func (_SolaceSCW *SolaceSCWTransactorSession) TransferERC1155(_token common.Address, _recipient common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _SolaceSCW.Contract.TransferERC1155(&_SolaceSCW.TransactOpts, _token, _recipient, _id, _amount, _data)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address _token, address _recipient, uint256 _amount) returns()
func (_SolaceSCW *SolaceSCWTransactor) TransferERC20(opts *bind.TransactOpts, _token common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SolaceSCW.contract.Transact(opts, "transferERC20", _token, _recipient, _amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address _token, address _recipient, uint256 _amount) returns()
func (_SolaceSCW *SolaceSCWSession) TransferERC20(_token common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SolaceSCW.Contract.TransferERC20(&_SolaceSCW.TransactOpts, _token, _recipient, _amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address _token, address _recipient, uint256 _amount) returns()
func (_SolaceSCW *SolaceSCWTransactorSession) TransferERC20(_token common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SolaceSCW.Contract.TransferERC20(&_SolaceSCW.TransactOpts, _token, _recipient, _amount)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address _nft, address _recipient, uint256 _tokenId) returns()
func (_SolaceSCW *SolaceSCWTransactor) TransferERC721(opts *bind.TransactOpts, _nft common.Address, _recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SolaceSCW.contract.Transact(opts, "transferERC721", _nft, _recipient, _tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address _nft, address _recipient, uint256 _tokenId) returns()
func (_SolaceSCW *SolaceSCWSession) TransferERC721(_nft common.Address, _recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SolaceSCW.Contract.TransferERC721(&_SolaceSCW.TransactOpts, _nft, _recipient, _tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address _nft, address _recipient, uint256 _tokenId) returns()
func (_SolaceSCW *SolaceSCWTransactorSession) TransferERC721(_nft common.Address, _recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SolaceSCW.Contract.TransferERC721(&_SolaceSCW.TransactOpts, _nft, _recipient, _tokenId)
}

// SolaceSCWERC20TransferredIterator is returned from FilterERC20Transferred and is used to iterate over the raw logs and unpacked data for ERC20Transferred events raised by the SolaceSCW contract.
type SolaceSCWERC20TransferredIterator struct {
	Event *SolaceSCWERC20Transferred // Event containing the contract specifics and raw log

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
func (it *SolaceSCWERC20TransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolaceSCWERC20Transferred)
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
		it.Event = new(SolaceSCWERC20Transferred)
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
func (it *SolaceSCWERC20TransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolaceSCWERC20TransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolaceSCWERC20Transferred represents a ERC20Transferred event raised by the SolaceSCW contract.
type SolaceSCWERC20Transferred struct {
	Signer    common.Address
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Transferred is a free log retrieval operation binding the contract event 0x8dfac83616afc4a99a89fe63243ae807798d10eb5a1449e5472853efc3c022cf.
//
// Solidity: event ERC20Transferred(address indexed signer, address indexed token, address indexed recipient, uint256 amount)
func (_SolaceSCW *SolaceSCWFilterer) FilterERC20Transferred(opts *bind.FilterOpts, signer []common.Address, token []common.Address, recipient []common.Address) (*SolaceSCWERC20TransferredIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SolaceSCW.contract.FilterLogs(opts, "ERC20Transferred", signerRule, tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWERC20TransferredIterator{contract: _SolaceSCW.contract, event: "ERC20Transferred", logs: logs, sub: sub}, nil
}

// WatchERC20Transferred is a free log subscription operation binding the contract event 0x8dfac83616afc4a99a89fe63243ae807798d10eb5a1449e5472853efc3c022cf.
//
// Solidity: event ERC20Transferred(address indexed signer, address indexed token, address indexed recipient, uint256 amount)
func (_SolaceSCW *SolaceSCWFilterer) WatchERC20Transferred(opts *bind.WatchOpts, sink chan<- *SolaceSCWERC20Transferred, signer []common.Address, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SolaceSCW.contract.WatchLogs(opts, "ERC20Transferred", signerRule, tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolaceSCWERC20Transferred)
				if err := _SolaceSCW.contract.UnpackLog(event, "ERC20Transferred", log); err != nil {
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

// ParseERC20Transferred is a log parse operation binding the contract event 0x8dfac83616afc4a99a89fe63243ae807798d10eb5a1449e5472853efc3c022cf.
//
// Solidity: event ERC20Transferred(address indexed signer, address indexed token, address indexed recipient, uint256 amount)
func (_SolaceSCW *SolaceSCWFilterer) ParseERC20Transferred(log types.Log) (*SolaceSCWERC20Transferred, error) {
	event := new(SolaceSCWERC20Transferred)
	if err := _SolaceSCW.contract.UnpackLog(event, "ERC20Transferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolaceSCWEphemeralSignerAddedIterator is returned from FilterEphemeralSignerAdded and is used to iterate over the raw logs and unpacked data for EphemeralSignerAdded events raised by the SolaceSCW contract.
type SolaceSCWEphemeralSignerAddedIterator struct {
	Event *SolaceSCWEphemeralSignerAdded // Event containing the contract specifics and raw log

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
func (it *SolaceSCWEphemeralSignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolaceSCWEphemeralSignerAdded)
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
		it.Event = new(SolaceSCWEphemeralSignerAdded)
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
func (it *SolaceSCWEphemeralSignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolaceSCWEphemeralSignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolaceSCWEphemeralSignerAdded represents a EphemeralSignerAdded event raised by the SolaceSCW contract.
type SolaceSCWEphemeralSignerAdded struct {
	Signer          common.Address
	ContractAddress common.Address
	ExpiryTime      *big.Int
	MaxTokenAmount  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEphemeralSignerAdded is a free log retrieval operation binding the contract event 0x9aff964c4697101f6df989f49e930aa78f79c975292b98c3361bbcd000ce33bb.
//
// Solidity: event EphemeralSignerAdded(address indexed signer, address indexed contractAddress, uint256 expiryTime, uint256 maxTokenAmount)
func (_SolaceSCW *SolaceSCWFilterer) FilterEphemeralSignerAdded(opts *bind.FilterOpts, signer []common.Address, contractAddress []common.Address) (*SolaceSCWEphemeralSignerAddedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _SolaceSCW.contract.FilterLogs(opts, "EphemeralSignerAdded", signerRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWEphemeralSignerAddedIterator{contract: _SolaceSCW.contract, event: "EphemeralSignerAdded", logs: logs, sub: sub}, nil
}

// WatchEphemeralSignerAdded is a free log subscription operation binding the contract event 0x9aff964c4697101f6df989f49e930aa78f79c975292b98c3361bbcd000ce33bb.
//
// Solidity: event EphemeralSignerAdded(address indexed signer, address indexed contractAddress, uint256 expiryTime, uint256 maxTokenAmount)
func (_SolaceSCW *SolaceSCWFilterer) WatchEphemeralSignerAdded(opts *bind.WatchOpts, sink chan<- *SolaceSCWEphemeralSignerAdded, signer []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _SolaceSCW.contract.WatchLogs(opts, "EphemeralSignerAdded", signerRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolaceSCWEphemeralSignerAdded)
				if err := _SolaceSCW.contract.UnpackLog(event, "EphemeralSignerAdded", log); err != nil {
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

// ParseEphemeralSignerAdded is a log parse operation binding the contract event 0x9aff964c4697101f6df989f49e930aa78f79c975292b98c3361bbcd000ce33bb.
//
// Solidity: event EphemeralSignerAdded(address indexed signer, address indexed contractAddress, uint256 expiryTime, uint256 maxTokenAmount)
func (_SolaceSCW *SolaceSCWFilterer) ParseEphemeralSignerAdded(log types.Log) (*SolaceSCWEphemeralSignerAdded, error) {
	event := new(SolaceSCWEphemeralSignerAdded)
	if err := _SolaceSCW.contract.UnpackLog(event, "EphemeralSignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolaceSCWEphemeralSignerRemovedIterator is returned from FilterEphemeralSignerRemoved and is used to iterate over the raw logs and unpacked data for EphemeralSignerRemoved events raised by the SolaceSCW contract.
type SolaceSCWEphemeralSignerRemovedIterator struct {
	Event *SolaceSCWEphemeralSignerRemoved // Event containing the contract specifics and raw log

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
func (it *SolaceSCWEphemeralSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolaceSCWEphemeralSignerRemoved)
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
		it.Event = new(SolaceSCWEphemeralSignerRemoved)
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
func (it *SolaceSCWEphemeralSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolaceSCWEphemeralSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolaceSCWEphemeralSignerRemoved represents a EphemeralSignerRemoved event raised by the SolaceSCW contract.
type SolaceSCWEphemeralSignerRemoved struct {
	Signer          common.Address
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEphemeralSignerRemoved is a free log retrieval operation binding the contract event 0xa281e9a61766387cc8cc81cd89e6af1fa54e3e7c26065087cdc49384af718c78.
//
// Solidity: event EphemeralSignerRemoved(address indexed signer, address indexed contractAddress)
func (_SolaceSCW *SolaceSCWFilterer) FilterEphemeralSignerRemoved(opts *bind.FilterOpts, signer []common.Address, contractAddress []common.Address) (*SolaceSCWEphemeralSignerRemovedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _SolaceSCW.contract.FilterLogs(opts, "EphemeralSignerRemoved", signerRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &SolaceSCWEphemeralSignerRemovedIterator{contract: _SolaceSCW.contract, event: "EphemeralSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchEphemeralSignerRemoved is a free log subscription operation binding the contract event 0xa281e9a61766387cc8cc81cd89e6af1fa54e3e7c26065087cdc49384af718c78.
//
// Solidity: event EphemeralSignerRemoved(address indexed signer, address indexed contractAddress)
func (_SolaceSCW *SolaceSCWFilterer) WatchEphemeralSignerRemoved(opts *bind.WatchOpts, sink chan<- *SolaceSCWEphemeralSignerRemoved, signer []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _SolaceSCW.contract.WatchLogs(opts, "EphemeralSignerRemoved", signerRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolaceSCWEphemeralSignerRemoved)
				if err := _SolaceSCW.contract.UnpackLog(event, "EphemeralSignerRemoved", log); err != nil {
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

// ParseEphemeralSignerRemoved is a log parse operation binding the contract event 0xa281e9a61766387cc8cc81cd89e6af1fa54e3e7c26065087cdc49384af718c78.
//
// Solidity: event EphemeralSignerRemoved(address indexed signer, address indexed contractAddress)
func (_SolaceSCW *SolaceSCWFilterer) ParseEphemeralSignerRemoved(log types.Log) (*SolaceSCWEphemeralSignerRemoved, error) {
	event := new(SolaceSCWEphemeralSignerRemoved)
	if err := _SolaceSCW.contract.UnpackLog(event, "EphemeralSignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolaceSCWInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SolaceSCW contract.
type SolaceSCWInitializedIterator struct {
	Event *SolaceSCWInitialized // Event containing the contract specifics and raw log

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
func (it *SolaceSCWInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolaceSCWInitialized)
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
		it.Event = new(SolaceSCWInitialized)
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
func (it *SolaceSCWInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolaceSCWInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolaceSCWInitialized represents a Initialized event raised by the SolaceSCW contract.
type SolaceSCWInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SolaceSCW *SolaceSCWFilterer) FilterInitialized(opts *bind.FilterOpts) (*SolaceSCWInitializedIterator, error) {

	logs, sub, err := _SolaceSCW.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SolaceSCWInitializedIterator{contract: _SolaceSCW.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SolaceSCW *SolaceSCWFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SolaceSCWInitialized) (event.Subscription, error) {

	logs, sub, err := _SolaceSCW.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolaceSCWInitialized)
				if err := _SolaceSCW.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SolaceSCW *SolaceSCWFilterer) ParseInitialized(log types.Log) (*SolaceSCWInitialized, error) {
	event := new(SolaceSCWInitialized)
	if err := _SolaceSCW.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
