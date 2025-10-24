// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CycleAssetAddExistTokenBatchParam is an auto generated low-level Go binding around an user-defined struct.
type CycleAssetAddExistTokenBatchParam struct {
	TokenId        uint32
	Token          common.Address
	SharedDecimals uint8
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_polygonZkEVMBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountBlacklisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountNotFormatted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeExceedsCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCallData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"InvalidDecimals\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"}],\"name\":\"InvalidDestinationNetworkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNetworkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"}],\"name\":\"NotCounterPartContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToBridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCycleSupportRouting\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCycleSupportSettingBridgeMessageWithValueStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPolygonZkEVMBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"}],\"name\":\"TokenAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIdIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"}],\"name\":\"TokenNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expect\",\"type\":\"uint256\"}],\"name\":\"UnexpectedTransferAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"sharedDecimals\",\"type\":\"uint8\"}],\"name\":\"AddToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"destinationFunc\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"sharedDecimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"Bridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"destinationFunc\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"sharedDecimals\",\"type\":\"uint8\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"DeleteBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"destinationFunc\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"sharedDecimals\",\"type\":\"uint8\"}],\"name\":\"Forward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"addBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sharedDecimals\",\"type\":\"uint8\"}],\"name\":\"addExistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sharedDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structCycleAsset.AddExistTokenBatchParam[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"addExistTokenBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"addNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"localDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"sharedDecimals\",\"type\":\"uint8\"}],\"name\":\"addWrappedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"bridgeMessageWithValueMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"bridgeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callExecutor\",\"outputs\":[{\"internalType\":\"contractICallExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"counterPartMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cycleFee\",\"outputs\":[{\"internalType\":\"contractICycleFee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"deleteBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"estimateBridgeToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"formatAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onMessageReceived\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonZkEVMBridge\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkId\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setBrigeMessageWithValueFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"setCallExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_counterPartAddress\",\"type\":\"address\"}],\"name\":\"setCounterPart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setCycleFeeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"tokenInfoMap\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_Bridge *BridgeCaller) Blacklist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "blacklist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_Bridge *BridgeSession) Blacklist(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Blacklist(&_Bridge.CallOpts, arg0)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) Blacklist(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Blacklist(&_Bridge.CallOpts, arg0)
}

// BridgeMessageWithValueMap is a free data retrieval call binding the contract method 0x574c8611.
//
// Solidity: function bridgeMessageWithValueMap(uint32 ) view returns(bool)
func (_Bridge *BridgeCaller) BridgeMessageWithValueMap(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "bridgeMessageWithValueMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeMessageWithValueMap is a free data retrieval call binding the contract method 0x574c8611.
//
// Solidity: function bridgeMessageWithValueMap(uint32 ) view returns(bool)
func (_Bridge *BridgeSession) BridgeMessageWithValueMap(arg0 uint32) (bool, error) {
	return _Bridge.Contract.BridgeMessageWithValueMap(&_Bridge.CallOpts, arg0)
}

// BridgeMessageWithValueMap is a free data retrieval call binding the contract method 0x574c8611.
//
// Solidity: function bridgeMessageWithValueMap(uint32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) BridgeMessageWithValueMap(arg0 uint32) (bool, error) {
	return _Bridge.Contract.BridgeMessageWithValueMap(&_Bridge.CallOpts, arg0)
}

// CallExecutor is a free data retrieval call binding the contract method 0x990969f6.
//
// Solidity: function callExecutor() view returns(address)
func (_Bridge *BridgeCaller) CallExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "callExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallExecutor is a free data retrieval call binding the contract method 0x990969f6.
//
// Solidity: function callExecutor() view returns(address)
func (_Bridge *BridgeSession) CallExecutor() (common.Address, error) {
	return _Bridge.Contract.CallExecutor(&_Bridge.CallOpts)
}

// CallExecutor is a free data retrieval call binding the contract method 0x990969f6.
//
// Solidity: function callExecutor() view returns(address)
func (_Bridge *BridgeCallerSession) CallExecutor() (common.Address, error) {
	return _Bridge.Contract.CallExecutor(&_Bridge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Bridge *BridgeCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Bridge *BridgeSession) ChainId() (*big.Int, error) {
	return _Bridge.Contract.ChainId(&_Bridge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Bridge *BridgeCallerSession) ChainId() (*big.Int, error) {
	return _Bridge.Contract.ChainId(&_Bridge.CallOpts)
}

// CounterPartMap is a free data retrieval call binding the contract method 0x38d4de2a.
//
// Solidity: function counterPartMap(uint32 ) view returns(address)
func (_Bridge *BridgeCaller) CounterPartMap(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "counterPartMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CounterPartMap is a free data retrieval call binding the contract method 0x38d4de2a.
//
// Solidity: function counterPartMap(uint32 ) view returns(address)
func (_Bridge *BridgeSession) CounterPartMap(arg0 uint32) (common.Address, error) {
	return _Bridge.Contract.CounterPartMap(&_Bridge.CallOpts, arg0)
}

// CounterPartMap is a free data retrieval call binding the contract method 0x38d4de2a.
//
// Solidity: function counterPartMap(uint32 ) view returns(address)
func (_Bridge *BridgeCallerSession) CounterPartMap(arg0 uint32) (common.Address, error) {
	return _Bridge.Contract.CounterPartMap(&_Bridge.CallOpts, arg0)
}

// CycleFee is a free data retrieval call binding the contract method 0x31b9bab3.
//
// Solidity: function cycleFee() view returns(address)
func (_Bridge *BridgeCaller) CycleFee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "cycleFee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CycleFee is a free data retrieval call binding the contract method 0x31b9bab3.
//
// Solidity: function cycleFee() view returns(address)
func (_Bridge *BridgeSession) CycleFee() (common.Address, error) {
	return _Bridge.Contract.CycleFee(&_Bridge.CallOpts)
}

// CycleFee is a free data retrieval call binding the contract method 0x31b9bab3.
//
// Solidity: function cycleFee() view returns(address)
func (_Bridge *BridgeCallerSession) CycleFee() (common.Address, error) {
	return _Bridge.Contract.CycleFee(&_Bridge.CallOpts)
}

// EstimateBridgeToken is a free data retrieval call binding the contract method 0x506ec62d.
//
// Solidity: function estimateBridgeToken(uint32 , address , uint32 , uint256 , bytes , uint256 ) view returns(uint256 totalFee, uint256 serviceFee)
func (_Bridge *BridgeCaller) EstimateBridgeToken(opts *bind.CallOpts, arg0 uint32, arg1 common.Address, arg2 uint32, arg3 *big.Int, arg4 []byte, arg5 *big.Int) (struct {
	TotalFee   *big.Int
	ServiceFee *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "estimateBridgeToken", arg0, arg1, arg2, arg3, arg4, arg5)

	outstruct := new(struct {
		TotalFee   *big.Int
		ServiceFee *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ServiceFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateBridgeToken is a free data retrieval call binding the contract method 0x506ec62d.
//
// Solidity: function estimateBridgeToken(uint32 , address , uint32 , uint256 , bytes , uint256 ) view returns(uint256 totalFee, uint256 serviceFee)
func (_Bridge *BridgeSession) EstimateBridgeToken(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 *big.Int, arg4 []byte, arg5 *big.Int) (struct {
	TotalFee   *big.Int
	ServiceFee *big.Int
}, error) {
	return _Bridge.Contract.EstimateBridgeToken(&_Bridge.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// EstimateBridgeToken is a free data retrieval call binding the contract method 0x506ec62d.
//
// Solidity: function estimateBridgeToken(uint32 , address , uint32 , uint256 , bytes , uint256 ) view returns(uint256 totalFee, uint256 serviceFee)
func (_Bridge *BridgeCallerSession) EstimateBridgeToken(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 *big.Int, arg4 []byte, arg5 *big.Int) (struct {
	TotalFee   *big.Int
	ServiceFee *big.Int
}, error) {
	return _Bridge.Contract.EstimateBridgeToken(&_Bridge.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Bridge *BridgeCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Bridge *BridgeSession) FeeRecipient() (common.Address, error) {
	return _Bridge.Contract.FeeRecipient(&_Bridge.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Bridge *BridgeCallerSession) FeeRecipient() (common.Address, error) {
	return _Bridge.Contract.FeeRecipient(&_Bridge.CallOpts)
}

// FormatAmount is a free data retrieval call binding the contract method 0xd2e38a2f.
//
// Solidity: function formatAmount(uint32 tokenId, uint256 amount) view returns(uint256)
func (_Bridge *BridgeCaller) FormatAmount(opts *bind.CallOpts, tokenId uint32, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "formatAmount", tokenId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FormatAmount is a free data retrieval call binding the contract method 0xd2e38a2f.
//
// Solidity: function formatAmount(uint32 tokenId, uint256 amount) view returns(uint256)
func (_Bridge *BridgeSession) FormatAmount(tokenId uint32, amount *big.Int) (*big.Int, error) {
	return _Bridge.Contract.FormatAmount(&_Bridge.CallOpts, tokenId, amount)
}

// FormatAmount is a free data retrieval call binding the contract method 0xd2e38a2f.
//
// Solidity: function formatAmount(uint32 tokenId, uint256 amount) view returns(uint256)
func (_Bridge *BridgeCallerSession) FormatAmount(tokenId uint32, amount *big.Int) (*big.Int, error) {
	return _Bridge.Contract.FormatAmount(&_Bridge.CallOpts, tokenId, amount)
}

// GetGlobalId is a free data retrieval call binding the contract method 0xbef33731.
//
// Solidity: function getGlobalId() view returns(uint256)
func (_Bridge *BridgeCaller) GetGlobalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getGlobalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalId is a free data retrieval call binding the contract method 0xbef33731.
//
// Solidity: function getGlobalId() view returns(uint256)
func (_Bridge *BridgeSession) GetGlobalId() (*big.Int, error) {
	return _Bridge.Contract.GetGlobalId(&_Bridge.CallOpts)
}

// GetGlobalId is a free data retrieval call binding the contract method 0xbef33731.
//
// Solidity: function getGlobalId() view returns(uint256)
func (_Bridge *BridgeCallerSession) GetGlobalId() (*big.Int, error) {
	return _Bridge.Contract.GetGlobalId(&_Bridge.CallOpts)
}

// NativeTokenExist is a free data retrieval call binding the contract method 0x28822fc6.
//
// Solidity: function nativeTokenExist() view returns(bool)
func (_Bridge *BridgeCaller) NativeTokenExist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nativeTokenExist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NativeTokenExist is a free data retrieval call binding the contract method 0x28822fc6.
//
// Solidity: function nativeTokenExist() view returns(bool)
func (_Bridge *BridgeSession) NativeTokenExist() (bool, error) {
	return _Bridge.Contract.NativeTokenExist(&_Bridge.CallOpts)
}

// NativeTokenExist is a free data retrieval call binding the contract method 0x28822fc6.
//
// Solidity: function nativeTokenExist() view returns(bool)
func (_Bridge *BridgeCallerSession) NativeTokenExist() (bool, error) {
	return _Bridge.Contract.NativeTokenExist(&_Bridge.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridge *BridgeCaller) NetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "networkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridge *BridgeSession) NetworkID() (uint32, error) {
	return _Bridge.Contract.NetworkID(&_Bridge.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridge *BridgeCallerSession) NetworkID() (uint32, error) {
	return _Bridge.Contract.NetworkID(&_Bridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeSession) Nonce() (*big.Int, error) {
	return _Bridge.Contract.Nonce(&_Bridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeCallerSession) Nonce() (*big.Int, error) {
	return _Bridge.Contract.Nonce(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// PolygonZkEVMBridge is a free data retrieval call binding the contract method 0x5d43792c.
//
// Solidity: function polygonZkEVMBridge() view returns(address)
func (_Bridge *BridgeCaller) PolygonZkEVMBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "polygonZkEVMBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonZkEVMBridge is a free data retrieval call binding the contract method 0x5d43792c.
//
// Solidity: function polygonZkEVMBridge() view returns(address)
func (_Bridge *BridgeSession) PolygonZkEVMBridge() (common.Address, error) {
	return _Bridge.Contract.PolygonZkEVMBridge(&_Bridge.CallOpts)
}

// PolygonZkEVMBridge is a free data retrieval call binding the contract method 0x5d43792c.
//
// Solidity: function polygonZkEVMBridge() view returns(address)
func (_Bridge *BridgeCallerSession) PolygonZkEVMBridge() (common.Address, error) {
	return _Bridge.Contract.PolygonZkEVMBridge(&_Bridge.CallOpts)
}

// TokenInfoMap is a free data retrieval call binding the contract method 0x644f83ab.
//
// Solidity: function tokenInfoMap(uint32 ) view returns(uint32 flags, address tokenAddress)
func (_Bridge *BridgeCaller) TokenInfoMap(opts *bind.CallOpts, arg0 uint32) (struct {
	Flags        uint32
	TokenAddress common.Address
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tokenInfoMap", arg0)

	outstruct := new(struct {
		Flags        uint32
		TokenAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Flags = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// TokenInfoMap is a free data retrieval call binding the contract method 0x644f83ab.
//
// Solidity: function tokenInfoMap(uint32 ) view returns(uint32 flags, address tokenAddress)
func (_Bridge *BridgeSession) TokenInfoMap(arg0 uint32) (struct {
	Flags        uint32
	TokenAddress common.Address
}, error) {
	return _Bridge.Contract.TokenInfoMap(&_Bridge.CallOpts, arg0)
}

// TokenInfoMap is a free data retrieval call binding the contract method 0x644f83ab.
//
// Solidity: function tokenInfoMap(uint32 ) view returns(uint32 flags, address tokenAddress)
func (_Bridge *BridgeCallerSession) TokenInfoMap(arg0 uint32) (struct {
	Flags        uint32
	TokenAddress common.Address
}, error) {
	return _Bridge.Contract.TokenInfoMap(&_Bridge.CallOpts, arg0)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x3d2cc56c.
//
// Solidity: function addBlacklist(address[] accounts) returns()
func (_Bridge *BridgeTransactor) AddBlacklist(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addBlacklist", accounts)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x3d2cc56c.
//
// Solidity: function addBlacklist(address[] accounts) returns()
func (_Bridge *BridgeSession) AddBlacklist(accounts []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddBlacklist(&_Bridge.TransactOpts, accounts)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x3d2cc56c.
//
// Solidity: function addBlacklist(address[] accounts) returns()
func (_Bridge *BridgeTransactorSession) AddBlacklist(accounts []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddBlacklist(&_Bridge.TransactOpts, accounts)
}

// AddExistToken is a paid mutator transaction binding the contract method 0xb876d80d.
//
// Solidity: function addExistToken(uint32 tokenId, address token, uint8 sharedDecimals) returns()
func (_Bridge *BridgeTransactor) AddExistToken(opts *bind.TransactOpts, tokenId uint32, token common.Address, sharedDecimals uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addExistToken", tokenId, token, sharedDecimals)
}

// AddExistToken is a paid mutator transaction binding the contract method 0xb876d80d.
//
// Solidity: function addExistToken(uint32 tokenId, address token, uint8 sharedDecimals) returns()
func (_Bridge *BridgeSession) AddExistToken(tokenId uint32, token common.Address, sharedDecimals uint8) (*types.Transaction, error) {
	return _Bridge.Contract.AddExistToken(&_Bridge.TransactOpts, tokenId, token, sharedDecimals)
}

// AddExistToken is a paid mutator transaction binding the contract method 0xb876d80d.
//
// Solidity: function addExistToken(uint32 tokenId, address token, uint8 sharedDecimals) returns()
func (_Bridge *BridgeTransactorSession) AddExistToken(tokenId uint32, token common.Address, sharedDecimals uint8) (*types.Transaction, error) {
	return _Bridge.Contract.AddExistToken(&_Bridge.TransactOpts, tokenId, token, sharedDecimals)
}

// AddExistTokenBatch is a paid mutator transaction binding the contract method 0x1a535b0c.
//
// Solidity: function addExistTokenBatch((uint32,address,uint8)[] params) returns()
func (_Bridge *BridgeTransactor) AddExistTokenBatch(opts *bind.TransactOpts, params []CycleAssetAddExistTokenBatchParam) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addExistTokenBatch", params)
}

// AddExistTokenBatch is a paid mutator transaction binding the contract method 0x1a535b0c.
//
// Solidity: function addExistTokenBatch((uint32,address,uint8)[] params) returns()
func (_Bridge *BridgeSession) AddExistTokenBatch(params []CycleAssetAddExistTokenBatchParam) (*types.Transaction, error) {
	return _Bridge.Contract.AddExistTokenBatch(&_Bridge.TransactOpts, params)
}

// AddExistTokenBatch is a paid mutator transaction binding the contract method 0x1a535b0c.
//
// Solidity: function addExistTokenBatch((uint32,address,uint8)[] params) returns()
func (_Bridge *BridgeTransactorSession) AddExistTokenBatch(params []CycleAssetAddExistTokenBatchParam) (*types.Transaction, error) {
	return _Bridge.Contract.AddExistTokenBatch(&_Bridge.TransactOpts, params)
}

// AddNativeToken is a paid mutator transaction binding the contract method 0x790060ef.
//
// Solidity: function addNativeToken(uint32 tokenId, string name, string symbol, uint8 decimals) returns()
func (_Bridge *BridgeTransactor) AddNativeToken(opts *bind.TransactOpts, tokenId uint32, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addNativeToken", tokenId, name, symbol, decimals)
}

// AddNativeToken is a paid mutator transaction binding the contract method 0x790060ef.
//
// Solidity: function addNativeToken(uint32 tokenId, string name, string symbol, uint8 decimals) returns()
func (_Bridge *BridgeSession) AddNativeToken(tokenId uint32, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Bridge.Contract.AddNativeToken(&_Bridge.TransactOpts, tokenId, name, symbol, decimals)
}

// AddNativeToken is a paid mutator transaction binding the contract method 0x790060ef.
//
// Solidity: function addNativeToken(uint32 tokenId, string name, string symbol, uint8 decimals) returns()
func (_Bridge *BridgeTransactorSession) AddNativeToken(tokenId uint32, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Bridge.Contract.AddNativeToken(&_Bridge.TransactOpts, tokenId, name, symbol, decimals)
}

// AddWrappedToken is a paid mutator transaction binding the contract method 0x0b7a083a.
//
// Solidity: function addWrappedToken(uint32 tokenId, string name, string symbol, uint8 localDecimals, uint8 sharedDecimals) returns()
func (_Bridge *BridgeTransactor) AddWrappedToken(opts *bind.TransactOpts, tokenId uint32, name string, symbol string, localDecimals uint8, sharedDecimals uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addWrappedToken", tokenId, name, symbol, localDecimals, sharedDecimals)
}

// AddWrappedToken is a paid mutator transaction binding the contract method 0x0b7a083a.
//
// Solidity: function addWrappedToken(uint32 tokenId, string name, string symbol, uint8 localDecimals, uint8 sharedDecimals) returns()
func (_Bridge *BridgeSession) AddWrappedToken(tokenId uint32, name string, symbol string, localDecimals uint8, sharedDecimals uint8) (*types.Transaction, error) {
	return _Bridge.Contract.AddWrappedToken(&_Bridge.TransactOpts, tokenId, name, symbol, localDecimals, sharedDecimals)
}

// AddWrappedToken is a paid mutator transaction binding the contract method 0x0b7a083a.
//
// Solidity: function addWrappedToken(uint32 tokenId, string name, string symbol, uint8 localDecimals, uint8 sharedDecimals) returns()
func (_Bridge *BridgeTransactorSession) AddWrappedToken(tokenId uint32, name string, symbol string, localDecimals uint8, sharedDecimals uint8) (*types.Transaction, error) {
	return _Bridge.Contract.AddWrappedToken(&_Bridge.TransactOpts, tokenId, name, symbol, localDecimals, sharedDecimals)
}

// BridgeToken is a paid mutator transaction binding the contract method 0xf1b7f829.
//
// Solidity: function bridgeToken(uint32 destinationNetwork, address destinationAddress, uint32 tokenId, uint256 amount, bytes callData, uint256 feeAmount) payable returns()
func (_Bridge *BridgeTransactor) BridgeToken(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, tokenId uint32, amount *big.Int, callData []byte, feeAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "bridgeToken", destinationNetwork, destinationAddress, tokenId, amount, callData, feeAmount)
}

// BridgeToken is a paid mutator transaction binding the contract method 0xf1b7f829.
//
// Solidity: function bridgeToken(uint32 destinationNetwork, address destinationAddress, uint32 tokenId, uint256 amount, bytes callData, uint256 feeAmount) payable returns()
func (_Bridge *BridgeSession) BridgeToken(destinationNetwork uint32, destinationAddress common.Address, tokenId uint32, amount *big.Int, callData []byte, feeAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeToken(&_Bridge.TransactOpts, destinationNetwork, destinationAddress, tokenId, amount, callData, feeAmount)
}

// BridgeToken is a paid mutator transaction binding the contract method 0xf1b7f829.
//
// Solidity: function bridgeToken(uint32 destinationNetwork, address destinationAddress, uint32 tokenId, uint256 amount, bytes callData, uint256 feeAmount) payable returns()
func (_Bridge *BridgeTransactorSession) BridgeToken(destinationNetwork uint32, destinationAddress common.Address, tokenId uint32, amount *big.Int, callData []byte, feeAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeToken(&_Bridge.TransactOpts, destinationNetwork, destinationAddress, tokenId, amount, callData, feeAmount)
}

// DeleteBlacklist is a paid mutator transaction binding the contract method 0x3dc8b59f.
//
// Solidity: function deleteBlacklist(address[] accounts) returns()
func (_Bridge *BridgeTransactor) DeleteBlacklist(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deleteBlacklist", accounts)
}

// DeleteBlacklist is a paid mutator transaction binding the contract method 0x3dc8b59f.
//
// Solidity: function deleteBlacklist(address[] accounts) returns()
func (_Bridge *BridgeSession) DeleteBlacklist(accounts []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.DeleteBlacklist(&_Bridge.TransactOpts, accounts)
}

// DeleteBlacklist is a paid mutator transaction binding the contract method 0x3dc8b59f.
//
// Solidity: function deleteBlacklist(address[] accounts) returns()
func (_Bridge *BridgeTransactorSession) DeleteBlacklist(accounts []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.DeleteBlacklist(&_Bridge.TransactOpts, accounts)
}

// OnMessageReceived is a paid mutator transaction binding the contract method 0x1806b5f2.
//
// Solidity: function onMessageReceived(address originAddress, uint32 originNetwork, bytes data) payable returns()
func (_Bridge *BridgeTransactor) OnMessageReceived(opts *bind.TransactOpts, originAddress common.Address, originNetwork uint32, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onMessageReceived", originAddress, originNetwork, data)
}

// OnMessageReceived is a paid mutator transaction binding the contract method 0x1806b5f2.
//
// Solidity: function onMessageReceived(address originAddress, uint32 originNetwork, bytes data) payable returns()
func (_Bridge *BridgeSession) OnMessageReceived(originAddress common.Address, originNetwork uint32, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnMessageReceived(&_Bridge.TransactOpts, originAddress, originNetwork, data)
}

// OnMessageReceived is a paid mutator transaction binding the contract method 0x1806b5f2.
//
// Solidity: function onMessageReceived(address originAddress, uint32 originNetwork, bytes data) payable returns()
func (_Bridge *BridgeTransactorSession) OnMessageReceived(originAddress common.Address, originNetwork uint32, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnMessageReceived(&_Bridge.TransactOpts, originAddress, originNetwork, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// SetBrigeMessageWithValueFlag is a paid mutator transaction binding the contract method 0x9ece28b2.
//
// Solidity: function setBrigeMessageWithValueFlag(uint32 networkId, bool enable) returns()
func (_Bridge *BridgeTransactor) SetBrigeMessageWithValueFlag(opts *bind.TransactOpts, networkId uint32, enable bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setBrigeMessageWithValueFlag", networkId, enable)
}

// SetBrigeMessageWithValueFlag is a paid mutator transaction binding the contract method 0x9ece28b2.
//
// Solidity: function setBrigeMessageWithValueFlag(uint32 networkId, bool enable) returns()
func (_Bridge *BridgeSession) SetBrigeMessageWithValueFlag(networkId uint32, enable bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetBrigeMessageWithValueFlag(&_Bridge.TransactOpts, networkId, enable)
}

// SetBrigeMessageWithValueFlag is a paid mutator transaction binding the contract method 0x9ece28b2.
//
// Solidity: function setBrigeMessageWithValueFlag(uint32 networkId, bool enable) returns()
func (_Bridge *BridgeTransactorSession) SetBrigeMessageWithValueFlag(networkId uint32, enable bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetBrigeMessageWithValueFlag(&_Bridge.TransactOpts, networkId, enable)
}

// SetCallExecutor is a paid mutator transaction binding the contract method 0x4b1a85f9.
//
// Solidity: function setCallExecutor(address executor) returns()
func (_Bridge *BridgeTransactor) SetCallExecutor(opts *bind.TransactOpts, executor common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setCallExecutor", executor)
}

// SetCallExecutor is a paid mutator transaction binding the contract method 0x4b1a85f9.
//
// Solidity: function setCallExecutor(address executor) returns()
func (_Bridge *BridgeSession) SetCallExecutor(executor common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetCallExecutor(&_Bridge.TransactOpts, executor)
}

// SetCallExecutor is a paid mutator transaction binding the contract method 0x4b1a85f9.
//
// Solidity: function setCallExecutor(address executor) returns()
func (_Bridge *BridgeTransactorSession) SetCallExecutor(executor common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetCallExecutor(&_Bridge.TransactOpts, executor)
}

// SetCounterPart is a paid mutator transaction binding the contract method 0xa240799a.
//
// Solidity: function setCounterPart(uint32 _networkId, address _counterPartAddress) returns()
func (_Bridge *BridgeTransactor) SetCounterPart(opts *bind.TransactOpts, _networkId uint32, _counterPartAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setCounterPart", _networkId, _counterPartAddress)
}

// SetCounterPart is a paid mutator transaction binding the contract method 0xa240799a.
//
// Solidity: function setCounterPart(uint32 _networkId, address _counterPartAddress) returns()
func (_Bridge *BridgeSession) SetCounterPart(_networkId uint32, _counterPartAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetCounterPart(&_Bridge.TransactOpts, _networkId, _counterPartAddress)
}

// SetCounterPart is a paid mutator transaction binding the contract method 0xa240799a.
//
// Solidity: function setCounterPart(uint32 _networkId, address _counterPartAddress) returns()
func (_Bridge *BridgeTransactorSession) SetCounterPart(_networkId uint32, _counterPartAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetCounterPart(&_Bridge.TransactOpts, _networkId, _counterPartAddress)
}

// SetCycleFeeContract is a paid mutator transaction binding the contract method 0x155647c0.
//
// Solidity: function setCycleFeeContract(address contractAddress) returns()
func (_Bridge *BridgeTransactor) SetCycleFeeContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setCycleFeeContract", contractAddress)
}

// SetCycleFeeContract is a paid mutator transaction binding the contract method 0x155647c0.
//
// Solidity: function setCycleFeeContract(address contractAddress) returns()
func (_Bridge *BridgeSession) SetCycleFeeContract(contractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetCycleFeeContract(&_Bridge.TransactOpts, contractAddress)
}

// SetCycleFeeContract is a paid mutator transaction binding the contract method 0x155647c0.
//
// Solidity: function setCycleFeeContract(address contractAddress) returns()
func (_Bridge *BridgeTransactorSession) SetCycleFeeContract(contractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetCycleFeeContract(&_Bridge.TransactOpts, contractAddress)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_Bridge *BridgeTransactor) SetFeeRecipient(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setFeeRecipient", recipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_Bridge *BridgeSession) SetFeeRecipient(recipient common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetFeeRecipient(&_Bridge.TransactOpts, recipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_Bridge *BridgeTransactorSession) SetFeeRecipient(recipient common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetFeeRecipient(&_Bridge.TransactOpts, recipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// BridgeAddBlacklistIterator is returned from FilterAddBlacklist and is used to iterate over the raw logs and unpacked data for AddBlacklist events raised by the Bridge contract.
type BridgeAddBlacklistIterator struct {
	Event *BridgeAddBlacklist // Event containing the contract specifics and raw log

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
func (it *BridgeAddBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAddBlacklist)
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
		it.Event = new(BridgeAddBlacklist)
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
func (it *BridgeAddBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAddBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAddBlacklist represents a AddBlacklist event raised by the Bridge contract.
type BridgeAddBlacklist struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddBlacklist is a free log retrieval operation binding the contract event 0x7e239e822fa537514cc6b38d8350bde5ce06a8f9282c77161b926fc077a81026.
//
// Solidity: event AddBlacklist(address account)
func (_Bridge *BridgeFilterer) FilterAddBlacklist(opts *bind.FilterOpts) (*BridgeAddBlacklistIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "AddBlacklist")
	if err != nil {
		return nil, err
	}
	return &BridgeAddBlacklistIterator{contract: _Bridge.contract, event: "AddBlacklist", logs: logs, sub: sub}, nil
}

// WatchAddBlacklist is a free log subscription operation binding the contract event 0x7e239e822fa537514cc6b38d8350bde5ce06a8f9282c77161b926fc077a81026.
//
// Solidity: event AddBlacklist(address account)
func (_Bridge *BridgeFilterer) WatchAddBlacklist(opts *bind.WatchOpts, sink chan<- *BridgeAddBlacklist) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "AddBlacklist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAddBlacklist)
				if err := _Bridge.contract.UnpackLog(event, "AddBlacklist", log); err != nil {
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

// ParseAddBlacklist is a log parse operation binding the contract event 0x7e239e822fa537514cc6b38d8350bde5ce06a8f9282c77161b926fc077a81026.
//
// Solidity: event AddBlacklist(address account)
func (_Bridge *BridgeFilterer) ParseAddBlacklist(log types.Log) (*BridgeAddBlacklist, error) {
	event := new(BridgeAddBlacklist)
	if err := _Bridge.contract.UnpackLog(event, "AddBlacklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeAddTokenIterator is returned from FilterAddToken and is used to iterate over the raw logs and unpacked data for AddToken events raised by the Bridge contract.
type BridgeAddTokenIterator struct {
	Event *BridgeAddToken // Event containing the contract specifics and raw log

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
func (it *BridgeAddTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAddToken)
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
		it.Event = new(BridgeAddToken)
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
func (it *BridgeAddTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAddTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAddToken represents a AddToken event raised by the Bridge contract.
type BridgeAddToken struct {
	TokenId        uint32
	Flags          uint32
	Name           string
	Symbol         string
	SharedDecimals uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddToken is a free log retrieval operation binding the contract event 0xa8778e48bd9593778650efe827160de5d410a2c3167a6471077b3d43502073b0.
//
// Solidity: event AddToken(uint32 tokenId, uint32 flags, string name, string symbol, uint8 sharedDecimals)
func (_Bridge *BridgeFilterer) FilterAddToken(opts *bind.FilterOpts) (*BridgeAddTokenIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "AddToken")
	if err != nil {
		return nil, err
	}
	return &BridgeAddTokenIterator{contract: _Bridge.contract, event: "AddToken", logs: logs, sub: sub}, nil
}

// WatchAddToken is a free log subscription operation binding the contract event 0xa8778e48bd9593778650efe827160de5d410a2c3167a6471077b3d43502073b0.
//
// Solidity: event AddToken(uint32 tokenId, uint32 flags, string name, string symbol, uint8 sharedDecimals)
func (_Bridge *BridgeFilterer) WatchAddToken(opts *bind.WatchOpts, sink chan<- *BridgeAddToken) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "AddToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAddToken)
				if err := _Bridge.contract.UnpackLog(event, "AddToken", log); err != nil {
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

// ParseAddToken is a log parse operation binding the contract event 0xa8778e48bd9593778650efe827160de5d410a2c3167a6471077b3d43502073b0.
//
// Solidity: event AddToken(uint32 tokenId, uint32 flags, string name, string symbol, uint8 sharedDecimals)
func (_Bridge *BridgeFilterer) ParseAddToken(log types.Log) (*BridgeAddToken, error) {
	event := new(BridgeAddToken)
	if err := _Bridge.contract.UnpackLog(event, "AddToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBridgeIterator is returned from FilterBridge and is used to iterate over the raw logs and unpacked data for Bridge events raised by the Bridge contract.
type BridgeBridgeIterator struct {
	Event *BridgeBridge // Event containing the contract specifics and raw log

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
func (it *BridgeBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBridge)
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
		it.Event = new(BridgeBridge)
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
func (it *BridgeBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBridge represents a Bridge event raised by the Bridge contract.
type BridgeBridge struct {
	GlobalId           *big.Int
	DestinationNetwork uint32
	DestinationAddress common.Address
	DestinationFunc    [4]byte
	TokenId            uint32
	SdAmount           *big.Int
	SharedDecimals     uint8
	FeeAmount          *big.Int
	FeeRecipient       common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBridge is a free log retrieval operation binding the contract event 0xa352ebf5b2deda290503f4ca11d1d10893bbe0919ebe4ecf9bf2e25e404b1b51.
//
// Solidity: event Bridge(uint256 globalId, uint32 destinationNetwork, address destinationAddress, bytes4 destinationFunc, uint32 tokenId, uint256 sdAmount, uint8 sharedDecimals, uint256 feeAmount, address feeRecipient)
func (_Bridge *BridgeFilterer) FilterBridge(opts *bind.FilterOpts) (*BridgeBridgeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Bridge")
	if err != nil {
		return nil, err
	}
	return &BridgeBridgeIterator{contract: _Bridge.contract, event: "Bridge", logs: logs, sub: sub}, nil
}

// WatchBridge is a free log subscription operation binding the contract event 0xa352ebf5b2deda290503f4ca11d1d10893bbe0919ebe4ecf9bf2e25e404b1b51.
//
// Solidity: event Bridge(uint256 globalId, uint32 destinationNetwork, address destinationAddress, bytes4 destinationFunc, uint32 tokenId, uint256 sdAmount, uint8 sharedDecimals, uint256 feeAmount, address feeRecipient)
func (_Bridge *BridgeFilterer) WatchBridge(opts *bind.WatchOpts, sink chan<- *BridgeBridge) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Bridge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBridge)
				if err := _Bridge.contract.UnpackLog(event, "Bridge", log); err != nil {
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

// ParseBridge is a log parse operation binding the contract event 0xa352ebf5b2deda290503f4ca11d1d10893bbe0919ebe4ecf9bf2e25e404b1b51.
//
// Solidity: event Bridge(uint256 globalId, uint32 destinationNetwork, address destinationAddress, bytes4 destinationFunc, uint32 tokenId, uint256 sdAmount, uint8 sharedDecimals, uint256 feeAmount, address feeRecipient)
func (_Bridge *BridgeFilterer) ParseBridge(log types.Log) (*BridgeBridge, error) {
	event := new(BridgeBridge)
	if err := _Bridge.contract.UnpackLog(event, "Bridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Bridge contract.
type BridgeClaimIterator struct {
	Event *BridgeClaim // Event containing the contract specifics and raw log

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
func (it *BridgeClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeClaim)
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
		it.Event = new(BridgeClaim)
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
func (it *BridgeClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeClaim represents a Claim event raised by the Bridge contract.
type BridgeClaim struct {
	GlobalId           *big.Int
	DestinationAddress common.Address
	DestinationFunc    [4]byte
	TokenId            uint32
	SdAmount           *big.Int
	SharedDecimals     uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0xb6397d9d5c14c0fcb90a5c764a69e0b67d770bdbd13ff18345bde4ce1d2437ac.
//
// Solidity: event Claim(uint256 globalId, address destinationAddress, bytes4 destinationFunc, uint32 tokenId, uint256 sdAmount, uint8 sharedDecimals)
func (_Bridge *BridgeFilterer) FilterClaim(opts *bind.FilterOpts) (*BridgeClaimIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &BridgeClaimIterator{contract: _Bridge.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0xb6397d9d5c14c0fcb90a5c764a69e0b67d770bdbd13ff18345bde4ce1d2437ac.
//
// Solidity: event Claim(uint256 globalId, address destinationAddress, bytes4 destinationFunc, uint32 tokenId, uint256 sdAmount, uint8 sharedDecimals)
func (_Bridge *BridgeFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *BridgeClaim) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeClaim)
				if err := _Bridge.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0xb6397d9d5c14c0fcb90a5c764a69e0b67d770bdbd13ff18345bde4ce1d2437ac.
//
// Solidity: event Claim(uint256 globalId, address destinationAddress, bytes4 destinationFunc, uint32 tokenId, uint256 sdAmount, uint8 sharedDecimals)
func (_Bridge *BridgeFilterer) ParseClaim(log types.Log) (*BridgeClaim, error) {
	event := new(BridgeClaim)
	if err := _Bridge.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDeleteBlacklistIterator is returned from FilterDeleteBlacklist and is used to iterate over the raw logs and unpacked data for DeleteBlacklist events raised by the Bridge contract.
type BridgeDeleteBlacklistIterator struct {
	Event *BridgeDeleteBlacklist // Event containing the contract specifics and raw log

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
func (it *BridgeDeleteBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDeleteBlacklist)
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
		it.Event = new(BridgeDeleteBlacklist)
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
func (it *BridgeDeleteBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDeleteBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDeleteBlacklist represents a DeleteBlacklist event raised by the Bridge contract.
type BridgeDeleteBlacklist struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeleteBlacklist is a free log retrieval operation binding the contract event 0xbceb396df10b0dfcd3897f25215ee21a8e8ca47dff0df47c827fbab30cff1339.
//
// Solidity: event DeleteBlacklist(address account)
func (_Bridge *BridgeFilterer) FilterDeleteBlacklist(opts *bind.FilterOpts) (*BridgeDeleteBlacklistIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DeleteBlacklist")
	if err != nil {
		return nil, err
	}
	return &BridgeDeleteBlacklistIterator{contract: _Bridge.contract, event: "DeleteBlacklist", logs: logs, sub: sub}, nil
}

// WatchDeleteBlacklist is a free log subscription operation binding the contract event 0xbceb396df10b0dfcd3897f25215ee21a8e8ca47dff0df47c827fbab30cff1339.
//
// Solidity: event DeleteBlacklist(address account)
func (_Bridge *BridgeFilterer) WatchDeleteBlacklist(opts *bind.WatchOpts, sink chan<- *BridgeDeleteBlacklist) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DeleteBlacklist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDeleteBlacklist)
				if err := _Bridge.contract.UnpackLog(event, "DeleteBlacklist", log); err != nil {
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

// ParseDeleteBlacklist is a log parse operation binding the contract event 0xbceb396df10b0dfcd3897f25215ee21a8e8ca47dff0df47c827fbab30cff1339.
//
// Solidity: event DeleteBlacklist(address account)
func (_Bridge *BridgeFilterer) ParseDeleteBlacklist(log types.Log) (*BridgeDeleteBlacklist, error) {
	event := new(BridgeDeleteBlacklist)
	if err := _Bridge.contract.UnpackLog(event, "DeleteBlacklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeForwardIterator is returned from FilterForward and is used to iterate over the raw logs and unpacked data for Forward events raised by the Bridge contract.
type BridgeForwardIterator struct {
	Event *BridgeForward // Event containing the contract specifics and raw log

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
func (it *BridgeForwardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeForward)
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
		it.Event = new(BridgeForward)
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
func (it *BridgeForwardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeForwardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeForward represents a Forward event raised by the Bridge contract.
type BridgeForward struct {
	GlobalId           *big.Int
	DestinationNetwork uint32
	DestinationAddress common.Address
	DestinationFunc    [4]byte
	TokenId            uint32
	SdAmount           *big.Int
	SharedDecimals     uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForward is a free log retrieval operation binding the contract event 0xbb7be7094fee86385ca75dca7ed2557bc23958eba940fa061889c1ce5135554f.
//
// Solidity: event Forward(uint256 globalId, uint32 destinationNetwork, address destinationAddress, bytes4 destinationFunc, uint32 tokenId, uint256 sdAmount, uint8 sharedDecimals)
func (_Bridge *BridgeFilterer) FilterForward(opts *bind.FilterOpts) (*BridgeForwardIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Forward")
	if err != nil {
		return nil, err
	}
	return &BridgeForwardIterator{contract: _Bridge.contract, event: "Forward", logs: logs, sub: sub}, nil
}

// WatchForward is a free log subscription operation binding the contract event 0xbb7be7094fee86385ca75dca7ed2557bc23958eba940fa061889c1ce5135554f.
//
// Solidity: event Forward(uint256 globalId, uint32 destinationNetwork, address destinationAddress, bytes4 destinationFunc, uint32 tokenId, uint256 sdAmount, uint8 sharedDecimals)
func (_Bridge *BridgeFilterer) WatchForward(opts *bind.WatchOpts, sink chan<- *BridgeForward) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Forward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeForward)
				if err := _Bridge.contract.UnpackLog(event, "Forward", log); err != nil {
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

// ParseForward is a log parse operation binding the contract event 0xbb7be7094fee86385ca75dca7ed2557bc23958eba940fa061889c1ce5135554f.
//
// Solidity: event Forward(uint256 globalId, uint32 destinationNetwork, address destinationAddress, bytes4 destinationFunc, uint32 tokenId, uint256 sdAmount, uint8 sharedDecimals)
func (_Bridge *BridgeFilterer) ParseForward(log types.Log) (*BridgeForward, error) {
	event := new(BridgeForward)
	if err := _Bridge.contract.UnpackLog(event, "Forward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
