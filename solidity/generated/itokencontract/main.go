// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itokencontract

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

// ItokencontractMetaData contains all meta data concerning the Itokencontract contract.
var ItokencontractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaidTokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenPrice\",\"type\":\"uint256\"}],\"name\":\"PaymentSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"name\":\"TokenContractParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"TokenMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenFactoryAddr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken_\",\"type\":\"uint256\"}],\"name\":\"__TokenContract_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"}],\"name\":\"existingTokenURIs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"}],\"name\":\"getUserTokenIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"}],\"name\":\"mintToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"contractITokenFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newTokenName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newTokenSymbol_\",\"type\":\"string\"}],\"name\":\"updateTokenContractParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"}],\"name\":\"withdrawPaidTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ItokencontractABI is the input ABI used to generate the binding from.
// Deprecated: Use ItokencontractMetaData.ABI instead.
var ItokencontractABI = ItokencontractMetaData.ABI

// Itokencontract is an auto generated Go binding around an Ethereum contract.
type Itokencontract struct {
	ItokencontractCaller     // Read-only binding to the contract
	ItokencontractTransactor // Write-only binding to the contract
	ItokencontractFilterer   // Log filterer for contract events
}

// ItokencontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ItokencontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokencontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ItokencontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokencontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ItokencontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokencontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ItokencontractSession struct {
	Contract     *Itokencontract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ItokencontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ItokencontractCallerSession struct {
	Contract *ItokencontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ItokencontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ItokencontractTransactorSession struct {
	Contract     *ItokencontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ItokencontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ItokencontractRaw struct {
	Contract *Itokencontract // Generic contract binding to access the raw methods on
}

// ItokencontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ItokencontractCallerRaw struct {
	Contract *ItokencontractCaller // Generic read-only contract binding to access the raw methods on
}

// ItokencontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ItokencontractTransactorRaw struct {
	Contract *ItokencontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewItokencontract creates a new instance of Itokencontract, bound to a specific deployed contract.
func NewItokencontract(address common.Address, backend bind.ContractBackend) (*Itokencontract, error) {
	contract, err := bindItokencontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Itokencontract{ItokencontractCaller: ItokencontractCaller{contract: contract}, ItokencontractTransactor: ItokencontractTransactor{contract: contract}, ItokencontractFilterer: ItokencontractFilterer{contract: contract}}, nil
}

// NewItokencontractCaller creates a new read-only instance of Itokencontract, bound to a specific deployed contract.
func NewItokencontractCaller(address common.Address, caller bind.ContractCaller) (*ItokencontractCaller, error) {
	contract, err := bindItokencontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ItokencontractCaller{contract: contract}, nil
}

// NewItokencontractTransactor creates a new write-only instance of Itokencontract, bound to a specific deployed contract.
func NewItokencontractTransactor(address common.Address, transactor bind.ContractTransactor) (*ItokencontractTransactor, error) {
	contract, err := bindItokencontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ItokencontractTransactor{contract: contract}, nil
}

// NewItokencontractFilterer creates a new log filterer instance of Itokencontract, bound to a specific deployed contract.
func NewItokencontractFilterer(address common.Address, filterer bind.ContractFilterer) (*ItokencontractFilterer, error) {
	contract, err := bindItokencontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ItokencontractFilterer{contract: contract}, nil
}

// bindItokencontract binds a generic wrapper to an already deployed contract.
func bindItokencontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ItokencontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itokencontract *ItokencontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Itokencontract.Contract.ItokencontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itokencontract *ItokencontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itokencontract.Contract.ItokencontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itokencontract *ItokencontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Itokencontract.Contract.ItokencontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itokencontract *ItokencontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Itokencontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itokencontract *ItokencontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itokencontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itokencontract *ItokencontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Itokencontract.Contract.contract.Transact(opts, method, params...)
}

// ExistingTokenURIs is a free data retrieval call binding the contract method 0xb8f4b821.
//
// Solidity: function existingTokenURIs(string tokenURI_) view returns(bool)
func (_Itokencontract *ItokencontractCaller) ExistingTokenURIs(opts *bind.CallOpts, tokenURI_ string) (bool, error) {
	var out []interface{}
	err := _Itokencontract.contract.Call(opts, &out, "existingTokenURIs", tokenURI_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistingTokenURIs is a free data retrieval call binding the contract method 0xb8f4b821.
//
// Solidity: function existingTokenURIs(string tokenURI_) view returns(bool)
func (_Itokencontract *ItokencontractSession) ExistingTokenURIs(tokenURI_ string) (bool, error) {
	return _Itokencontract.Contract.ExistingTokenURIs(&_Itokencontract.CallOpts, tokenURI_)
}

// ExistingTokenURIs is a free data retrieval call binding the contract method 0xb8f4b821.
//
// Solidity: function existingTokenURIs(string tokenURI_) view returns(bool)
func (_Itokencontract *ItokencontractCallerSession) ExistingTokenURIs(tokenURI_ string) (bool, error) {
	return _Itokencontract.Contract.ExistingTokenURIs(&_Itokencontract.CallOpts, tokenURI_)
}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0x595ee598.
//
// Solidity: function getUserTokenIDs(address userAddr_) view returns(uint256[] tokenIDs_)
func (_Itokencontract *ItokencontractCaller) GetUserTokenIDs(opts *bind.CallOpts, userAddr_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Itokencontract.contract.Call(opts, &out, "getUserTokenIDs", userAddr_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0x595ee598.
//
// Solidity: function getUserTokenIDs(address userAddr_) view returns(uint256[] tokenIDs_)
func (_Itokencontract *ItokencontractSession) GetUserTokenIDs(userAddr_ common.Address) ([]*big.Int, error) {
	return _Itokencontract.Contract.GetUserTokenIDs(&_Itokencontract.CallOpts, userAddr_)
}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0x595ee598.
//
// Solidity: function getUserTokenIDs(address userAddr_) view returns(uint256[] tokenIDs_)
func (_Itokencontract *ItokencontractCallerSession) GetUserTokenIDs(userAddr_ common.Address) ([]*big.Int, error) {
	return _Itokencontract.Contract.GetUserTokenIDs(&_Itokencontract.CallOpts, userAddr_)
}

// PricePerOneToken is a free data retrieval call binding the contract method 0x1935fba9.
//
// Solidity: function pricePerOneToken() view returns(uint256)
func (_Itokencontract *ItokencontractCaller) PricePerOneToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Itokencontract.contract.Call(opts, &out, "pricePerOneToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerOneToken is a free data retrieval call binding the contract method 0x1935fba9.
//
// Solidity: function pricePerOneToken() view returns(uint256)
func (_Itokencontract *ItokencontractSession) PricePerOneToken() (*big.Int, error) {
	return _Itokencontract.Contract.PricePerOneToken(&_Itokencontract.CallOpts)
}

// PricePerOneToken is a free data retrieval call binding the contract method 0x1935fba9.
//
// Solidity: function pricePerOneToken() view returns(uint256)
func (_Itokencontract *ItokencontractCallerSession) PricePerOneToken() (*big.Int, error) {
	return _Itokencontract.Contract.PricePerOneToken(&_Itokencontract.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Itokencontract *ItokencontractCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Itokencontract.contract.Call(opts, &out, "tokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Itokencontract *ItokencontractSession) TokenFactory() (common.Address, error) {
	return _Itokencontract.Contract.TokenFactory(&_Itokencontract.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Itokencontract *ItokencontractCallerSession) TokenFactory() (common.Address, error) {
	return _Itokencontract.Contract.TokenFactory(&_Itokencontract.CallOpts)
}

// TokenContractInit is a paid mutator transaction binding the contract method 0xdac51535.
//
// Solidity: function __TokenContract_init(string tokenName_, string tokenSymbol_, address tokenFactoryAddr_, uint256 pricePerOneToken_) returns()
func (_Itokencontract *ItokencontractTransactor) TokenContractInit(opts *bind.TransactOpts, tokenName_ string, tokenSymbol_ string, tokenFactoryAddr_ common.Address, pricePerOneToken_ *big.Int) (*types.Transaction, error) {
	return _Itokencontract.contract.Transact(opts, "__TokenContract_init", tokenName_, tokenSymbol_, tokenFactoryAddr_, pricePerOneToken_)
}

// TokenContractInit is a paid mutator transaction binding the contract method 0xdac51535.
//
// Solidity: function __TokenContract_init(string tokenName_, string tokenSymbol_, address tokenFactoryAddr_, uint256 pricePerOneToken_) returns()
func (_Itokencontract *ItokencontractSession) TokenContractInit(tokenName_ string, tokenSymbol_ string, tokenFactoryAddr_ common.Address, pricePerOneToken_ *big.Int) (*types.Transaction, error) {
	return _Itokencontract.Contract.TokenContractInit(&_Itokencontract.TransactOpts, tokenName_, tokenSymbol_, tokenFactoryAddr_, pricePerOneToken_)
}

// TokenContractInit is a paid mutator transaction binding the contract method 0xdac51535.
//
// Solidity: function __TokenContract_init(string tokenName_, string tokenSymbol_, address tokenFactoryAddr_, uint256 pricePerOneToken_) returns()
func (_Itokencontract *ItokencontractTransactorSession) TokenContractInit(tokenName_ string, tokenSymbol_ string, tokenFactoryAddr_ common.Address, pricePerOneToken_ *big.Int) (*types.Transaction, error) {
	return _Itokencontract.Contract.TokenContractInit(&_Itokencontract.TransactOpts, tokenName_, tokenSymbol_, tokenFactoryAddr_, pricePerOneToken_)
}

// MintToken is a paid mutator transaction binding the contract method 0x7cec18b3.
//
// Solidity: function mintToken(address paymentTokenAddress_, uint256 paymentTokenPrice_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) payable returns()
func (_Itokencontract *ItokencontractTransactor) MintToken(opts *bind.TransactOpts, paymentTokenAddress_ common.Address, paymentTokenPrice_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Itokencontract.contract.Transact(opts, "mintToken", paymentTokenAddress_, paymentTokenPrice_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// MintToken is a paid mutator transaction binding the contract method 0x7cec18b3.
//
// Solidity: function mintToken(address paymentTokenAddress_, uint256 paymentTokenPrice_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) payable returns()
func (_Itokencontract *ItokencontractSession) MintToken(paymentTokenAddress_ common.Address, paymentTokenPrice_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Itokencontract.Contract.MintToken(&_Itokencontract.TransactOpts, paymentTokenAddress_, paymentTokenPrice_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// MintToken is a paid mutator transaction binding the contract method 0x7cec18b3.
//
// Solidity: function mintToken(address paymentTokenAddress_, uint256 paymentTokenPrice_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) payable returns()
func (_Itokencontract *ItokencontractTransactorSession) MintToken(paymentTokenAddress_ common.Address, paymentTokenPrice_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Itokencontract.Contract.MintToken(&_Itokencontract.TransactOpts, paymentTokenAddress_, paymentTokenPrice_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Itokencontract *ItokencontractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itokencontract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Itokencontract *ItokencontractSession) Pause() (*types.Transaction, error) {
	return _Itokencontract.Contract.Pause(&_Itokencontract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Itokencontract *ItokencontractTransactorSession) Pause() (*types.Transaction, error) {
	return _Itokencontract.Contract.Pause(&_Itokencontract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Itokencontract *ItokencontractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itokencontract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Itokencontract *ItokencontractSession) Unpause() (*types.Transaction, error) {
	return _Itokencontract.Contract.Unpause(&_Itokencontract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Itokencontract *ItokencontractTransactorSession) Unpause() (*types.Transaction, error) {
	return _Itokencontract.Contract.Unpause(&_Itokencontract.TransactOpts)
}

// UpdateTokenContractParams is a paid mutator transaction binding the contract method 0x541a738c.
//
// Solidity: function updateTokenContractParams(uint256 newPrice_, string newTokenName_, string newTokenSymbol_) returns()
func (_Itokencontract *ItokencontractTransactor) UpdateTokenContractParams(opts *bind.TransactOpts, newPrice_ *big.Int, newTokenName_ string, newTokenSymbol_ string) (*types.Transaction, error) {
	return _Itokencontract.contract.Transact(opts, "updateTokenContractParams", newPrice_, newTokenName_, newTokenSymbol_)
}

// UpdateTokenContractParams is a paid mutator transaction binding the contract method 0x541a738c.
//
// Solidity: function updateTokenContractParams(uint256 newPrice_, string newTokenName_, string newTokenSymbol_) returns()
func (_Itokencontract *ItokencontractSession) UpdateTokenContractParams(newPrice_ *big.Int, newTokenName_ string, newTokenSymbol_ string) (*types.Transaction, error) {
	return _Itokencontract.Contract.UpdateTokenContractParams(&_Itokencontract.TransactOpts, newPrice_, newTokenName_, newTokenSymbol_)
}

// UpdateTokenContractParams is a paid mutator transaction binding the contract method 0x541a738c.
//
// Solidity: function updateTokenContractParams(uint256 newPrice_, string newTokenName_, string newTokenSymbol_) returns()
func (_Itokencontract *ItokencontractTransactorSession) UpdateTokenContractParams(newPrice_ *big.Int, newTokenName_ string, newTokenSymbol_ string) (*types.Transaction, error) {
	return _Itokencontract.Contract.UpdateTokenContractParams(&_Itokencontract.TransactOpts, newPrice_, newTokenName_, newTokenSymbol_)
}

// WithdrawPaidTokens is a paid mutator transaction binding the contract method 0xd3fd39aa.
//
// Solidity: function withdrawPaidTokens(address tokenAddr_, address recipient_) returns()
func (_Itokencontract *ItokencontractTransactor) WithdrawPaidTokens(opts *bind.TransactOpts, tokenAddr_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _Itokencontract.contract.Transact(opts, "withdrawPaidTokens", tokenAddr_, recipient_)
}

// WithdrawPaidTokens is a paid mutator transaction binding the contract method 0xd3fd39aa.
//
// Solidity: function withdrawPaidTokens(address tokenAddr_, address recipient_) returns()
func (_Itokencontract *ItokencontractSession) WithdrawPaidTokens(tokenAddr_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _Itokencontract.Contract.WithdrawPaidTokens(&_Itokencontract.TransactOpts, tokenAddr_, recipient_)
}

// WithdrawPaidTokens is a paid mutator transaction binding the contract method 0xd3fd39aa.
//
// Solidity: function withdrawPaidTokens(address tokenAddr_, address recipient_) returns()
func (_Itokencontract *ItokencontractTransactorSession) WithdrawPaidTokens(tokenAddr_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _Itokencontract.Contract.WithdrawPaidTokens(&_Itokencontract.TransactOpts, tokenAddr_, recipient_)
}

// ItokencontractPaidTokensWithdrawnIterator is returned from FilterPaidTokensWithdrawn and is used to iterate over the raw logs and unpacked data for PaidTokensWithdrawn events raised by the Itokencontract contract.
type ItokencontractPaidTokensWithdrawnIterator struct {
	Event *ItokencontractPaidTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *ItokencontractPaidTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItokencontractPaidTokensWithdrawn)
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
		it.Event = new(ItokencontractPaidTokensWithdrawn)
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
func (it *ItokencontractPaidTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItokencontractPaidTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItokencontractPaidTokensWithdrawn represents a PaidTokensWithdrawn event raised by the Itokencontract contract.
type ItokencontractPaidTokensWithdrawn struct {
	TokenAddr common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPaidTokensWithdrawn is a free log retrieval operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Itokencontract *ItokencontractFilterer) FilterPaidTokensWithdrawn(opts *bind.FilterOpts, tokenAddr []common.Address) (*ItokencontractPaidTokensWithdrawnIterator, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _Itokencontract.contract.FilterLogs(opts, "PaidTokensWithdrawn", tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &ItokencontractPaidTokensWithdrawnIterator{contract: _Itokencontract.contract, event: "PaidTokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPaidTokensWithdrawn is a free log subscription operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Itokencontract *ItokencontractFilterer) WatchPaidTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *ItokencontractPaidTokensWithdrawn, tokenAddr []common.Address) (event.Subscription, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _Itokencontract.contract.WatchLogs(opts, "PaidTokensWithdrawn", tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItokencontractPaidTokensWithdrawn)
				if err := _Itokencontract.contract.UnpackLog(event, "PaidTokensWithdrawn", log); err != nil {
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

// ParsePaidTokensWithdrawn is a log parse operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Itokencontract *ItokencontractFilterer) ParsePaidTokensWithdrawn(log types.Log) (*ItokencontractPaidTokensWithdrawn, error) {
	event := new(ItokencontractPaidTokensWithdrawn)
	if err := _Itokencontract.contract.UnpackLog(event, "PaidTokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ItokencontractPaymentSuccessfulIterator is returned from FilterPaymentSuccessful and is used to iterate over the raw logs and unpacked data for PaymentSuccessful events raised by the Itokencontract contract.
type ItokencontractPaymentSuccessfulIterator struct {
	Event *ItokencontractPaymentSuccessful // Event containing the contract specifics and raw log

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
func (it *ItokencontractPaymentSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItokencontractPaymentSuccessful)
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
		it.Event = new(ItokencontractPaymentSuccessful)
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
func (it *ItokencontractPaymentSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItokencontractPaymentSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItokencontractPaymentSuccessful represents a PaymentSuccessful event raised by the Itokencontract contract.
type ItokencontractPaymentSuccessful struct {
	TokenAddress common.Address
	TokenAmount  *big.Int
	TokenPrice   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPaymentSuccessful is a free log retrieval operation binding the contract event 0xa056c95b265afa858f1320d1f0aad1aa914254d7fc80f83a8a0132bf2f29edbe.
//
// Solidity: event PaymentSuccessful(address indexed tokenAddress, uint256 tokenAmount, uint256 tokenPrice)
func (_Itokencontract *ItokencontractFilterer) FilterPaymentSuccessful(opts *bind.FilterOpts, tokenAddress []common.Address) (*ItokencontractPaymentSuccessfulIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Itokencontract.contract.FilterLogs(opts, "PaymentSuccessful", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ItokencontractPaymentSuccessfulIterator{contract: _Itokencontract.contract, event: "PaymentSuccessful", logs: logs, sub: sub}, nil
}

// WatchPaymentSuccessful is a free log subscription operation binding the contract event 0xa056c95b265afa858f1320d1f0aad1aa914254d7fc80f83a8a0132bf2f29edbe.
//
// Solidity: event PaymentSuccessful(address indexed tokenAddress, uint256 tokenAmount, uint256 tokenPrice)
func (_Itokencontract *ItokencontractFilterer) WatchPaymentSuccessful(opts *bind.WatchOpts, sink chan<- *ItokencontractPaymentSuccessful, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Itokencontract.contract.WatchLogs(opts, "PaymentSuccessful", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItokencontractPaymentSuccessful)
				if err := _Itokencontract.contract.UnpackLog(event, "PaymentSuccessful", log); err != nil {
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

// ParsePaymentSuccessful is a log parse operation binding the contract event 0xa056c95b265afa858f1320d1f0aad1aa914254d7fc80f83a8a0132bf2f29edbe.
//
// Solidity: event PaymentSuccessful(address indexed tokenAddress, uint256 tokenAmount, uint256 tokenPrice)
func (_Itokencontract *ItokencontractFilterer) ParsePaymentSuccessful(log types.Log) (*ItokencontractPaymentSuccessful, error) {
	event := new(ItokencontractPaymentSuccessful)
	if err := _Itokencontract.contract.UnpackLog(event, "PaymentSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ItokencontractTokenContractParamsUpdatedIterator is returned from FilterTokenContractParamsUpdated and is used to iterate over the raw logs and unpacked data for TokenContractParamsUpdated events raised by the Itokencontract contract.
type ItokencontractTokenContractParamsUpdatedIterator struct {
	Event *ItokencontractTokenContractParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ItokencontractTokenContractParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItokencontractTokenContractParamsUpdated)
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
		it.Event = new(ItokencontractTokenContractParamsUpdated)
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
func (it *ItokencontractTokenContractParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItokencontractTokenContractParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItokencontractTokenContractParamsUpdated represents a TokenContractParamsUpdated event raised by the Itokencontract contract.
type ItokencontractTokenContractParamsUpdated struct {
	NewPrice    *big.Int
	TokenName   string
	TokenSymbol string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenContractParamsUpdated is a free log retrieval operation binding the contract event 0x3419ca7e094a7b9d428666d41c7b45306d09be116d969b39e10447b02c2ee50f.
//
// Solidity: event TokenContractParamsUpdated(uint256 newPrice, string tokenName, string tokenSymbol)
func (_Itokencontract *ItokencontractFilterer) FilterTokenContractParamsUpdated(opts *bind.FilterOpts) (*ItokencontractTokenContractParamsUpdatedIterator, error) {

	logs, sub, err := _Itokencontract.contract.FilterLogs(opts, "TokenContractParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &ItokencontractTokenContractParamsUpdatedIterator{contract: _Itokencontract.contract, event: "TokenContractParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenContractParamsUpdated is a free log subscription operation binding the contract event 0x3419ca7e094a7b9d428666d41c7b45306d09be116d969b39e10447b02c2ee50f.
//
// Solidity: event TokenContractParamsUpdated(uint256 newPrice, string tokenName, string tokenSymbol)
func (_Itokencontract *ItokencontractFilterer) WatchTokenContractParamsUpdated(opts *bind.WatchOpts, sink chan<- *ItokencontractTokenContractParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _Itokencontract.contract.WatchLogs(opts, "TokenContractParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItokencontractTokenContractParamsUpdated)
				if err := _Itokencontract.contract.UnpackLog(event, "TokenContractParamsUpdated", log); err != nil {
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

// ParseTokenContractParamsUpdated is a log parse operation binding the contract event 0x3419ca7e094a7b9d428666d41c7b45306d09be116d969b39e10447b02c2ee50f.
//
// Solidity: event TokenContractParamsUpdated(uint256 newPrice, string tokenName, string tokenSymbol)
func (_Itokencontract *ItokencontractFilterer) ParseTokenContractParamsUpdated(log types.Log) (*ItokencontractTokenContractParamsUpdated, error) {
	event := new(ItokencontractTokenContractParamsUpdated)
	if err := _Itokencontract.contract.UnpackLog(event, "TokenContractParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ItokencontractTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the Itokencontract contract.
type ItokencontractTokenMintedIterator struct {
	Event *ItokencontractTokenMinted // Event containing the contract specifics and raw log

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
func (it *ItokencontractTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItokencontractTokenMinted)
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
		it.Event = new(ItokencontractTokenMinted)
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
func (it *ItokencontractTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItokencontractTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItokencontractTokenMinted represents a TokenMinted event raised by the Itokencontract contract.
type ItokencontractTokenMinted struct {
	Recipient common.Address
	TokenId   *big.Int
	TokenURI  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0xdf92894dc4675a7333caa5903b69cf5d8e8ec0d3f361c88207b6688e525703bb.
//
// Solidity: event TokenMinted(address indexed recipient, uint256 tokenId, string tokenURI)
func (_Itokencontract *ItokencontractFilterer) FilterTokenMinted(opts *bind.FilterOpts, recipient []common.Address) (*ItokencontractTokenMintedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Itokencontract.contract.FilterLogs(opts, "TokenMinted", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ItokencontractTokenMintedIterator{contract: _Itokencontract.contract, event: "TokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0xdf92894dc4675a7333caa5903b69cf5d8e8ec0d3f361c88207b6688e525703bb.
//
// Solidity: event TokenMinted(address indexed recipient, uint256 tokenId, string tokenURI)
func (_Itokencontract *ItokencontractFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *ItokencontractTokenMinted, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Itokencontract.contract.WatchLogs(opts, "TokenMinted", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItokencontractTokenMinted)
				if err := _Itokencontract.contract.UnpackLog(event, "TokenMinted", log); err != nil {
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

// ParseTokenMinted is a log parse operation binding the contract event 0xdf92894dc4675a7333caa5903b69cf5d8e8ec0d3f361c88207b6688e525703bb.
//
// Solidity: event TokenMinted(address indexed recipient, uint256 tokenId, string tokenURI)
func (_Itokencontract *ItokencontractFilterer) ParseTokenMinted(log types.Log) (*ItokencontractTokenMinted, error) {
	event := new(ItokencontractTokenMinted)
	if err := _Itokencontract.contract.UnpackLog(event, "TokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
