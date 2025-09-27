// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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
	_ = abi.ConvertType
)

// SourceoftruthMetaData contains all meta data concerning the Sourceoftruth contract.
var SourceoftruthMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"HashStored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"}],\"name\":\"storeHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"}],\"name\":\"checkHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"storedHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SourceoftruthABI is the input ABI used to generate the binding from.
// Deprecated: Use SourceoftruthMetaData.ABI instead.
var SourceoftruthABI = SourceoftruthMetaData.ABI

// Sourceoftruth is an auto generated Go binding around an Ethereum contract.
type Sourceoftruth struct {
	SourceoftruthCaller     // Read-only binding to the contract
	SourceoftruthTransactor // Write-only binding to the contract
	SourceoftruthFilterer   // Log filterer for contract events
}

// SourceoftruthCaller is an auto generated read-only Go binding around an Ethereum contract.
type SourceoftruthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SourceoftruthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SourceoftruthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SourceoftruthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SourceoftruthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SourceoftruthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SourceoftruthSession struct {
	Contract     *Sourceoftruth    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SourceoftruthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SourceoftruthCallerSession struct {
	Contract *SourceoftruthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SourceoftruthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SourceoftruthTransactorSession struct {
	Contract     *SourceoftruthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SourceoftruthRaw is an auto generated low-level Go binding around an Ethereum contract.
type SourceoftruthRaw struct {
	Contract *Sourceoftruth // Generic contract binding to access the raw methods on
}

// SourceoftruthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SourceoftruthCallerRaw struct {
	Contract *SourceoftruthCaller // Generic read-only contract binding to access the raw methods on
}

// SourceoftruthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SourceoftruthTransactorRaw struct {
	Contract *SourceoftruthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSourceoftruth creates a new instance of Sourceoftruth, bound to a specific deployed contract.
func NewSourceoftruth(address common.Address, backend bind.ContractBackend) (*Sourceoftruth, error) {
	contract, err := bindSourceoftruth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sourceoftruth{SourceoftruthCaller: SourceoftruthCaller{contract: contract}, SourceoftruthTransactor: SourceoftruthTransactor{contract: contract}, SourceoftruthFilterer: SourceoftruthFilterer{contract: contract}}, nil
}

// NewSourceoftruthCaller creates a new read-only instance of Sourceoftruth, bound to a specific deployed contract.
func NewSourceoftruthCaller(address common.Address, caller bind.ContractCaller) (*SourceoftruthCaller, error) {
	contract, err := bindSourceoftruth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SourceoftruthCaller{contract: contract}, nil
}

// NewSourceoftruthTransactor creates a new write-only instance of Sourceoftruth, bound to a specific deployed contract.
func NewSourceoftruthTransactor(address common.Address, transactor bind.ContractTransactor) (*SourceoftruthTransactor, error) {
	contract, err := bindSourceoftruth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SourceoftruthTransactor{contract: contract}, nil
}

// NewSourceoftruthFilterer creates a new log filterer instance of Sourceoftruth, bound to a specific deployed contract.
func NewSourceoftruthFilterer(address common.Address, filterer bind.ContractFilterer) (*SourceoftruthFilterer, error) {
	contract, err := bindSourceoftruth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SourceoftruthFilterer{contract: contract}, nil
}

// bindSourceoftruth binds a generic wrapper to an already deployed contract.
func bindSourceoftruth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SourceoftruthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sourceoftruth *SourceoftruthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sourceoftruth.Contract.SourceoftruthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sourceoftruth *SourceoftruthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sourceoftruth.Contract.SourceoftruthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sourceoftruth *SourceoftruthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sourceoftruth.Contract.SourceoftruthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sourceoftruth *SourceoftruthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sourceoftruth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sourceoftruth *SourceoftruthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sourceoftruth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sourceoftruth *SourceoftruthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sourceoftruth.Contract.contract.Transact(opts, method, params...)
}

// CheckHash is a free data retrieval call binding the contract method 0xe670f7cd.
//
// Solidity: function checkHash(string fileHash) view returns(bool)
func (_Sourceoftruth *SourceoftruthCaller) CheckHash(opts *bind.CallOpts, fileHash string) (bool, error) {
	var out []interface{}
	err := _Sourceoftruth.contract.Call(opts, &out, "checkHash", fileHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckHash is a free data retrieval call binding the contract method 0xe670f7cd.
//
// Solidity: function checkHash(string fileHash) view returns(bool)
func (_Sourceoftruth *SourceoftruthSession) CheckHash(fileHash string) (bool, error) {
	return _Sourceoftruth.Contract.CheckHash(&_Sourceoftruth.CallOpts, fileHash)
}

// CheckHash is a free data retrieval call binding the contract method 0xe670f7cd.
//
// Solidity: function checkHash(string fileHash) view returns(bool)
func (_Sourceoftruth *SourceoftruthCallerSession) CheckHash(fileHash string) (bool, error) {
	return _Sourceoftruth.Contract.CheckHash(&_Sourceoftruth.CallOpts, fileHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sourceoftruth *SourceoftruthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sourceoftruth.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sourceoftruth *SourceoftruthSession) Owner() (common.Address, error) {
	return _Sourceoftruth.Contract.Owner(&_Sourceoftruth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sourceoftruth *SourceoftruthCallerSession) Owner() (common.Address, error) {
	return _Sourceoftruth.Contract.Owner(&_Sourceoftruth.CallOpts)
}

// StoredHashes is a free data retrieval call binding the contract method 0x16663ea6.
//
// Solidity: function storedHashes(string ) view returns(bool)
func (_Sourceoftruth *SourceoftruthCaller) StoredHashes(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Sourceoftruth.contract.Call(opts, &out, "storedHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StoredHashes is a free data retrieval call binding the contract method 0x16663ea6.
//
// Solidity: function storedHashes(string ) view returns(bool)
func (_Sourceoftruth *SourceoftruthSession) StoredHashes(arg0 string) (bool, error) {
	return _Sourceoftruth.Contract.StoredHashes(&_Sourceoftruth.CallOpts, arg0)
}

// StoredHashes is a free data retrieval call binding the contract method 0x16663ea6.
//
// Solidity: function storedHashes(string ) view returns(bool)
func (_Sourceoftruth *SourceoftruthCallerSession) StoredHashes(arg0 string) (bool, error) {
	return _Sourceoftruth.Contract.StoredHashes(&_Sourceoftruth.CallOpts, arg0)
}

// StoreHash is a paid mutator transaction binding the contract method 0x71dc61cb.
//
// Solidity: function storeHash(string fileHash) returns()
func (_Sourceoftruth *SourceoftruthTransactor) StoreHash(opts *bind.TransactOpts, fileHash string) (*types.Transaction, error) {
	return _Sourceoftruth.contract.Transact(opts, "storeHash", fileHash)
}

// StoreHash is a paid mutator transaction binding the contract method 0x71dc61cb.
//
// Solidity: function storeHash(string fileHash) returns()
func (_Sourceoftruth *SourceoftruthSession) StoreHash(fileHash string) (*types.Transaction, error) {
	return _Sourceoftruth.Contract.StoreHash(&_Sourceoftruth.TransactOpts, fileHash)
}

// StoreHash is a paid mutator transaction binding the contract method 0x71dc61cb.
//
// Solidity: function storeHash(string fileHash) returns()
func (_Sourceoftruth *SourceoftruthTransactorSession) StoreHash(fileHash string) (*types.Transaction, error) {
	return _Sourceoftruth.Contract.StoreHash(&_Sourceoftruth.TransactOpts, fileHash)
}

// SourceoftruthHashStoredIterator is returned from FilterHashStored and is used to iterate over the raw logs and unpacked data for HashStored events raised by the Sourceoftruth contract.
type SourceoftruthHashStoredIterator struct {
	Event *SourceoftruthHashStored // Event containing the contract specifics and raw log

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
func (it *SourceoftruthHashStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SourceoftruthHashStored)
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
		it.Event = new(SourceoftruthHashStored)
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
func (it *SourceoftruthHashStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SourceoftruthHashStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SourceoftruthHashStored represents a HashStored event raised by the Sourceoftruth contract.
type SourceoftruthHashStored struct {
	FileHash  string
	StoredBy  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHashStored is a free log retrieval operation binding the contract event 0xaea7310eab8defd2229b800a4f9aa9397e8fa2e614f2721e5a34ded9221da4f9.
//
// Solidity: event HashStored(string fileHash, address indexed storedBy, uint256 timestamp)
func (_Sourceoftruth *SourceoftruthFilterer) FilterHashStored(opts *bind.FilterOpts, storedBy []common.Address) (*SourceoftruthHashStoredIterator, error) {

	var storedByRule []interface{}
	for _, storedByItem := range storedBy {
		storedByRule = append(storedByRule, storedByItem)
	}

	logs, sub, err := _Sourceoftruth.contract.FilterLogs(opts, "HashStored", storedByRule)
	if err != nil {
		return nil, err
	}
	return &SourceoftruthHashStoredIterator{contract: _Sourceoftruth.contract, event: "HashStored", logs: logs, sub: sub}, nil
}

// WatchHashStored is a free log subscription operation binding the contract event 0xaea7310eab8defd2229b800a4f9aa9397e8fa2e614f2721e5a34ded9221da4f9.
//
// Solidity: event HashStored(string fileHash, address indexed storedBy, uint256 timestamp)
func (_Sourceoftruth *SourceoftruthFilterer) WatchHashStored(opts *bind.WatchOpts, sink chan<- *SourceoftruthHashStored, storedBy []common.Address) (event.Subscription, error) {

	var storedByRule []interface{}
	for _, storedByItem := range storedBy {
		storedByRule = append(storedByRule, storedByItem)
	}

	logs, sub, err := _Sourceoftruth.contract.WatchLogs(opts, "HashStored", storedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SourceoftruthHashStored)
				if err := _Sourceoftruth.contract.UnpackLog(event, "HashStored", log); err != nil {
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

// ParseHashStored is a log parse operation binding the contract event 0xaea7310eab8defd2229b800a4f9aa9397e8fa2e614f2721e5a34ded9221da4f9.
//
// Solidity: event HashStored(string fileHash, address indexed storedBy, uint256 timestamp)
func (_Sourceoftruth *SourceoftruthFilterer) ParseHashStored(log types.Log) (*SourceoftruthHashStored, error) {
	event := new(SourceoftruthHashStored)
	if err := _Sourceoftruth.contract.UnpackLog(event, "HashStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}