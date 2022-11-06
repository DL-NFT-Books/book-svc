// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itokenfactory

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

// ItokenfactoryMetaData contains all meta data concerning the Itokenfactory contract.
var ItokenfactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"adminsToUpdate\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdding\",\"type\":\"bool\"}],\"name\":\"AdminsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newBaseTokenContractsURI\",\"type\":\"string\"}],\"name\":\"BaseTokenContractsURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenContractId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTokenContractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"name\":\"TokenContractDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"adminsArr_\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"priceDecimals_\",\"type\":\"uint8\"}],\"name\":\"__TokenFactory_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenContractsURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenContractId_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"}],\"name\":\"deployTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenContractsArr_\",\"type\":\"address[]\"}],\"name\":\"getBaseTokenContractsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"}],\"internalType\":\"structITokenFactory.BaseTokenContractInfo[]\",\"name\":\"tokenContractsInfoArr_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContractsImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getTokenContractsPart\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"}],\"name\":\"getUserNFTsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddr\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"}],\"internalType\":\"structITokenFactory.UserNFTsInfo[]\",\"name\":\"userNFTsInfoArr_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolsBeacon\",\"outputs\":[{\"internalType\":\"contractProxyBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenContractsURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"}],\"name\":\"setNewImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenContractId_\",\"type\":\"uint256\"}],\"name\":\"tokenContractByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"adminsToUpdate_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isAdding_\",\"type\":\"bool\"}],\"name\":\"updateAdmins\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ItokenfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ItokenfactoryMetaData.ABI instead.
var ItokenfactoryABI = ItokenfactoryMetaData.ABI

// Itokenfactory is an auto generated Go binding around an Ethereum contract.
type Itokenfactory struct {
	ItokenfactoryCaller     // Read-only binding to the contract
	ItokenfactoryTransactor // Write-only binding to the contract
	ItokenfactoryFilterer   // Log filterer for contract events
}

// ItokenfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ItokenfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokenfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ItokenfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokenfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ItokenfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokenfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ItokenfactorySession struct {
	Contract     *Itokenfactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ItokenfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ItokenfactoryCallerSession struct {
	Contract *ItokenfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ItokenfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ItokenfactoryTransactorSession struct {
	Contract     *ItokenfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ItokenfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ItokenfactoryRaw struct {
	Contract *Itokenfactory // Generic contract binding to access the raw methods on
}

// ItokenfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ItokenfactoryCallerRaw struct {
	Contract *ItokenfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ItokenfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ItokenfactoryTransactorRaw struct {
	Contract *ItokenfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewItokenfactory creates a new instance of Itokenfactory, bound to a specific deployed contract.
func NewItokenfactory(address common.Address, backend bind.ContractBackend) (*Itokenfactory, error) {
	contract, err := bindItokenfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Itokenfactory{ItokenfactoryCaller: ItokenfactoryCaller{contract: contract}, ItokenfactoryTransactor: ItokenfactoryTransactor{contract: contract}, ItokenfactoryFilterer: ItokenfactoryFilterer{contract: contract}}, nil
}

// NewItokenfactoryCaller creates a new read-only instance of Itokenfactory, bound to a specific deployed contract.
func NewItokenfactoryCaller(address common.Address, caller bind.ContractCaller) (*ItokenfactoryCaller, error) {
	contract, err := bindItokenfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ItokenfactoryCaller{contract: contract}, nil
}

// NewItokenfactoryTransactor creates a new write-only instance of Itokenfactory, bound to a specific deployed contract.
func NewItokenfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ItokenfactoryTransactor, error) {
	contract, err := bindItokenfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ItokenfactoryTransactor{contract: contract}, nil
}

// NewItokenfactoryFilterer creates a new log filterer instance of Itokenfactory, bound to a specific deployed contract.
func NewItokenfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ItokenfactoryFilterer, error) {
	contract, err := bindItokenfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ItokenfactoryFilterer{contract: contract}, nil
}

// bindItokenfactory binds a generic wrapper to an already deployed contract.
func bindItokenfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ItokenfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itokenfactory *ItokenfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Itokenfactory.Contract.ItokenfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itokenfactory *ItokenfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itokenfactory.Contract.ItokenfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itokenfactory *ItokenfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Itokenfactory.Contract.ItokenfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itokenfactory *ItokenfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Itokenfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itokenfactory *ItokenfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itokenfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itokenfactory *ItokenfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Itokenfactory.Contract.contract.Transact(opts, method, params...)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Itokenfactory *ItokenfactoryCaller) BaseTokenContractsURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "baseTokenContractsURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Itokenfactory *ItokenfactorySession) BaseTokenContractsURI() (string, error) {
	return _Itokenfactory.Contract.BaseTokenContractsURI(&_Itokenfactory.CallOpts)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Itokenfactory *ItokenfactoryCallerSession) BaseTokenContractsURI() (string, error) {
	return _Itokenfactory.Contract.BaseTokenContractsURI(&_Itokenfactory.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Itokenfactory *ItokenfactoryCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Itokenfactory *ItokenfactorySession) GetAdmins() ([]common.Address, error) {
	return _Itokenfactory.Contract.GetAdmins(&_Itokenfactory.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Itokenfactory *ItokenfactoryCallerSession) GetAdmins() ([]common.Address, error) {
	return _Itokenfactory.Contract.GetAdmins(&_Itokenfactory.CallOpts)
}

// GetBaseTokenContractsInfo is a free data retrieval call binding the contract method 0x59f7278c.
//
// Solidity: function getBaseTokenContractsInfo(address[] tokenContractsArr_) view returns((address,uint256)[] tokenContractsInfoArr_)
func (_Itokenfactory *ItokenfactoryCaller) GetBaseTokenContractsInfo(opts *bind.CallOpts, tokenContractsArr_ []common.Address) ([]ITokenFactoryBaseTokenContractInfo, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "getBaseTokenContractsInfo", tokenContractsArr_)

	if err != nil {
		return *new([]ITokenFactoryBaseTokenContractInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITokenFactoryBaseTokenContractInfo)).(*[]ITokenFactoryBaseTokenContractInfo)

	return out0, err

}

// GetBaseTokenContractsInfo is a free data retrieval call binding the contract method 0x59f7278c.
//
// Solidity: function getBaseTokenContractsInfo(address[] tokenContractsArr_) view returns((address,uint256)[] tokenContractsInfoArr_)
func (_Itokenfactory *ItokenfactorySession) GetBaseTokenContractsInfo(tokenContractsArr_ []common.Address) ([]ITokenFactoryBaseTokenContractInfo, error) {
	return _Itokenfactory.Contract.GetBaseTokenContractsInfo(&_Itokenfactory.CallOpts, tokenContractsArr_)
}

// GetBaseTokenContractsInfo is a free data retrieval call binding the contract method 0x59f7278c.
//
// Solidity: function getBaseTokenContractsInfo(address[] tokenContractsArr_) view returns((address,uint256)[] tokenContractsInfoArr_)
func (_Itokenfactory *ItokenfactoryCallerSession) GetBaseTokenContractsInfo(tokenContractsArr_ []common.Address) ([]ITokenFactoryBaseTokenContractInfo, error) {
	return _Itokenfactory.Contract.GetBaseTokenContractsInfo(&_Itokenfactory.CallOpts, tokenContractsArr_)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Itokenfactory *ItokenfactoryCaller) GetTokenContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "getTokenContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Itokenfactory *ItokenfactorySession) GetTokenContractsCount() (*big.Int, error) {
	return _Itokenfactory.Contract.GetTokenContractsCount(&_Itokenfactory.CallOpts)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Itokenfactory *ItokenfactoryCallerSession) GetTokenContractsCount() (*big.Int, error) {
	return _Itokenfactory.Contract.GetTokenContractsCount(&_Itokenfactory.CallOpts)
}

// GetTokenContractsImpl is a free data retrieval call binding the contract method 0xfff7a62c.
//
// Solidity: function getTokenContractsImpl() view returns(address)
func (_Itokenfactory *ItokenfactoryCaller) GetTokenContractsImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "getTokenContractsImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenContractsImpl is a free data retrieval call binding the contract method 0xfff7a62c.
//
// Solidity: function getTokenContractsImpl() view returns(address)
func (_Itokenfactory *ItokenfactorySession) GetTokenContractsImpl() (common.Address, error) {
	return _Itokenfactory.Contract.GetTokenContractsImpl(&_Itokenfactory.CallOpts)
}

// GetTokenContractsImpl is a free data retrieval call binding the contract method 0xfff7a62c.
//
// Solidity: function getTokenContractsImpl() view returns(address)
func (_Itokenfactory *ItokenfactoryCallerSession) GetTokenContractsImpl() (common.Address, error) {
	return _Itokenfactory.Contract.GetTokenContractsImpl(&_Itokenfactory.CallOpts)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Itokenfactory *ItokenfactoryCaller) GetTokenContractsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "getTokenContractsPart", offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Itokenfactory *ItokenfactorySession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Itokenfactory.Contract.GetTokenContractsPart(&_Itokenfactory.CallOpts, offset_, limit_)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Itokenfactory *ItokenfactoryCallerSession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Itokenfactory.Contract.GetTokenContractsPart(&_Itokenfactory.CallOpts, offset_, limit_)
}

// GetUserNFTsInfo is a free data retrieval call binding the contract method 0x066d344b.
//
// Solidity: function getUserNFTsInfo(address userAddr_) view returns((address,uint256[])[] userNFTsInfoArr_)
func (_Itokenfactory *ItokenfactoryCaller) GetUserNFTsInfo(opts *bind.CallOpts, userAddr_ common.Address) ([]ITokenFactoryUserNFTsInfo, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "getUserNFTsInfo", userAddr_)

	if err != nil {
		return *new([]ITokenFactoryUserNFTsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITokenFactoryUserNFTsInfo)).(*[]ITokenFactoryUserNFTsInfo)

	return out0, err

}

// GetUserNFTsInfo is a free data retrieval call binding the contract method 0x066d344b.
//
// Solidity: function getUserNFTsInfo(address userAddr_) view returns((address,uint256[])[] userNFTsInfoArr_)
func (_Itokenfactory *ItokenfactorySession) GetUserNFTsInfo(userAddr_ common.Address) ([]ITokenFactoryUserNFTsInfo, error) {
	return _Itokenfactory.Contract.GetUserNFTsInfo(&_Itokenfactory.CallOpts, userAddr_)
}

// GetUserNFTsInfo is a free data retrieval call binding the contract method 0x066d344b.
//
// Solidity: function getUserNFTsInfo(address userAddr_) view returns((address,uint256[])[] userNFTsInfoArr_)
func (_Itokenfactory *ItokenfactoryCallerSession) GetUserNFTsInfo(userAddr_ common.Address) ([]ITokenFactoryUserNFTsInfo, error) {
	return _Itokenfactory.Contract.GetUserNFTsInfo(&_Itokenfactory.CallOpts, userAddr_)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address userAddr_) view returns(bool)
func (_Itokenfactory *ItokenfactoryCaller) IsAdmin(opts *bind.CallOpts, userAddr_ common.Address) (bool, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "isAdmin", userAddr_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address userAddr_) view returns(bool)
func (_Itokenfactory *ItokenfactorySession) IsAdmin(userAddr_ common.Address) (bool, error) {
	return _Itokenfactory.Contract.IsAdmin(&_Itokenfactory.CallOpts, userAddr_)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address userAddr_) view returns(bool)
func (_Itokenfactory *ItokenfactoryCallerSession) IsAdmin(userAddr_ common.Address) (bool, error) {
	return _Itokenfactory.Contract.IsAdmin(&_Itokenfactory.CallOpts, userAddr_)
}

// PoolsBeacon is a free data retrieval call binding the contract method 0xb3f251e7.
//
// Solidity: function poolsBeacon() view returns(address)
func (_Itokenfactory *ItokenfactoryCaller) PoolsBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "poolsBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolsBeacon is a free data retrieval call binding the contract method 0xb3f251e7.
//
// Solidity: function poolsBeacon() view returns(address)
func (_Itokenfactory *ItokenfactorySession) PoolsBeacon() (common.Address, error) {
	return _Itokenfactory.Contract.PoolsBeacon(&_Itokenfactory.CallOpts)
}

// PoolsBeacon is a free data retrieval call binding the contract method 0xb3f251e7.
//
// Solidity: function poolsBeacon() view returns(address)
func (_Itokenfactory *ItokenfactoryCallerSession) PoolsBeacon() (common.Address, error) {
	return _Itokenfactory.Contract.PoolsBeacon(&_Itokenfactory.CallOpts)
}

// PriceDecimals is a free data retrieval call binding the contract method 0x05300b28.
//
// Solidity: function priceDecimals() view returns(uint8)
func (_Itokenfactory *ItokenfactoryCaller) PriceDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "priceDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PriceDecimals is a free data retrieval call binding the contract method 0x05300b28.
//
// Solidity: function priceDecimals() view returns(uint8)
func (_Itokenfactory *ItokenfactorySession) PriceDecimals() (uint8, error) {
	return _Itokenfactory.Contract.PriceDecimals(&_Itokenfactory.CallOpts)
}

// PriceDecimals is a free data retrieval call binding the contract method 0x05300b28.
//
// Solidity: function priceDecimals() view returns(uint8)
func (_Itokenfactory *ItokenfactoryCallerSession) PriceDecimals() (uint8, error) {
	return _Itokenfactory.Contract.PriceDecimals(&_Itokenfactory.CallOpts)
}

// TokenContractByIndex is a free data retrieval call binding the contract method 0xb7e6a3ec.
//
// Solidity: function tokenContractByIndex(uint256 tokenContractId_) view returns(address)
func (_Itokenfactory *ItokenfactoryCaller) TokenContractByIndex(opts *bind.CallOpts, tokenContractId_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Itokenfactory.contract.Call(opts, &out, "tokenContractByIndex", tokenContractId_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenContractByIndex is a free data retrieval call binding the contract method 0xb7e6a3ec.
//
// Solidity: function tokenContractByIndex(uint256 tokenContractId_) view returns(address)
func (_Itokenfactory *ItokenfactorySession) TokenContractByIndex(tokenContractId_ *big.Int) (common.Address, error) {
	return _Itokenfactory.Contract.TokenContractByIndex(&_Itokenfactory.CallOpts, tokenContractId_)
}

// TokenContractByIndex is a free data retrieval call binding the contract method 0xb7e6a3ec.
//
// Solidity: function tokenContractByIndex(uint256 tokenContractId_) view returns(address)
func (_Itokenfactory *ItokenfactoryCallerSession) TokenContractByIndex(tokenContractId_ *big.Int) (common.Address, error) {
	return _Itokenfactory.Contract.TokenContractByIndex(&_Itokenfactory.CallOpts, tokenContractId_)
}

// TokenFactoryInit is a paid mutator transaction binding the contract method 0xae13b76e.
//
// Solidity: function __TokenFactory_init(address[] adminsArr_, string baseTokenContractsURI_, uint8 priceDecimals_) returns()
func (_Itokenfactory *ItokenfactoryTransactor) TokenFactoryInit(opts *bind.TransactOpts, adminsArr_ []common.Address, baseTokenContractsURI_ string, priceDecimals_ uint8) (*types.Transaction, error) {
	return _Itokenfactory.contract.Transact(opts, "__TokenFactory_init", adminsArr_, baseTokenContractsURI_, priceDecimals_)
}

// TokenFactoryInit is a paid mutator transaction binding the contract method 0xae13b76e.
//
// Solidity: function __TokenFactory_init(address[] adminsArr_, string baseTokenContractsURI_, uint8 priceDecimals_) returns()
func (_Itokenfactory *ItokenfactorySession) TokenFactoryInit(adminsArr_ []common.Address, baseTokenContractsURI_ string, priceDecimals_ uint8) (*types.Transaction, error) {
	return _Itokenfactory.Contract.TokenFactoryInit(&_Itokenfactory.TransactOpts, adminsArr_, baseTokenContractsURI_, priceDecimals_)
}

// TokenFactoryInit is a paid mutator transaction binding the contract method 0xae13b76e.
//
// Solidity: function __TokenFactory_init(address[] adminsArr_, string baseTokenContractsURI_, uint8 priceDecimals_) returns()
func (_Itokenfactory *ItokenfactoryTransactorSession) TokenFactoryInit(adminsArr_ []common.Address, baseTokenContractsURI_ string, priceDecimals_ uint8) (*types.Transaction, error) {
	return _Itokenfactory.Contract.TokenFactoryInit(&_Itokenfactory.TransactOpts, adminsArr_, baseTokenContractsURI_, priceDecimals_)
}

// DeployTokenContract is a paid mutator transaction binding the contract method 0xda8aa768.
//
// Solidity: function deployTokenContract(uint256 tokenContractId_, string tokenName_, string tokenSymbol_, uint256 pricePerOneToken_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_Itokenfactory *ItokenfactoryTransactor) DeployTokenContract(opts *bind.TransactOpts, tokenContractId_ *big.Int, tokenName_ string, tokenSymbol_ string, pricePerOneToken_ *big.Int, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Itokenfactory.contract.Transact(opts, "deployTokenContract", tokenContractId_, tokenName_, tokenSymbol_, pricePerOneToken_, r_, s_, v_)
}

// DeployTokenContract is a paid mutator transaction binding the contract method 0xda8aa768.
//
// Solidity: function deployTokenContract(uint256 tokenContractId_, string tokenName_, string tokenSymbol_, uint256 pricePerOneToken_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_Itokenfactory *ItokenfactorySession) DeployTokenContract(tokenContractId_ *big.Int, tokenName_ string, tokenSymbol_ string, pricePerOneToken_ *big.Int, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Itokenfactory.Contract.DeployTokenContract(&_Itokenfactory.TransactOpts, tokenContractId_, tokenName_, tokenSymbol_, pricePerOneToken_, r_, s_, v_)
}

// DeployTokenContract is a paid mutator transaction binding the contract method 0xda8aa768.
//
// Solidity: function deployTokenContract(uint256 tokenContractId_, string tokenName_, string tokenSymbol_, uint256 pricePerOneToken_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_Itokenfactory *ItokenfactoryTransactorSession) DeployTokenContract(tokenContractId_ *big.Int, tokenName_ string, tokenSymbol_ string, pricePerOneToken_ *big.Int, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Itokenfactory.Contract.DeployTokenContract(&_Itokenfactory.TransactOpts, tokenContractId_, tokenName_, tokenSymbol_, pricePerOneToken_, r_, s_, v_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Itokenfactory *ItokenfactoryTransactor) SetBaseTokenContractsURI(opts *bind.TransactOpts, baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Itokenfactory.contract.Transact(opts, "setBaseTokenContractsURI", baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Itokenfactory *ItokenfactorySession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Itokenfactory.Contract.SetBaseTokenContractsURI(&_Itokenfactory.TransactOpts, baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Itokenfactory *ItokenfactoryTransactorSession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Itokenfactory.Contract.SetBaseTokenContractsURI(&_Itokenfactory.TransactOpts, baseTokenContractsURI_)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address newImplementation_) returns()
func (_Itokenfactory *ItokenfactoryTransactor) SetNewImplementation(opts *bind.TransactOpts, newImplementation_ common.Address) (*types.Transaction, error) {
	return _Itokenfactory.contract.Transact(opts, "setNewImplementation", newImplementation_)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address newImplementation_) returns()
func (_Itokenfactory *ItokenfactorySession) SetNewImplementation(newImplementation_ common.Address) (*types.Transaction, error) {
	return _Itokenfactory.Contract.SetNewImplementation(&_Itokenfactory.TransactOpts, newImplementation_)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address newImplementation_) returns()
func (_Itokenfactory *ItokenfactoryTransactorSession) SetNewImplementation(newImplementation_ common.Address) (*types.Transaction, error) {
	return _Itokenfactory.Contract.SetNewImplementation(&_Itokenfactory.TransactOpts, newImplementation_)
}

// UpdateAdmins is a paid mutator transaction binding the contract method 0x809f976a.
//
// Solidity: function updateAdmins(address[] adminsToUpdate_, bool isAdding_) returns()
func (_Itokenfactory *ItokenfactoryTransactor) UpdateAdmins(opts *bind.TransactOpts, adminsToUpdate_ []common.Address, isAdding_ bool) (*types.Transaction, error) {
	return _Itokenfactory.contract.Transact(opts, "updateAdmins", adminsToUpdate_, isAdding_)
}

// UpdateAdmins is a paid mutator transaction binding the contract method 0x809f976a.
//
// Solidity: function updateAdmins(address[] adminsToUpdate_, bool isAdding_) returns()
func (_Itokenfactory *ItokenfactorySession) UpdateAdmins(adminsToUpdate_ []common.Address, isAdding_ bool) (*types.Transaction, error) {
	return _Itokenfactory.Contract.UpdateAdmins(&_Itokenfactory.TransactOpts, adminsToUpdate_, isAdding_)
}

// UpdateAdmins is a paid mutator transaction binding the contract method 0x809f976a.
//
// Solidity: function updateAdmins(address[] adminsToUpdate_, bool isAdding_) returns()
func (_Itokenfactory *ItokenfactoryTransactorSession) UpdateAdmins(adminsToUpdate_ []common.Address, isAdding_ bool) (*types.Transaction, error) {
	return _Itokenfactory.Contract.UpdateAdmins(&_Itokenfactory.TransactOpts, adminsToUpdate_, isAdding_)
}

// ItokenfactoryAdminsUpdatedIterator is returned from FilterAdminsUpdated and is used to iterate over the raw logs and unpacked data for AdminsUpdated events raised by the Itokenfactory contract.
type ItokenfactoryAdminsUpdatedIterator struct {
	Event *ItokenfactoryAdminsUpdated // Event containing the contract specifics and raw log

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
func (it *ItokenfactoryAdminsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItokenfactoryAdminsUpdated)
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
		it.Event = new(ItokenfactoryAdminsUpdated)
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
func (it *ItokenfactoryAdminsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItokenfactoryAdminsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItokenfactoryAdminsUpdated represents a AdminsUpdated event raised by the Itokenfactory contract.
type ItokenfactoryAdminsUpdated struct {
	AdminsToUpdate []common.Address
	IsAdding       bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAdminsUpdated is a free log retrieval operation binding the contract event 0x869da43fad8e93387e0f5b58027ae3bdc4c61756fc5219fa85a2e4f44d3ebea4.
//
// Solidity: event AdminsUpdated(address[] adminsToUpdate, bool isAdding)
func (_Itokenfactory *ItokenfactoryFilterer) FilterAdminsUpdated(opts *bind.FilterOpts) (*ItokenfactoryAdminsUpdatedIterator, error) {

	logs, sub, err := _Itokenfactory.contract.FilterLogs(opts, "AdminsUpdated")
	if err != nil {
		return nil, err
	}
	return &ItokenfactoryAdminsUpdatedIterator{contract: _Itokenfactory.contract, event: "AdminsUpdated", logs: logs, sub: sub}, nil
}

// WatchAdminsUpdated is a free log subscription operation binding the contract event 0x869da43fad8e93387e0f5b58027ae3bdc4c61756fc5219fa85a2e4f44d3ebea4.
//
// Solidity: event AdminsUpdated(address[] adminsToUpdate, bool isAdding)
func (_Itokenfactory *ItokenfactoryFilterer) WatchAdminsUpdated(opts *bind.WatchOpts, sink chan<- *ItokenfactoryAdminsUpdated) (event.Subscription, error) {

	logs, sub, err := _Itokenfactory.contract.WatchLogs(opts, "AdminsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItokenfactoryAdminsUpdated)
				if err := _Itokenfactory.contract.UnpackLog(event, "AdminsUpdated", log); err != nil {
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
func (_Itokenfactory *ItokenfactoryFilterer) ParseAdminsUpdated(log types.Log) (*ItokenfactoryAdminsUpdated, error) {
	event := new(ItokenfactoryAdminsUpdated)
	if err := _Itokenfactory.contract.UnpackLog(event, "AdminsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ItokenfactoryBaseTokenContractsURIUpdatedIterator is returned from FilterBaseTokenContractsURIUpdated and is used to iterate over the raw logs and unpacked data for BaseTokenContractsURIUpdated events raised by the Itokenfactory contract.
type ItokenfactoryBaseTokenContractsURIUpdatedIterator struct {
	Event *ItokenfactoryBaseTokenContractsURIUpdated // Event containing the contract specifics and raw log

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
func (it *ItokenfactoryBaseTokenContractsURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItokenfactoryBaseTokenContractsURIUpdated)
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
		it.Event = new(ItokenfactoryBaseTokenContractsURIUpdated)
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
func (it *ItokenfactoryBaseTokenContractsURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItokenfactoryBaseTokenContractsURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItokenfactoryBaseTokenContractsURIUpdated represents a BaseTokenContractsURIUpdated event raised by the Itokenfactory contract.
type ItokenfactoryBaseTokenContractsURIUpdated struct {
	NewBaseTokenContractsURI string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterBaseTokenContractsURIUpdated is a free log retrieval operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Itokenfactory *ItokenfactoryFilterer) FilterBaseTokenContractsURIUpdated(opts *bind.FilterOpts) (*ItokenfactoryBaseTokenContractsURIUpdatedIterator, error) {

	logs, sub, err := _Itokenfactory.contract.FilterLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return &ItokenfactoryBaseTokenContractsURIUpdatedIterator{contract: _Itokenfactory.contract, event: "BaseTokenContractsURIUpdated", logs: logs, sub: sub}, nil
}

// WatchBaseTokenContractsURIUpdated is a free log subscription operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Itokenfactory *ItokenfactoryFilterer) WatchBaseTokenContractsURIUpdated(opts *bind.WatchOpts, sink chan<- *ItokenfactoryBaseTokenContractsURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Itokenfactory.contract.WatchLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItokenfactoryBaseTokenContractsURIUpdated)
				if err := _Itokenfactory.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
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
func (_Itokenfactory *ItokenfactoryFilterer) ParseBaseTokenContractsURIUpdated(log types.Log) (*ItokenfactoryBaseTokenContractsURIUpdated, error) {
	event := new(ItokenfactoryBaseTokenContractsURIUpdated)
	if err := _Itokenfactory.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ItokenfactoryTokenContractDeployedIterator is returned from FilterTokenContractDeployed and is used to iterate over the raw logs and unpacked data for TokenContractDeployed events raised by the Itokenfactory contract.
type ItokenfactoryTokenContractDeployedIterator struct {
	Event *ItokenfactoryTokenContractDeployed // Event containing the contract specifics and raw log

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
func (it *ItokenfactoryTokenContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItokenfactoryTokenContractDeployed)
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
		it.Event = new(ItokenfactoryTokenContractDeployed)
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
func (it *ItokenfactoryTokenContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItokenfactoryTokenContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItokenfactoryTokenContractDeployed represents a TokenContractDeployed event raised by the Itokenfactory contract.
type ItokenfactoryTokenContractDeployed struct {
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
func (_Itokenfactory *ItokenfactoryFilterer) FilterTokenContractDeployed(opts *bind.FilterOpts) (*ItokenfactoryTokenContractDeployedIterator, error) {

	logs, sub, err := _Itokenfactory.contract.FilterLogs(opts, "TokenContractDeployed")
	if err != nil {
		return nil, err
	}
	return &ItokenfactoryTokenContractDeployedIterator{contract: _Itokenfactory.contract, event: "TokenContractDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenContractDeployed is a free log subscription operation binding the contract event 0x8412389c512c8db70cea087e4ae10396c0fc6d2f2fe1efff6fedd9dc87fc228f.
//
// Solidity: event TokenContractDeployed(uint256 tokenContractId, address newTokenContractAddr, uint256 pricePerOneToken, string tokenName, string tokenSymbol)
func (_Itokenfactory *ItokenfactoryFilterer) WatchTokenContractDeployed(opts *bind.WatchOpts, sink chan<- *ItokenfactoryTokenContractDeployed) (event.Subscription, error) {

	logs, sub, err := _Itokenfactory.contract.WatchLogs(opts, "TokenContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItokenfactoryTokenContractDeployed)
				if err := _Itokenfactory.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
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
func (_Itokenfactory *ItokenfactoryFilterer) ParseTokenContractDeployed(log types.Log) (*ItokenfactoryTokenContractDeployed, error) {
	event := new(ItokenfactoryTokenContractDeployed)
	if err := _Itokenfactory.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
