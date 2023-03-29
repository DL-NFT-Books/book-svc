// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractsregistry

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

// ContractsregistryMetaData contains all meta data concerning the Contractsregistry contract.
var ContractsregistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isProxy\",\"type\":\"bool\"}],\"name\":\"AddedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"RemovedContract\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MARKETPLACE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_MANAGER_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_FACTORY_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_REGISTRY_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__OwnableContractsRegistry_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"addContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"addProxyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketplaceContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyUpgrader\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoleManagerContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenFactoryContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenRegistryContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"hasContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"injectDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"injectDependenciesWithData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"justAddProxyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"removeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"upgradeContractAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractsregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsregistryMetaData.ABI instead.
var ContractsregistryABI = ContractsregistryMetaData.ABI

// Contractsregistry is an auto generated Go binding around an Ethereum contract.
type Contractsregistry struct {
	ContractsregistryCaller     // Read-only binding to the contract
	ContractsregistryTransactor // Write-only binding to the contract
	ContractsregistryFilterer   // Log filterer for contract events
}

// ContractsregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsregistrySession struct {
	Contract     *Contractsregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractsregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsregistryCallerSession struct {
	Contract *ContractsregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractsregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsregistryTransactorSession struct {
	Contract     *ContractsregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractsregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsregistryRaw struct {
	Contract *Contractsregistry // Generic contract binding to access the raw methods on
}

// ContractsregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsregistryCallerRaw struct {
	Contract *ContractsregistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsregistryTransactorRaw struct {
	Contract *ContractsregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractsregistry creates a new instance of Contractsregistry, bound to a specific deployed contract.
func NewContractsregistry(address common.Address, backend bind.ContractBackend) (*Contractsregistry, error) {
	contract, err := bindContractsregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractsregistry{ContractsregistryCaller: ContractsregistryCaller{contract: contract}, ContractsregistryTransactor: ContractsregistryTransactor{contract: contract}, ContractsregistryFilterer: ContractsregistryFilterer{contract: contract}}, nil
}

// NewContractsregistryCaller creates a new read-only instance of Contractsregistry, bound to a specific deployed contract.
func NewContractsregistryCaller(address common.Address, caller bind.ContractCaller) (*ContractsregistryCaller, error) {
	contract, err := bindContractsregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsregistryCaller{contract: contract}, nil
}

// NewContractsregistryTransactor creates a new write-only instance of Contractsregistry, bound to a specific deployed contract.
func NewContractsregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsregistryTransactor, error) {
	contract, err := bindContractsregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsregistryTransactor{contract: contract}, nil
}

// NewContractsregistryFilterer creates a new log filterer instance of Contractsregistry, bound to a specific deployed contract.
func NewContractsregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsregistryFilterer, error) {
	contract, err := bindContractsregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsregistryFilterer{contract: contract}, nil
}

// bindContractsregistry binds a generic wrapper to an already deployed contract.
func bindContractsregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsregistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractsregistry *ContractsregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractsregistry.Contract.ContractsregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractsregistry *ContractsregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractsregistry.Contract.ContractsregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractsregistry *ContractsregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractsregistry.Contract.ContractsregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractsregistry *ContractsregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractsregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractsregistry *ContractsregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractsregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractsregistry *ContractsregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractsregistry.Contract.contract.Transact(opts, method, params...)
}

// MARKETPLACENAME is a free data retrieval call binding the contract method 0x7fc06f1d.
//
// Solidity: function MARKETPLACE_NAME() view returns(string)
func (_Contractsregistry *ContractsregistryCaller) MARKETPLACENAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "MARKETPLACE_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MARKETPLACENAME is a free data retrieval call binding the contract method 0x7fc06f1d.
//
// Solidity: function MARKETPLACE_NAME() view returns(string)
func (_Contractsregistry *ContractsregistrySession) MARKETPLACENAME() (string, error) {
	return _Contractsregistry.Contract.MARKETPLACENAME(&_Contractsregistry.CallOpts)
}

// MARKETPLACENAME is a free data retrieval call binding the contract method 0x7fc06f1d.
//
// Solidity: function MARKETPLACE_NAME() view returns(string)
func (_Contractsregistry *ContractsregistryCallerSession) MARKETPLACENAME() (string, error) {
	return _Contractsregistry.Contract.MARKETPLACENAME(&_Contractsregistry.CallOpts)
}

// ROLEMANAGERNAME is a free data retrieval call binding the contract method 0x3976ca98.
//
// Solidity: function ROLE_MANAGER_NAME() view returns(string)
func (_Contractsregistry *ContractsregistryCaller) ROLEMANAGERNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "ROLE_MANAGER_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROLEMANAGERNAME is a free data retrieval call binding the contract method 0x3976ca98.
//
// Solidity: function ROLE_MANAGER_NAME() view returns(string)
func (_Contractsregistry *ContractsregistrySession) ROLEMANAGERNAME() (string, error) {
	return _Contractsregistry.Contract.ROLEMANAGERNAME(&_Contractsregistry.CallOpts)
}

// ROLEMANAGERNAME is a free data retrieval call binding the contract method 0x3976ca98.
//
// Solidity: function ROLE_MANAGER_NAME() view returns(string)
func (_Contractsregistry *ContractsregistryCallerSession) ROLEMANAGERNAME() (string, error) {
	return _Contractsregistry.Contract.ROLEMANAGERNAME(&_Contractsregistry.CallOpts)
}

// TOKENFACTORYNAME is a free data retrieval call binding the contract method 0x723c56e1.
//
// Solidity: function TOKEN_FACTORY_NAME() view returns(string)
func (_Contractsregistry *ContractsregistryCaller) TOKENFACTORYNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "TOKEN_FACTORY_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TOKENFACTORYNAME is a free data retrieval call binding the contract method 0x723c56e1.
//
// Solidity: function TOKEN_FACTORY_NAME() view returns(string)
func (_Contractsregistry *ContractsregistrySession) TOKENFACTORYNAME() (string, error) {
	return _Contractsregistry.Contract.TOKENFACTORYNAME(&_Contractsregistry.CallOpts)
}

// TOKENFACTORYNAME is a free data retrieval call binding the contract method 0x723c56e1.
//
// Solidity: function TOKEN_FACTORY_NAME() view returns(string)
func (_Contractsregistry *ContractsregistryCallerSession) TOKENFACTORYNAME() (string, error) {
	return _Contractsregistry.Contract.TOKENFACTORYNAME(&_Contractsregistry.CallOpts)
}

// TOKENREGISTRYNAME is a free data retrieval call binding the contract method 0x616c9fb1.
//
// Solidity: function TOKEN_REGISTRY_NAME() view returns(string)
func (_Contractsregistry *ContractsregistryCaller) TOKENREGISTRYNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "TOKEN_REGISTRY_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TOKENREGISTRYNAME is a free data retrieval call binding the contract method 0x616c9fb1.
//
// Solidity: function TOKEN_REGISTRY_NAME() view returns(string)
func (_Contractsregistry *ContractsregistrySession) TOKENREGISTRYNAME() (string, error) {
	return _Contractsregistry.Contract.TOKENREGISTRYNAME(&_Contractsregistry.CallOpts)
}

// TOKENREGISTRYNAME is a free data retrieval call binding the contract method 0x616c9fb1.
//
// Solidity: function TOKEN_REGISTRY_NAME() view returns(string)
func (_Contractsregistry *ContractsregistryCallerSession) TOKENREGISTRYNAME() (string, error) {
	return _Contractsregistry.Contract.TOKENREGISTRYNAME(&_Contractsregistry.CallOpts)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name_) view returns(address)
func (_Contractsregistry *ContractsregistryCaller) GetContract(opts *bind.CallOpts, name_ string) (common.Address, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "getContract", name_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name_) view returns(address)
func (_Contractsregistry *ContractsregistrySession) GetContract(name_ string) (common.Address, error) {
	return _Contractsregistry.Contract.GetContract(&_Contractsregistry.CallOpts, name_)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name_) view returns(address)
func (_Contractsregistry *ContractsregistryCallerSession) GetContract(name_ string) (common.Address, error) {
	return _Contractsregistry.Contract.GetContract(&_Contractsregistry.CallOpts, name_)
}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name_) view returns(address)
func (_Contractsregistry *ContractsregistryCaller) GetImplementation(opts *bind.CallOpts, name_ string) (common.Address, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "getImplementation", name_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name_) view returns(address)
func (_Contractsregistry *ContractsregistrySession) GetImplementation(name_ string) (common.Address, error) {
	return _Contractsregistry.Contract.GetImplementation(&_Contractsregistry.CallOpts, name_)
}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name_) view returns(address)
func (_Contractsregistry *ContractsregistryCallerSession) GetImplementation(name_ string) (common.Address, error) {
	return _Contractsregistry.Contract.GetImplementation(&_Contractsregistry.CallOpts, name_)
}

// GetMarketplaceContract is a free data retrieval call binding the contract method 0xaff46572.
//
// Solidity: function getMarketplaceContract() view returns(address)
func (_Contractsregistry *ContractsregistryCaller) GetMarketplaceContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "getMarketplaceContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMarketplaceContract is a free data retrieval call binding the contract method 0xaff46572.
//
// Solidity: function getMarketplaceContract() view returns(address)
func (_Contractsregistry *ContractsregistrySession) GetMarketplaceContract() (common.Address, error) {
	return _Contractsregistry.Contract.GetMarketplaceContract(&_Contractsregistry.CallOpts)
}

// GetMarketplaceContract is a free data retrieval call binding the contract method 0xaff46572.
//
// Solidity: function getMarketplaceContract() view returns(address)
func (_Contractsregistry *ContractsregistryCallerSession) GetMarketplaceContract() (common.Address, error) {
	return _Contractsregistry.Contract.GetMarketplaceContract(&_Contractsregistry.CallOpts)
}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_Contractsregistry *ContractsregistryCaller) GetProxyUpgrader(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "getProxyUpgrader")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_Contractsregistry *ContractsregistrySession) GetProxyUpgrader() (common.Address, error) {
	return _Contractsregistry.Contract.GetProxyUpgrader(&_Contractsregistry.CallOpts)
}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_Contractsregistry *ContractsregistryCallerSession) GetProxyUpgrader() (common.Address, error) {
	return _Contractsregistry.Contract.GetProxyUpgrader(&_Contractsregistry.CallOpts)
}

// GetRoleManagerContract is a free data retrieval call binding the contract method 0xef3454de.
//
// Solidity: function getRoleManagerContract() view returns(address)
func (_Contractsregistry *ContractsregistryCaller) GetRoleManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "getRoleManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleManagerContract is a free data retrieval call binding the contract method 0xef3454de.
//
// Solidity: function getRoleManagerContract() view returns(address)
func (_Contractsregistry *ContractsregistrySession) GetRoleManagerContract() (common.Address, error) {
	return _Contractsregistry.Contract.GetRoleManagerContract(&_Contractsregistry.CallOpts)
}

// GetRoleManagerContract is a free data retrieval call binding the contract method 0xef3454de.
//
// Solidity: function getRoleManagerContract() view returns(address)
func (_Contractsregistry *ContractsregistryCallerSession) GetRoleManagerContract() (common.Address, error) {
	return _Contractsregistry.Contract.GetRoleManagerContract(&_Contractsregistry.CallOpts)
}

// GetTokenFactoryContract is a free data retrieval call binding the contract method 0x1a271056.
//
// Solidity: function getTokenFactoryContract() view returns(address)
func (_Contractsregistry *ContractsregistryCaller) GetTokenFactoryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "getTokenFactoryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenFactoryContract is a free data retrieval call binding the contract method 0x1a271056.
//
// Solidity: function getTokenFactoryContract() view returns(address)
func (_Contractsregistry *ContractsregistrySession) GetTokenFactoryContract() (common.Address, error) {
	return _Contractsregistry.Contract.GetTokenFactoryContract(&_Contractsregistry.CallOpts)
}

// GetTokenFactoryContract is a free data retrieval call binding the contract method 0x1a271056.
//
// Solidity: function getTokenFactoryContract() view returns(address)
func (_Contractsregistry *ContractsregistryCallerSession) GetTokenFactoryContract() (common.Address, error) {
	return _Contractsregistry.Contract.GetTokenFactoryContract(&_Contractsregistry.CallOpts)
}

// GetTokenRegistryContract is a free data retrieval call binding the contract method 0x3e61489f.
//
// Solidity: function getTokenRegistryContract() view returns(address)
func (_Contractsregistry *ContractsregistryCaller) GetTokenRegistryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "getTokenRegistryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenRegistryContract is a free data retrieval call binding the contract method 0x3e61489f.
//
// Solidity: function getTokenRegistryContract() view returns(address)
func (_Contractsregistry *ContractsregistrySession) GetTokenRegistryContract() (common.Address, error) {
	return _Contractsregistry.Contract.GetTokenRegistryContract(&_Contractsregistry.CallOpts)
}

// GetTokenRegistryContract is a free data retrieval call binding the contract method 0x3e61489f.
//
// Solidity: function getTokenRegistryContract() view returns(address)
func (_Contractsregistry *ContractsregistryCallerSession) GetTokenRegistryContract() (common.Address, error) {
	return _Contractsregistry.Contract.GetTokenRegistryContract(&_Contractsregistry.CallOpts)
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name_) view returns(bool)
func (_Contractsregistry *ContractsregistryCaller) HasContract(opts *bind.CallOpts, name_ string) (bool, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "hasContract", name_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name_) view returns(bool)
func (_Contractsregistry *ContractsregistrySession) HasContract(name_ string) (bool, error) {
	return _Contractsregistry.Contract.HasContract(&_Contractsregistry.CallOpts, name_)
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name_) view returns(bool)
func (_Contractsregistry *ContractsregistryCallerSession) HasContract(name_ string) (bool, error) {
	return _Contractsregistry.Contract.HasContract(&_Contractsregistry.CallOpts, name_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractsregistry *ContractsregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractsregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractsregistry *ContractsregistrySession) Owner() (common.Address, error) {
	return _Contractsregistry.Contract.Owner(&_Contractsregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractsregistry *ContractsregistryCallerSession) Owner() (common.Address, error) {
	return _Contractsregistry.Contract.Owner(&_Contractsregistry.CallOpts)
}

// OwnableContractsRegistryInit is a paid mutator transaction binding the contract method 0x1bae7a65.
//
// Solidity: function __OwnableContractsRegistry_init() returns()
func (_Contractsregistry *ContractsregistryTransactor) OwnableContractsRegistryInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "__OwnableContractsRegistry_init")
}

// OwnableContractsRegistryInit is a paid mutator transaction binding the contract method 0x1bae7a65.
//
// Solidity: function __OwnableContractsRegistry_init() returns()
func (_Contractsregistry *ContractsregistrySession) OwnableContractsRegistryInit() (*types.Transaction, error) {
	return _Contractsregistry.Contract.OwnableContractsRegistryInit(&_Contractsregistry.TransactOpts)
}

// OwnableContractsRegistryInit is a paid mutator transaction binding the contract method 0x1bae7a65.
//
// Solidity: function __OwnableContractsRegistry_init() returns()
func (_Contractsregistry *ContractsregistryTransactorSession) OwnableContractsRegistryInit() (*types.Transaction, error) {
	return _Contractsregistry.Contract.OwnableContractsRegistryInit(&_Contractsregistry.TransactOpts)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_Contractsregistry *ContractsregistryTransactor) AddContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "addContract", name_, contractAddress_)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_Contractsregistry *ContractsregistrySession) AddContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.Contract.AddContract(&_Contractsregistry.TransactOpts, name_, contractAddress_)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_Contractsregistry *ContractsregistryTransactorSession) AddContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.Contract.AddContract(&_Contractsregistry.TransactOpts, name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_Contractsregistry *ContractsregistryTransactor) AddProxyContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "addProxyContract", name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_Contractsregistry *ContractsregistrySession) AddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.Contract.AddProxyContract(&_Contractsregistry.TransactOpts, name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_Contractsregistry *ContractsregistryTransactorSession) AddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.Contract.AddProxyContract(&_Contractsregistry.TransactOpts, name_, contractAddress_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_Contractsregistry *ContractsregistryTransactor) InjectDependencies(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "injectDependencies", name_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_Contractsregistry *ContractsregistrySession) InjectDependencies(name_ string) (*types.Transaction, error) {
	return _Contractsregistry.Contract.InjectDependencies(&_Contractsregistry.TransactOpts, name_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_Contractsregistry *ContractsregistryTransactorSession) InjectDependencies(name_ string) (*types.Transaction, error) {
	return _Contractsregistry.Contract.InjectDependencies(&_Contractsregistry.TransactOpts, name_)
}

// InjectDependenciesWithData is a paid mutator transaction binding the contract method 0xbe96dc3e.
//
// Solidity: function injectDependenciesWithData(string name_, bytes data_) returns()
func (_Contractsregistry *ContractsregistryTransactor) InjectDependenciesWithData(opts *bind.TransactOpts, name_ string, data_ []byte) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "injectDependenciesWithData", name_, data_)
}

// InjectDependenciesWithData is a paid mutator transaction binding the contract method 0xbe96dc3e.
//
// Solidity: function injectDependenciesWithData(string name_, bytes data_) returns()
func (_Contractsregistry *ContractsregistrySession) InjectDependenciesWithData(name_ string, data_ []byte) (*types.Transaction, error) {
	return _Contractsregistry.Contract.InjectDependenciesWithData(&_Contractsregistry.TransactOpts, name_, data_)
}

// InjectDependenciesWithData is a paid mutator transaction binding the contract method 0xbe96dc3e.
//
// Solidity: function injectDependenciesWithData(string name_, bytes data_) returns()
func (_Contractsregistry *ContractsregistryTransactorSession) InjectDependenciesWithData(name_ string, data_ []byte) (*types.Transaction, error) {
	return _Contractsregistry.Contract.InjectDependenciesWithData(&_Contractsregistry.TransactOpts, name_, data_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_Contractsregistry *ContractsregistryTransactor) JustAddProxyContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "justAddProxyContract", name_, contractAddress_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_Contractsregistry *ContractsregistrySession) JustAddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.Contract.JustAddProxyContract(&_Contractsregistry.TransactOpts, name_, contractAddress_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_Contractsregistry *ContractsregistryTransactorSession) JustAddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.Contract.JustAddProxyContract(&_Contractsregistry.TransactOpts, name_, contractAddress_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_Contractsregistry *ContractsregistryTransactor) RemoveContract(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "removeContract", name_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_Contractsregistry *ContractsregistrySession) RemoveContract(name_ string) (*types.Transaction, error) {
	return _Contractsregistry.Contract.RemoveContract(&_Contractsregistry.TransactOpts, name_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_Contractsregistry *ContractsregistryTransactorSession) RemoveContract(name_ string) (*types.Transaction, error) {
	return _Contractsregistry.Contract.RemoveContract(&_Contractsregistry.TransactOpts, name_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contractsregistry *ContractsregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contractsregistry *ContractsregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Contractsregistry.Contract.RenounceOwnership(&_Contractsregistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contractsregistry *ContractsregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contractsregistry.Contract.RenounceOwnership(&_Contractsregistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contractsregistry *ContractsregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contractsregistry *ContractsregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contractsregistry.Contract.TransferOwnership(&_Contractsregistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contractsregistry *ContractsregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contractsregistry.Contract.TransferOwnership(&_Contractsregistry.TransactOpts, newOwner)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_Contractsregistry *ContractsregistryTransactor) UpgradeContract(opts *bind.TransactOpts, name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "upgradeContract", name_, newImplementation_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_Contractsregistry *ContractsregistrySession) UpgradeContract(name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.Contract.UpgradeContract(&_Contractsregistry.TransactOpts, name_, newImplementation_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_Contractsregistry *ContractsregistryTransactorSession) UpgradeContract(name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _Contractsregistry.Contract.UpgradeContract(&_Contractsregistry.TransactOpts, name_, newImplementation_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_Contractsregistry *ContractsregistryTransactor) UpgradeContractAndCall(opts *bind.TransactOpts, name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Contractsregistry.contract.Transact(opts, "upgradeContractAndCall", name_, newImplementation_, data_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_Contractsregistry *ContractsregistrySession) UpgradeContractAndCall(name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Contractsregistry.Contract.UpgradeContractAndCall(&_Contractsregistry.TransactOpts, name_, newImplementation_, data_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_Contractsregistry *ContractsregistryTransactorSession) UpgradeContractAndCall(name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Contractsregistry.Contract.UpgradeContractAndCall(&_Contractsregistry.TransactOpts, name_, newImplementation_, data_)
}

// ContractsregistryAddedContractIterator is returned from FilterAddedContract and is used to iterate over the raw logs and unpacked data for AddedContract events raised by the Contractsregistry contract.
type ContractsregistryAddedContractIterator struct {
	Event *ContractsregistryAddedContract // Event containing the contract specifics and raw log

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
func (it *ContractsregistryAddedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsregistryAddedContract)
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
		it.Event = new(ContractsregistryAddedContract)
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
func (it *ContractsregistryAddedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsregistryAddedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsregistryAddedContract represents a AddedContract event raised by the Contractsregistry contract.
type ContractsregistryAddedContract struct {
	Name            string
	ContractAddress common.Address
	IsProxy         bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddedContract is a free log retrieval operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_Contractsregistry *ContractsregistryFilterer) FilterAddedContract(opts *bind.FilterOpts) (*ContractsregistryAddedContractIterator, error) {

	logs, sub, err := _Contractsregistry.contract.FilterLogs(opts, "AddedContract")
	if err != nil {
		return nil, err
	}
	return &ContractsregistryAddedContractIterator{contract: _Contractsregistry.contract, event: "AddedContract", logs: logs, sub: sub}, nil
}

// WatchAddedContract is a free log subscription operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_Contractsregistry *ContractsregistryFilterer) WatchAddedContract(opts *bind.WatchOpts, sink chan<- *ContractsregistryAddedContract) (event.Subscription, error) {

	logs, sub, err := _Contractsregistry.contract.WatchLogs(opts, "AddedContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsregistryAddedContract)
				if err := _Contractsregistry.contract.UnpackLog(event, "AddedContract", log); err != nil {
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

// ParseAddedContract is a log parse operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_Contractsregistry *ContractsregistryFilterer) ParseAddedContract(log types.Log) (*ContractsregistryAddedContract, error) {
	event := new(ContractsregistryAddedContract)
	if err := _Contractsregistry.contract.UnpackLog(event, "AddedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsregistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contractsregistry contract.
type ContractsregistryOwnershipTransferredIterator struct {
	Event *ContractsregistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsregistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsregistryOwnershipTransferred)
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
		it.Event = new(ContractsregistryOwnershipTransferred)
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
func (it *ContractsregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsregistryOwnershipTransferred represents a OwnershipTransferred event raised by the Contractsregistry contract.
type ContractsregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contractsregistry *ContractsregistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsregistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contractsregistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsregistryOwnershipTransferredIterator{contract: _Contractsregistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contractsregistry *ContractsregistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsregistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contractsregistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsregistryOwnershipTransferred)
				if err := _Contractsregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contractsregistry *ContractsregistryFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsregistryOwnershipTransferred, error) {
	event := new(ContractsregistryOwnershipTransferred)
	if err := _Contractsregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsregistryRemovedContractIterator is returned from FilterRemovedContract and is used to iterate over the raw logs and unpacked data for RemovedContract events raised by the Contractsregistry contract.
type ContractsregistryRemovedContractIterator struct {
	Event *ContractsregistryRemovedContract // Event containing the contract specifics and raw log

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
func (it *ContractsregistryRemovedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsregistryRemovedContract)
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
		it.Event = new(ContractsregistryRemovedContract)
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
func (it *ContractsregistryRemovedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsregistryRemovedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsregistryRemovedContract represents a RemovedContract event raised by the Contractsregistry contract.
type ContractsregistryRemovedContract struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedContract is a free log retrieval operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_Contractsregistry *ContractsregistryFilterer) FilterRemovedContract(opts *bind.FilterOpts) (*ContractsregistryRemovedContractIterator, error) {

	logs, sub, err := _Contractsregistry.contract.FilterLogs(opts, "RemovedContract")
	if err != nil {
		return nil, err
	}
	return &ContractsregistryRemovedContractIterator{contract: _Contractsregistry.contract, event: "RemovedContract", logs: logs, sub: sub}, nil
}

// WatchRemovedContract is a free log subscription operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_Contractsregistry *ContractsregistryFilterer) WatchRemovedContract(opts *bind.WatchOpts, sink chan<- *ContractsregistryRemovedContract) (event.Subscription, error) {

	logs, sub, err := _Contractsregistry.contract.WatchLogs(opts, "RemovedContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsregistryRemovedContract)
				if err := _Contractsregistry.contract.UnpackLog(event, "RemovedContract", log); err != nil {
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

// ParseRemovedContract is a log parse operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_Contractsregistry *ContractsregistryFilterer) ParseRemovedContract(log types.Log) (*ContractsregistryRemovedContract, error) {
	event := new(ContractsregistryRemovedContract)
	if err := _Contractsregistry.contract.UnpackLog(event, "RemovedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
