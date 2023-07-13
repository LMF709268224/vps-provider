package blockchain // Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

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

// OrderMetaData contains all meta data concerning the Order contract.
var OrderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderPlaced\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"placeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OrderABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderMetaData.ABI instead.
var OrderABI = OrderMetaData.ABI

// Order is an auto generated Go binding around an Ethereum contract.
type Order struct {
	OrderCaller     // Read-only binding to the contract
	OrderTransactor // Write-only binding to the contract
	OrderFilterer   // Log filterer for contract events
}

// OrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderSession struct {
	Contract     *Order            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderCallerSession struct {
	Contract *OrderCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderTransactorSession struct {
	Contract     *OrderTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderRaw struct {
	Contract *Order // Generic contract binding to access the raw methods on
}

// OrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderCallerRaw struct {
	Contract *OrderCaller // Generic read-only contract binding to access the raw methods on
}

// OrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderTransactorRaw struct {
	Contract *OrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrder creates a new instance of Order, bound to a specific deployed contract.
func NewOrder(address common.Address, backend bind.ContractBackend) (*Order, error) {
	contract, err := bindOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Order{OrderCaller: OrderCaller{contract: contract}, OrderTransactor: OrderTransactor{contract: contract}, OrderFilterer: OrderFilterer{contract: contract}}, nil
}

// NewOrderCaller creates a new read-only instance of Order, bound to a specific deployed contract.
func NewOrderCaller(address common.Address, caller bind.ContractCaller) (*OrderCaller, error) {
	contract, err := bindOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderCaller{contract: contract}, nil
}

// NewOrderTransactor creates a new write-only instance of Order, bound to a specific deployed contract.
func NewOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderTransactor, error) {
	contract, err := bindOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderTransactor{contract: contract}, nil
}

// NewOrderFilterer creates a new log filterer instance of Order, bound to a specific deployed contract.
func NewOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderFilterer, error) {
	contract, err := bindOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderFilterer{contract: contract}, nil
}

// bindOrder binds a generic wrapper to an already deployed contract.
func bindOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Order *OrderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Order.Contract.OrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Order *OrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Order.Contract.OrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Order *OrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Order.Contract.OrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Order *OrderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Order.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Order *OrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Order.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Order *OrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Order.Contract.contract.Transact(opts, method, params...)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 _orderId) view returns(uint256, uint256, uint256)
func (_Order *OrderCaller) GetOrder(opts *bind.CallOpts, _orderId *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "getOrder", _orderId)
	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 _orderId) view returns(uint256, uint256, uint256)
func (_Order *OrderSession) GetOrder(_orderId *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Order.Contract.GetOrder(&_Order.CallOpts, _orderId)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 _orderId) view returns(uint256, uint256, uint256)
func (_Order *OrderCallerSession) GetOrder(_orderId *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Order.Contract.GetOrder(&_Order.CallOpts, _orderId)
}

// GetOrderCount is a free data retrieval call binding the contract method 0x8d0a5fbb.
//
// Solidity: function getOrderCount() view returns(uint256)
func (_Order *OrderCaller) GetOrderCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "getOrderCount")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// GetOrderCount is a free data retrieval call binding the contract method 0x8d0a5fbb.
//
// Solidity: function getOrderCount() view returns(uint256)
func (_Order *OrderSession) GetOrderCount() (*big.Int, error) {
	return _Order.Contract.GetOrderCount(&_Order.CallOpts)
}

// GetOrderCount is a free data retrieval call binding the contract method 0x8d0a5fbb.
//
// Solidity: function getOrderCount() view returns(uint256)
func (_Order *OrderCallerSession) GetOrderCount() (*big.Int, error) {
	return _Order.Contract.GetOrderCount(&_Order.CallOpts)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x843f61e2.
//
// Solidity: function placeOrder(uint256 _orderId, uint256 _amount) returns()
func (_Order *OrderTransactor) PlaceOrder(opts *bind.TransactOpts, _orderId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "placeOrder", _orderId, _amount)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x843f61e2.
//
// Solidity: function placeOrder(uint256 _orderId, uint256 _amount) returns()
func (_Order *OrderSession) PlaceOrder(_orderId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Order.Contract.PlaceOrder(&_Order.TransactOpts, _orderId, _amount)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x843f61e2.
//
// Solidity: function placeOrder(uint256 _orderId, uint256 _amount) returns()
func (_Order *OrderTransactorSession) PlaceOrder(_orderId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Order.Contract.PlaceOrder(&_Order.TransactOpts, _orderId, _amount)
}

// OrderOrderPlacedIterator is returned from FilterOrderPlaced and is used to iterate over the raw logs and unpacked data for OrderPlaced events raised by the Order contract.
type OrderOrderPlacedIterator struct {
	Event *OrderOrderPlaced // Event containing the contract specifics and raw log

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
func (it *OrderOrderPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderOrderPlaced)
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
		it.Event = new(OrderOrderPlaced)
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
func (it *OrderOrderPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderOrderPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderOrderPlaced represents a OrderPlaced event raised by the Order contract.
type OrderOrderPlaced struct {
	OrderId   *big.Int
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderPlaced is a free log retrieval operation binding the contract event 0x4e3753386201ff4fa2fe7c9a2bdce7cea2028030bb57c7596e089ec9aa5feadb.
//
// Solidity: event OrderPlaced(uint256 indexed orderId, uint256 amount, uint256 timestamp)
func (_Order *OrderFilterer) FilterOrderPlaced(opts *bind.FilterOpts, orderId []*big.Int) (*OrderOrderPlacedIterator, error) {
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Order.contract.FilterLogs(opts, "OrderPlaced", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &OrderOrderPlacedIterator{contract: _Order.contract, event: "OrderPlaced", logs: logs, sub: sub}, nil
}

// WatchOrderPlaced is a free log subscription operation binding the contract event 0x4e3753386201ff4fa2fe7c9a2bdce7cea2028030bb57c7596e089ec9aa5feadb.
//
// Solidity: event OrderPlaced(uint256 indexed orderId, uint256 amount, uint256 timestamp)
func (_Order *OrderFilterer) WatchOrderPlaced(opts *bind.WatchOpts, sink chan<- *OrderOrderPlaced, orderId []*big.Int) (event.Subscription, error) {
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Order.contract.WatchLogs(opts, "OrderPlaced", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderOrderPlaced)
				if err := _Order.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
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

// ParseOrderPlaced is a log parse operation binding the contract event 0x4e3753386201ff4fa2fe7c9a2bdce7cea2028030bb57c7596e089ec9aa5feadb.
//
// Solidity: event OrderPlaced(uint256 indexed orderId, uint256 amount, uint256 timestamp)
func (_Order *OrderFilterer) ParseOrderPlaced(log types.Log) (*OrderOrderPlaced, error) {
	event := new(OrderOrderPlaced)
	if err := _Order.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
