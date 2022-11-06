// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenfactory

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

// ITokenFactoryBaseTokenContractInfo is an auto generated low-level Go binding around an user-defined struct.
type ITokenFactoryBaseTokenContractInfo struct {
	TokenContractAddr common.Address
	PricePerOneToken  *big.Int
}

// ITokenFactoryUserNFTsInfo is an auto generated low-level Go binding around an user-defined struct.
type ITokenFactoryUserNFTsInfo struct {
	TokenContractAddr common.Address
	TokenIDs          []*big.Int
}

// TokenfactoryMetaData contains all meta data concerning the Tokenfactory contract.
var TokenfactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"adminsToUpdate\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdding\",\"type\":\"bool\"}],\"name\":\"AdminsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newBaseTokenContractsURI\",\"type\":\"string\"}],\"name\":\"BaseTokenContractsURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenContractId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTokenContractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"name\":\"TokenContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"adminsArr_\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"priceDecimals_\",\"type\":\"uint8\"}],\"name\":\"__TokenFactory_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenContractsURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenContractId_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"}],\"name\":\"deployTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenContractsArr_\",\"type\":\"address[]\"}],\"name\":\"getBaseTokenContractsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"}],\"internalType\":\"structITokenFactory.BaseTokenContractInfo[]\",\"name\":\"tokenContractsInfoArr_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContractsImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getTokenContractsPart\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"}],\"name\":\"getUserNFTsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddr\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"}],\"internalType\":\"structITokenFactory.UserNFTsInfo[]\",\"name\":\"userNFTsInfoArr_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolsBeacon\",\"outputs\":[{\"internalType\":\"contractProxyBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenContractsURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"}],\"name\":\"setNewImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenContractByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"adminsToUpdate_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isAdding_\",\"type\":\"bool\"}],\"name\":\"updateAdmins\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TokenfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenfactoryMetaData.ABI instead.
var TokenfactoryABI = TokenfactoryMetaData.ABI

// Tokenfactory is an auto generated Go binding around an Ethereum contract.
type Tokenfactory struct {
	TokenfactoryCaller     // Read-only binding to the contract
	TokenfactoryTransactor // Write-only binding to the contract
	TokenfactoryFilterer   // Log filterer for contract events
}

// TokenfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenfactorySession struct {
	Contract     *Tokenfactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenfactoryCallerSession struct {
	Contract *TokenfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenfactoryTransactorSession struct {
	Contract     *TokenfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenfactoryRaw struct {
	Contract *Tokenfactory // Generic contract binding to access the raw methods on
}

// TokenfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenfactoryCallerRaw struct {
	Contract *TokenfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenfactoryTransactorRaw struct {
	Contract *TokenfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenfactory creates a new instance of Tokenfactory, bound to a specific deployed contract.
func NewTokenfactory(address common.Address, backend bind.ContractBackend) (*Tokenfactory, error) {
	contract, err := bindTokenfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenfactory{TokenfactoryCaller: TokenfactoryCaller{contract: contract}, TokenfactoryTransactor: TokenfactoryTransactor{contract: contract}, TokenfactoryFilterer: TokenfactoryFilterer{contract: contract}}, nil
}

// NewTokenfactoryCaller creates a new read-only instance of Tokenfactory, bound to a specific deployed contract.
func NewTokenfactoryCaller(address common.Address, caller bind.ContractCaller) (*TokenfactoryCaller, error) {
	contract, err := bindTokenfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryCaller{contract: contract}, nil
}

// NewTokenfactoryTransactor creates a new write-only instance of Tokenfactory, bound to a specific deployed contract.
func NewTokenfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenfactoryTransactor, error) {
	contract, err := bindTokenfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryTransactor{contract: contract}, nil
}

// NewTokenfactoryFilterer creates a new log filterer instance of Tokenfactory, bound to a specific deployed contract.
func NewTokenfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenfactoryFilterer, error) {
	contract, err := bindTokenfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryFilterer{contract: contract}, nil
}

// bindTokenfactory binds a generic wrapper to an already deployed contract.
func bindTokenfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenfactory *TokenfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenfactory.Contract.TokenfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenfactory *TokenfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenfactory.Contract.TokenfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenfactory *TokenfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenfactory.Contract.TokenfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenfactory *TokenfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenfactory *TokenfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenfactory *TokenfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenfactory.Contract.contract.Transact(opts, method, params...)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Tokenfactory *TokenfactoryCaller) BaseTokenContractsURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "baseTokenContractsURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Tokenfactory *TokenfactorySession) BaseTokenContractsURI() (string, error) {
	return _Tokenfactory.Contract.BaseTokenContractsURI(&_Tokenfactory.CallOpts)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Tokenfactory *TokenfactoryCallerSession) BaseTokenContractsURI() (string, error) {
	return _Tokenfactory.Contract.BaseTokenContractsURI(&_Tokenfactory.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Tokenfactory *TokenfactoryCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Tokenfactory *TokenfactorySession) GetAdmins() ([]common.Address, error) {
	return _Tokenfactory.Contract.GetAdmins(&_Tokenfactory.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Tokenfactory *TokenfactoryCallerSession) GetAdmins() ([]common.Address, error) {
	return _Tokenfactory.Contract.GetAdmins(&_Tokenfactory.CallOpts)
}

// GetBaseTokenContractsInfo is a free data retrieval call binding the contract method 0x59f7278c.
//
// Solidity: function getBaseTokenContractsInfo(address[] tokenContractsArr_) view returns((address,uint256)[] tokenContractsInfoArr_)
func (_Tokenfactory *TokenfactoryCaller) GetBaseTokenContractsInfo(opts *bind.CallOpts, tokenContractsArr_ []common.Address) ([]ITokenFactoryBaseTokenContractInfo, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "getBaseTokenContractsInfo", tokenContractsArr_)

	if err != nil {
		return *new([]ITokenFactoryBaseTokenContractInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITokenFactoryBaseTokenContractInfo)).(*[]ITokenFactoryBaseTokenContractInfo)

	return out0, err

}

// GetBaseTokenContractsInfo is a free data retrieval call binding the contract method 0x59f7278c.
//
// Solidity: function getBaseTokenContractsInfo(address[] tokenContractsArr_) view returns((address,uint256)[] tokenContractsInfoArr_)
func (_Tokenfactory *TokenfactorySession) GetBaseTokenContractsInfo(tokenContractsArr_ []common.Address) ([]ITokenFactoryBaseTokenContractInfo, error) {
	return _Tokenfactory.Contract.GetBaseTokenContractsInfo(&_Tokenfactory.CallOpts, tokenContractsArr_)
}

// GetBaseTokenContractsInfo is a free data retrieval call binding the contract method 0x59f7278c.
//
// Solidity: function getBaseTokenContractsInfo(address[] tokenContractsArr_) view returns((address,uint256)[] tokenContractsInfoArr_)
func (_Tokenfactory *TokenfactoryCallerSession) GetBaseTokenContractsInfo(tokenContractsArr_ []common.Address) ([]ITokenFactoryBaseTokenContractInfo, error) {
	return _Tokenfactory.Contract.GetBaseTokenContractsInfo(&_Tokenfactory.CallOpts, tokenContractsArr_)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Tokenfactory *TokenfactoryCaller) GetTokenContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "getTokenContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Tokenfactory *TokenfactorySession) GetTokenContractsCount() (*big.Int, error) {
	return _Tokenfactory.Contract.GetTokenContractsCount(&_Tokenfactory.CallOpts)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Tokenfactory *TokenfactoryCallerSession) GetTokenContractsCount() (*big.Int, error) {
	return _Tokenfactory.Contract.GetTokenContractsCount(&_Tokenfactory.CallOpts)
}

// GetTokenContractsImpl is a free data retrieval call binding the contract method 0xfff7a62c.
//
// Solidity: function getTokenContractsImpl() view returns(address)
func (_Tokenfactory *TokenfactoryCaller) GetTokenContractsImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "getTokenContractsImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenContractsImpl is a free data retrieval call binding the contract method 0xfff7a62c.
//
// Solidity: function getTokenContractsImpl() view returns(address)
func (_Tokenfactory *TokenfactorySession) GetTokenContractsImpl() (common.Address, error) {
	return _Tokenfactory.Contract.GetTokenContractsImpl(&_Tokenfactory.CallOpts)
}

// GetTokenContractsImpl is a free data retrieval call binding the contract method 0xfff7a62c.
//
// Solidity: function getTokenContractsImpl() view returns(address)
func (_Tokenfactory *TokenfactoryCallerSession) GetTokenContractsImpl() (common.Address, error) {
	return _Tokenfactory.Contract.GetTokenContractsImpl(&_Tokenfactory.CallOpts)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Tokenfactory *TokenfactoryCaller) GetTokenContractsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "getTokenContractsPart", offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Tokenfactory *TokenfactorySession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Tokenfactory.Contract.GetTokenContractsPart(&_Tokenfactory.CallOpts, offset_, limit_)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Tokenfactory *TokenfactoryCallerSession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Tokenfactory.Contract.GetTokenContractsPart(&_Tokenfactory.CallOpts, offset_, limit_)
}

// GetUserNFTsInfo is a free data retrieval call binding the contract method 0x066d344b.
//
// Solidity: function getUserNFTsInfo(address userAddr_) view returns((address,uint256[])[] userNFTsInfoArr_)
func (_Tokenfactory *TokenfactoryCaller) GetUserNFTsInfo(opts *bind.CallOpts, userAddr_ common.Address) ([]ITokenFactoryUserNFTsInfo, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "getUserNFTsInfo", userAddr_)

	if err != nil {
		return *new([]ITokenFactoryUserNFTsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITokenFactoryUserNFTsInfo)).(*[]ITokenFactoryUserNFTsInfo)

	return out0, err

}

// GetUserNFTsInfo is a free data retrieval call binding the contract method 0x066d344b.
//
// Solidity: function getUserNFTsInfo(address userAddr_) view returns((address,uint256[])[] userNFTsInfoArr_)
func (_Tokenfactory *TokenfactorySession) GetUserNFTsInfo(userAddr_ common.Address) ([]ITokenFactoryUserNFTsInfo, error) {
	return _Tokenfactory.Contract.GetUserNFTsInfo(&_Tokenfactory.CallOpts, userAddr_)
}

// GetUserNFTsInfo is a free data retrieval call binding the contract method 0x066d344b.
//
// Solidity: function getUserNFTsInfo(address userAddr_) view returns((address,uint256[])[] userNFTsInfoArr_)
func (_Tokenfactory *TokenfactoryCallerSession) GetUserNFTsInfo(userAddr_ common.Address) ([]ITokenFactoryUserNFTsInfo, error) {
	return _Tokenfactory.Contract.GetUserNFTsInfo(&_Tokenfactory.CallOpts, userAddr_)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address userAddr_) view returns(bool)
func (_Tokenfactory *TokenfactoryCaller) IsAdmin(opts *bind.CallOpts, userAddr_ common.Address) (bool, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "isAdmin", userAddr_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address userAddr_) view returns(bool)
func (_Tokenfactory *TokenfactorySession) IsAdmin(userAddr_ common.Address) (bool, error) {
	return _Tokenfactory.Contract.IsAdmin(&_Tokenfactory.CallOpts, userAddr_)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address userAddr_) view returns(bool)
func (_Tokenfactory *TokenfactoryCallerSession) IsAdmin(userAddr_ common.Address) (bool, error) {
	return _Tokenfactory.Contract.IsAdmin(&_Tokenfactory.CallOpts, userAddr_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenfactory *TokenfactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenfactory *TokenfactorySession) Owner() (common.Address, error) {
	return _Tokenfactory.Contract.Owner(&_Tokenfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenfactory *TokenfactoryCallerSession) Owner() (common.Address, error) {
	return _Tokenfactory.Contract.Owner(&_Tokenfactory.CallOpts)
}

// PoolsBeacon is a free data retrieval call binding the contract method 0xb3f251e7.
//
// Solidity: function poolsBeacon() view returns(address)
func (_Tokenfactory *TokenfactoryCaller) PoolsBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "poolsBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolsBeacon is a free data retrieval call binding the contract method 0xb3f251e7.
//
// Solidity: function poolsBeacon() view returns(address)
func (_Tokenfactory *TokenfactorySession) PoolsBeacon() (common.Address, error) {
	return _Tokenfactory.Contract.PoolsBeacon(&_Tokenfactory.CallOpts)
}

// PoolsBeacon is a free data retrieval call binding the contract method 0xb3f251e7.
//
// Solidity: function poolsBeacon() view returns(address)
func (_Tokenfactory *TokenfactoryCallerSession) PoolsBeacon() (common.Address, error) {
	return _Tokenfactory.Contract.PoolsBeacon(&_Tokenfactory.CallOpts)
}

// PriceDecimals is a free data retrieval call binding the contract method 0x05300b28.
//
// Solidity: function priceDecimals() view returns(uint8)
func (_Tokenfactory *TokenfactoryCaller) PriceDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "priceDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PriceDecimals is a free data retrieval call binding the contract method 0x05300b28.
//
// Solidity: function priceDecimals() view returns(uint8)
func (_Tokenfactory *TokenfactorySession) PriceDecimals() (uint8, error) {
	return _Tokenfactory.Contract.PriceDecimals(&_Tokenfactory.CallOpts)
}

// PriceDecimals is a free data retrieval call binding the contract method 0x05300b28.
//
// Solidity: function priceDecimals() view returns(uint8)
func (_Tokenfactory *TokenfactoryCallerSession) PriceDecimals() (uint8, error) {
	return _Tokenfactory.Contract.PriceDecimals(&_Tokenfactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Tokenfactory *TokenfactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Tokenfactory *TokenfactorySession) ProxiableUUID() ([32]byte, error) {
	return _Tokenfactory.Contract.ProxiableUUID(&_Tokenfactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Tokenfactory *TokenfactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Tokenfactory.Contract.ProxiableUUID(&_Tokenfactory.CallOpts)
}

// TokenContractByIndex is a free data retrieval call binding the contract method 0xb7e6a3ec.
//
// Solidity: function tokenContractByIndex(uint256 ) view returns(address)
func (_Tokenfactory *TokenfactoryCaller) TokenContractByIndex(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "tokenContractByIndex", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenContractByIndex is a free data retrieval call binding the contract method 0xb7e6a3ec.
//
// Solidity: function tokenContractByIndex(uint256 ) view returns(address)
func (_Tokenfactory *TokenfactorySession) TokenContractByIndex(arg0 *big.Int) (common.Address, error) {
	return _Tokenfactory.Contract.TokenContractByIndex(&_Tokenfactory.CallOpts, arg0)
}

// TokenContractByIndex is a free data retrieval call binding the contract method 0xb7e6a3ec.
//
// Solidity: function tokenContractByIndex(uint256 ) view returns(address)
func (_Tokenfactory *TokenfactoryCallerSession) TokenContractByIndex(arg0 *big.Int) (common.Address, error) {
	return _Tokenfactory.Contract.TokenContractByIndex(&_Tokenfactory.CallOpts, arg0)
}

// TokenFactoryInit is a paid mutator transaction binding the contract method 0xae13b76e.
//
// Solidity: function __TokenFactory_init(address[] adminsArr_, string baseTokenContractsURI_, uint8 priceDecimals_) returns()
func (_Tokenfactory *TokenfactoryTransactor) TokenFactoryInit(opts *bind.TransactOpts, adminsArr_ []common.Address, baseTokenContractsURI_ string, priceDecimals_ uint8) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "__TokenFactory_init", adminsArr_, baseTokenContractsURI_, priceDecimals_)
}

// TokenFactoryInit is a paid mutator transaction binding the contract method 0xae13b76e.
//
// Solidity: function __TokenFactory_init(address[] adminsArr_, string baseTokenContractsURI_, uint8 priceDecimals_) returns()
func (_Tokenfactory *TokenfactorySession) TokenFactoryInit(adminsArr_ []common.Address, baseTokenContractsURI_ string, priceDecimals_ uint8) (*types.Transaction, error) {
	return _Tokenfactory.Contract.TokenFactoryInit(&_Tokenfactory.TransactOpts, adminsArr_, baseTokenContractsURI_, priceDecimals_)
}

// TokenFactoryInit is a paid mutator transaction binding the contract method 0xae13b76e.
//
// Solidity: function __TokenFactory_init(address[] adminsArr_, string baseTokenContractsURI_, uint8 priceDecimals_) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) TokenFactoryInit(adminsArr_ []common.Address, baseTokenContractsURI_ string, priceDecimals_ uint8) (*types.Transaction, error) {
	return _Tokenfactory.Contract.TokenFactoryInit(&_Tokenfactory.TransactOpts, adminsArr_, baseTokenContractsURI_, priceDecimals_)
}

// DeployTokenContract is a paid mutator transaction binding the contract method 0xda8aa768.
//
// Solidity: function deployTokenContract(uint256 tokenContractId_, string tokenName_, string tokenSymbol_, uint256 pricePerOneToken_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_Tokenfactory *TokenfactoryTransactor) DeployTokenContract(opts *bind.TransactOpts, tokenContractId_ *big.Int, tokenName_ string, tokenSymbol_ string, pricePerOneToken_ *big.Int, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "deployTokenContract", tokenContractId_, tokenName_, tokenSymbol_, pricePerOneToken_, r_, s_, v_)
}

// DeployTokenContract is a paid mutator transaction binding the contract method 0xda8aa768.
//
// Solidity: function deployTokenContract(uint256 tokenContractId_, string tokenName_, string tokenSymbol_, uint256 pricePerOneToken_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_Tokenfactory *TokenfactorySession) DeployTokenContract(tokenContractId_ *big.Int, tokenName_ string, tokenSymbol_ string, pricePerOneToken_ *big.Int, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Tokenfactory.Contract.DeployTokenContract(&_Tokenfactory.TransactOpts, tokenContractId_, tokenName_, tokenSymbol_, pricePerOneToken_, r_, s_, v_)
}

// DeployTokenContract is a paid mutator transaction binding the contract method 0xda8aa768.
//
// Solidity: function deployTokenContract(uint256 tokenContractId_, string tokenName_, string tokenSymbol_, uint256 pricePerOneToken_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) DeployTokenContract(tokenContractId_ *big.Int, tokenName_ string, tokenSymbol_ string, pricePerOneToken_ *big.Int, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Tokenfactory.Contract.DeployTokenContract(&_Tokenfactory.TransactOpts, tokenContractId_, tokenName_, tokenSymbol_, pricePerOneToken_, r_, s_, v_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenfactory *TokenfactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenfactory *TokenfactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenfactory.Contract.RenounceOwnership(&_Tokenfactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenfactory *TokenfactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenfactory.Contract.RenounceOwnership(&_Tokenfactory.TransactOpts)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Tokenfactory *TokenfactoryTransactor) SetBaseTokenContractsURI(opts *bind.TransactOpts, baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "setBaseTokenContractsURI", baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Tokenfactory *TokenfactorySession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetBaseTokenContractsURI(&_Tokenfactory.TransactOpts, baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetBaseTokenContractsURI(&_Tokenfactory.TransactOpts, baseTokenContractsURI_)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address newImplementation_) returns()
func (_Tokenfactory *TokenfactoryTransactor) SetNewImplementation(opts *bind.TransactOpts, newImplementation_ common.Address) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "setNewImplementation", newImplementation_)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address newImplementation_) returns()
func (_Tokenfactory *TokenfactorySession) SetNewImplementation(newImplementation_ common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetNewImplementation(&_Tokenfactory.TransactOpts, newImplementation_)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address newImplementation_) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) SetNewImplementation(newImplementation_ common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetNewImplementation(&_Tokenfactory.TransactOpts, newImplementation_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenfactory *TokenfactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenfactory *TokenfactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.TransferOwnership(&_Tokenfactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.TransferOwnership(&_Tokenfactory.TransactOpts, newOwner)
}

// UpdateAdmins is a paid mutator transaction binding the contract method 0x809f976a.
//
// Solidity: function updateAdmins(address[] adminsToUpdate_, bool isAdding_) returns()
func (_Tokenfactory *TokenfactoryTransactor) UpdateAdmins(opts *bind.TransactOpts, adminsToUpdate_ []common.Address, isAdding_ bool) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "updateAdmins", adminsToUpdate_, isAdding_)
}

// UpdateAdmins is a paid mutator transaction binding the contract method 0x809f976a.
//
// Solidity: function updateAdmins(address[] adminsToUpdate_, bool isAdding_) returns()
func (_Tokenfactory *TokenfactorySession) UpdateAdmins(adminsToUpdate_ []common.Address, isAdding_ bool) (*types.Transaction, error) {
	return _Tokenfactory.Contract.UpdateAdmins(&_Tokenfactory.TransactOpts, adminsToUpdate_, isAdding_)
}

// UpdateAdmins is a paid mutator transaction binding the contract method 0x809f976a.
//
// Solidity: function updateAdmins(address[] adminsToUpdate_, bool isAdding_) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) UpdateAdmins(adminsToUpdate_ []common.Address, isAdding_ bool) (*types.Transaction, error) {
	return _Tokenfactory.Contract.UpdateAdmins(&_Tokenfactory.TransactOpts, adminsToUpdate_, isAdding_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Tokenfactory *TokenfactoryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Tokenfactory *TokenfactorySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.UpgradeTo(&_Tokenfactory.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.UpgradeTo(&_Tokenfactory.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Tokenfactory *TokenfactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Tokenfactory *TokenfactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Tokenfactory.Contract.UpgradeToAndCall(&_Tokenfactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Tokenfactory *TokenfactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Tokenfactory.Contract.UpgradeToAndCall(&_Tokenfactory.TransactOpts, newImplementation, data)
}

// TokenfactoryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Tokenfactory contract.
type TokenfactoryAdminChangedIterator struct {
	Event *TokenfactoryAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenfactoryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryAdminChanged)
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
		it.Event = new(TokenfactoryAdminChanged)
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
func (it *TokenfactoryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryAdminChanged represents a AdminChanged event raised by the Tokenfactory contract.
type TokenfactoryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Tokenfactory *TokenfactoryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TokenfactoryAdminChangedIterator, error) {

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TokenfactoryAdminChangedIterator{contract: _Tokenfactory.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Tokenfactory *TokenfactoryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenfactoryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryAdminChanged)
				if err := _Tokenfactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Tokenfactory *TokenfactoryFilterer) ParseAdminChanged(log types.Log) (*TokenfactoryAdminChanged, error) {
	event := new(TokenfactoryAdminChanged)
	if err := _Tokenfactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryAdminsUpdatedIterator is returned from FilterAdminsUpdated and is used to iterate over the raw logs and unpacked data for AdminsUpdated events raised by the Tokenfactory contract.
type TokenfactoryAdminsUpdatedIterator struct {
	Event *TokenfactoryAdminsUpdated // Event containing the contract specifics and raw log

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
func (it *TokenfactoryAdminsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryAdminsUpdated)
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
		it.Event = new(TokenfactoryAdminsUpdated)
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
func (it *TokenfactoryAdminsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryAdminsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryAdminsUpdated represents a AdminsUpdated event raised by the Tokenfactory contract.
type TokenfactoryAdminsUpdated struct {
	AdminsToUpdate []common.Address
	IsAdding       bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAdminsUpdated is a free log retrieval operation binding the contract event 0x869da43fad8e93387e0f5b58027ae3bdc4c61756fc5219fa85a2e4f44d3ebea4.
//
// Solidity: event AdminsUpdated(address[] adminsToUpdate, bool isAdding)
func (_Tokenfactory *TokenfactoryFilterer) FilterAdminsUpdated(opts *bind.FilterOpts) (*TokenfactoryAdminsUpdatedIterator, error) {

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "AdminsUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenfactoryAdminsUpdatedIterator{contract: _Tokenfactory.contract, event: "AdminsUpdated", logs: logs, sub: sub}, nil
}

// WatchAdminsUpdated is a free log subscription operation binding the contract event 0x869da43fad8e93387e0f5b58027ae3bdc4c61756fc5219fa85a2e4f44d3ebea4.
//
// Solidity: event AdminsUpdated(address[] adminsToUpdate, bool isAdding)
func (_Tokenfactory *TokenfactoryFilterer) WatchAdminsUpdated(opts *bind.WatchOpts, sink chan<- *TokenfactoryAdminsUpdated) (event.Subscription, error) {

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "AdminsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryAdminsUpdated)
				if err := _Tokenfactory.contract.UnpackLog(event, "AdminsUpdated", log); err != nil {
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

// ParseAdminsUpdated is a log parse operation binding the contract event 0x869da43fad8e93387e0f5b58027ae3bdc4c61756fc5219fa85a2e4f44d3ebea4.
//
// Solidity: event AdminsUpdated(address[] adminsToUpdate, bool isAdding)
func (_Tokenfactory *TokenfactoryFilterer) ParseAdminsUpdated(log types.Log) (*TokenfactoryAdminsUpdated, error) {
	event := new(TokenfactoryAdminsUpdated)
	if err := _Tokenfactory.contract.UnpackLog(event, "AdminsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryBaseTokenContractsURIUpdatedIterator is returned from FilterBaseTokenContractsURIUpdated and is used to iterate over the raw logs and unpacked data for BaseTokenContractsURIUpdated events raised by the Tokenfactory contract.
type TokenfactoryBaseTokenContractsURIUpdatedIterator struct {
	Event *TokenfactoryBaseTokenContractsURIUpdated // Event containing the contract specifics and raw log

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
func (it *TokenfactoryBaseTokenContractsURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryBaseTokenContractsURIUpdated)
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
		it.Event = new(TokenfactoryBaseTokenContractsURIUpdated)
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
func (it *TokenfactoryBaseTokenContractsURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryBaseTokenContractsURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryBaseTokenContractsURIUpdated represents a BaseTokenContractsURIUpdated event raised by the Tokenfactory contract.
type TokenfactoryBaseTokenContractsURIUpdated struct {
	NewBaseTokenContractsURI string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterBaseTokenContractsURIUpdated is a free log retrieval operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Tokenfactory *TokenfactoryFilterer) FilterBaseTokenContractsURIUpdated(opts *bind.FilterOpts) (*TokenfactoryBaseTokenContractsURIUpdatedIterator, error) {

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenfactoryBaseTokenContractsURIUpdatedIterator{contract: _Tokenfactory.contract, event: "BaseTokenContractsURIUpdated", logs: logs, sub: sub}, nil
}

// WatchBaseTokenContractsURIUpdated is a free log subscription operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Tokenfactory *TokenfactoryFilterer) WatchBaseTokenContractsURIUpdated(opts *bind.WatchOpts, sink chan<- *TokenfactoryBaseTokenContractsURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryBaseTokenContractsURIUpdated)
				if err := _Tokenfactory.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
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

// ParseBaseTokenContractsURIUpdated is a log parse operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Tokenfactory *TokenfactoryFilterer) ParseBaseTokenContractsURIUpdated(log types.Log) (*TokenfactoryBaseTokenContractsURIUpdated, error) {
	event := new(TokenfactoryBaseTokenContractsURIUpdated)
	if err := _Tokenfactory.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Tokenfactory contract.
type TokenfactoryBeaconUpgradedIterator struct {
	Event *TokenfactoryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *TokenfactoryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryBeaconUpgraded)
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
		it.Event = new(TokenfactoryBeaconUpgraded)
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
func (it *TokenfactoryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryBeaconUpgraded represents a BeaconUpgraded event raised by the Tokenfactory contract.
type TokenfactoryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Tokenfactory *TokenfactoryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*TokenfactoryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryBeaconUpgradedIterator{contract: _Tokenfactory.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Tokenfactory *TokenfactoryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *TokenfactoryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryBeaconUpgraded)
				if err := _Tokenfactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Tokenfactory *TokenfactoryFilterer) ParseBeaconUpgraded(log types.Log) (*TokenfactoryBeaconUpgraded, error) {
	event := new(TokenfactoryBeaconUpgraded)
	if err := _Tokenfactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Tokenfactory contract.
type TokenfactoryInitializedIterator struct {
	Event *TokenfactoryInitialized // Event containing the contract specifics and raw log

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
func (it *TokenfactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryInitialized)
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
		it.Event = new(TokenfactoryInitialized)
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
func (it *TokenfactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryInitialized represents a Initialized event raised by the Tokenfactory contract.
type TokenfactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenfactory *TokenfactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenfactoryInitializedIterator, error) {

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenfactoryInitializedIterator{contract: _Tokenfactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenfactory *TokenfactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenfactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryInitialized)
				if err := _Tokenfactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Tokenfactory *TokenfactoryFilterer) ParseInitialized(log types.Log) (*TokenfactoryInitialized, error) {
	event := new(TokenfactoryInitialized)
	if err := _Tokenfactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokenfactory contract.
type TokenfactoryOwnershipTransferredIterator struct {
	Event *TokenfactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenfactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryOwnershipTransferred)
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
		it.Event = new(TokenfactoryOwnershipTransferred)
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
func (it *TokenfactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Tokenfactory contract.
type TokenfactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenfactory *TokenfactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenfactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryOwnershipTransferredIterator{contract: _Tokenfactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenfactory *TokenfactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenfactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryOwnershipTransferred)
				if err := _Tokenfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Tokenfactory *TokenfactoryFilterer) ParseOwnershipTransferred(log types.Log) (*TokenfactoryOwnershipTransferred, error) {
	event := new(TokenfactoryOwnershipTransferred)
	if err := _Tokenfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryTokenContractDeployedIterator is returned from FilterTokenContractDeployed and is used to iterate over the raw logs and unpacked data for TokenContractDeployed events raised by the Tokenfactory contract.
type TokenfactoryTokenContractDeployedIterator struct {
	Event *TokenfactoryTokenContractDeployed // Event containing the contract specifics and raw log

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
func (it *TokenfactoryTokenContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryTokenContractDeployed)
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
		it.Event = new(TokenfactoryTokenContractDeployed)
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
func (it *TokenfactoryTokenContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryTokenContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryTokenContractDeployed represents a TokenContractDeployed event raised by the Tokenfactory contract.
type TokenfactoryTokenContractDeployed struct {
	TokenContractId      *big.Int
	NewTokenContractAddr common.Address
	PricePerOneToken     *big.Int
	TokenName            string
	TokenSymbol          string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTokenContractDeployed is a free log retrieval operation binding the contract event 0x8412389c512c8db70cea087e4ae10396c0fc6d2f2fe1efff6fedd9dc87fc228f.
//
// Solidity: event TokenContractDeployed(uint256 tokenContractId, address newTokenContractAddr, uint256 pricePerOneToken, string tokenName, string tokenSymbol)
func (_Tokenfactory *TokenfactoryFilterer) FilterTokenContractDeployed(opts *bind.FilterOpts) (*TokenfactoryTokenContractDeployedIterator, error) {

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "TokenContractDeployed")
	if err != nil {
		return nil, err
	}
	return &TokenfactoryTokenContractDeployedIterator{contract: _Tokenfactory.contract, event: "TokenContractDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenContractDeployed is a free log subscription operation binding the contract event 0x8412389c512c8db70cea087e4ae10396c0fc6d2f2fe1efff6fedd9dc87fc228f.
//
// Solidity: event TokenContractDeployed(uint256 tokenContractId, address newTokenContractAddr, uint256 pricePerOneToken, string tokenName, string tokenSymbol)
func (_Tokenfactory *TokenfactoryFilterer) WatchTokenContractDeployed(opts *bind.WatchOpts, sink chan<- *TokenfactoryTokenContractDeployed) (event.Subscription, error) {

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "TokenContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryTokenContractDeployed)
				if err := _Tokenfactory.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
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

// ParseTokenContractDeployed is a log parse operation binding the contract event 0x8412389c512c8db70cea087e4ae10396c0fc6d2f2fe1efff6fedd9dc87fc228f.
//
// Solidity: event TokenContractDeployed(uint256 tokenContractId, address newTokenContractAddr, uint256 pricePerOneToken, string tokenName, string tokenSymbol)
func (_Tokenfactory *TokenfactoryFilterer) ParseTokenContractDeployed(log types.Log) (*TokenfactoryTokenContractDeployed, error) {
	event := new(TokenfactoryTokenContractDeployed)
	if err := _Tokenfactory.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Tokenfactory contract.
type TokenfactoryUpgradedIterator struct {
	Event *TokenfactoryUpgraded // Event containing the contract specifics and raw log

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
func (it *TokenfactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryUpgraded)
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
		it.Event = new(TokenfactoryUpgraded)
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
func (it *TokenfactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryUpgraded represents a Upgraded event raised by the Tokenfactory contract.
type TokenfactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Tokenfactory *TokenfactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TokenfactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryUpgradedIterator{contract: _Tokenfactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Tokenfactory *TokenfactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TokenfactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryUpgraded)
				if err := _Tokenfactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Tokenfactory *TokenfactoryFilterer) ParseUpgraded(log types.Log) (*TokenfactoryUpgraded, error) {
	event := new(TokenfactoryUpgraded)
	if err := _Tokenfactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
