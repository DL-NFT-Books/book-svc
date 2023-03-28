// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rolemanager

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

// RolemanagerMetaData contains all meta data concerning the Rolemanager contract.
var RolemanagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMINISTRATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKETPLACE_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_SUPERVISOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_FACTORY_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_REGISTRY_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__RoleManager_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles_\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts_\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager_\",\"type\":\"address\"}],\"name\":\"isMarketplaceManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"supervisor_\",\"type\":\"address\"}],\"name\":\"isRoleSupervisor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager_\",\"type\":\"address\"}],\"name\":\"isTokenFactoryManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager_\",\"type\":\"address\"}],\"name\":\"isTokenManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager_\",\"type\":\"address\"}],\"name\":\"isTokenRegistryManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager_\",\"type\":\"address\"}],\"name\":\"isWithdrawalManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractsRegistry_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RolemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use RolemanagerMetaData.ABI instead.
var RolemanagerABI = RolemanagerMetaData.ABI

// Rolemanager is an auto generated Go binding around an Ethereum contract.
type Rolemanager struct {
	RolemanagerCaller     // Read-only binding to the contract
	RolemanagerTransactor // Write-only binding to the contract
	RolemanagerFilterer   // Log filterer for contract events
}

// RolemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RolemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RolemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RolemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RolemanagerSession struct {
	Contract     *Rolemanager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RolemanagerCallerSession struct {
	Contract *RolemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RolemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RolemanagerTransactorSession struct {
	Contract     *RolemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RolemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RolemanagerRaw struct {
	Contract *Rolemanager // Generic contract binding to access the raw methods on
}

// RolemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RolemanagerCallerRaw struct {
	Contract *RolemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// RolemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RolemanagerTransactorRaw struct {
	Contract *RolemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRolemanager creates a new instance of Rolemanager, bound to a specific deployed contract.
func NewRolemanager(address common.Address, backend bind.ContractBackend) (*Rolemanager, error) {
	contract, err := bindRolemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rolemanager{RolemanagerCaller: RolemanagerCaller{contract: contract}, RolemanagerTransactor: RolemanagerTransactor{contract: contract}, RolemanagerFilterer: RolemanagerFilterer{contract: contract}}, nil
}

// NewRolemanagerCaller creates a new read-only instance of Rolemanager, bound to a specific deployed contract.
func NewRolemanagerCaller(address common.Address, caller bind.ContractCaller) (*RolemanagerCaller, error) {
	contract, err := bindRolemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolemanagerCaller{contract: contract}, nil
}

// NewRolemanagerTransactor creates a new write-only instance of Rolemanager, bound to a specific deployed contract.
func NewRolemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RolemanagerTransactor, error) {
	contract, err := bindRolemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolemanagerTransactor{contract: contract}, nil
}

// NewRolemanagerFilterer creates a new log filterer instance of Rolemanager, bound to a specific deployed contract.
func NewRolemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RolemanagerFilterer, error) {
	contract, err := bindRolemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolemanagerFilterer{contract: contract}, nil
}

// bindRolemanager binds a generic wrapper to an already deployed contract.
func bindRolemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolemanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rolemanager *RolemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rolemanager.Contract.RolemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rolemanager *RolemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rolemanager.Contract.RolemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rolemanager *RolemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rolemanager.Contract.RolemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rolemanager *RolemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rolemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rolemanager *RolemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rolemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rolemanager *RolemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rolemanager.Contract.contract.Transact(opts, method, params...)
}

// ADMINISTRATORROLE is a free data retrieval call binding the contract method 0xf45edb5f.
//
// Solidity: function ADMINISTRATOR_ROLE() view returns(bytes32)
func (_Rolemanager *RolemanagerCaller) ADMINISTRATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "ADMINISTRATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINISTRATORROLE is a free data retrieval call binding the contract method 0xf45edb5f.
//
// Solidity: function ADMINISTRATOR_ROLE() view returns(bytes32)
func (_Rolemanager *RolemanagerSession) ADMINISTRATORROLE() ([32]byte, error) {
	return _Rolemanager.Contract.ADMINISTRATORROLE(&_Rolemanager.CallOpts)
}

// ADMINISTRATORROLE is a free data retrieval call binding the contract method 0xf45edb5f.
//
// Solidity: function ADMINISTRATOR_ROLE() view returns(bytes32)
func (_Rolemanager *RolemanagerCallerSession) ADMINISTRATORROLE() ([32]byte, error) {
	return _Rolemanager.Contract.ADMINISTRATORROLE(&_Rolemanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Rolemanager *RolemanagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Rolemanager *RolemanagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Rolemanager.Contract.DEFAULTADMINROLE(&_Rolemanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Rolemanager *RolemanagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Rolemanager.Contract.DEFAULTADMINROLE(&_Rolemanager.CallOpts)
}

// MARKETPLACEMANAGER is a free data retrieval call binding the contract method 0xce2c2940.
//
// Solidity: function MARKETPLACE_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerCaller) MARKETPLACEMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "MARKETPLACE_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MARKETPLACEMANAGER is a free data retrieval call binding the contract method 0xce2c2940.
//
// Solidity: function MARKETPLACE_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerSession) MARKETPLACEMANAGER() ([32]byte, error) {
	return _Rolemanager.Contract.MARKETPLACEMANAGER(&_Rolemanager.CallOpts)
}

// MARKETPLACEMANAGER is a free data retrieval call binding the contract method 0xce2c2940.
//
// Solidity: function MARKETPLACE_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerCallerSession) MARKETPLACEMANAGER() ([32]byte, error) {
	return _Rolemanager.Contract.MARKETPLACEMANAGER(&_Rolemanager.CallOpts)
}

// ROLESUPERVISOR is a free data retrieval call binding the contract method 0x0d80af9b.
//
// Solidity: function ROLE_SUPERVISOR() view returns(bytes32)
func (_Rolemanager *RolemanagerCaller) ROLESUPERVISOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "ROLE_SUPERVISOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLESUPERVISOR is a free data retrieval call binding the contract method 0x0d80af9b.
//
// Solidity: function ROLE_SUPERVISOR() view returns(bytes32)
func (_Rolemanager *RolemanagerSession) ROLESUPERVISOR() ([32]byte, error) {
	return _Rolemanager.Contract.ROLESUPERVISOR(&_Rolemanager.CallOpts)
}

// ROLESUPERVISOR is a free data retrieval call binding the contract method 0x0d80af9b.
//
// Solidity: function ROLE_SUPERVISOR() view returns(bytes32)
func (_Rolemanager *RolemanagerCallerSession) ROLESUPERVISOR() ([32]byte, error) {
	return _Rolemanager.Contract.ROLESUPERVISOR(&_Rolemanager.CallOpts)
}

// TOKENFACTORYMANAGER is a free data retrieval call binding the contract method 0xfb70f02f.
//
// Solidity: function TOKEN_FACTORY_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerCaller) TOKENFACTORYMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "TOKEN_FACTORY_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENFACTORYMANAGER is a free data retrieval call binding the contract method 0xfb70f02f.
//
// Solidity: function TOKEN_FACTORY_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerSession) TOKENFACTORYMANAGER() ([32]byte, error) {
	return _Rolemanager.Contract.TOKENFACTORYMANAGER(&_Rolemanager.CallOpts)
}

// TOKENFACTORYMANAGER is a free data retrieval call binding the contract method 0xfb70f02f.
//
// Solidity: function TOKEN_FACTORY_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerCallerSession) TOKENFACTORYMANAGER() ([32]byte, error) {
	return _Rolemanager.Contract.TOKENFACTORYMANAGER(&_Rolemanager.CallOpts)
}

// TOKENMANAGER is a free data retrieval call binding the contract method 0xe0956e0f.
//
// Solidity: function TOKEN_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerCaller) TOKENMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "TOKEN_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENMANAGER is a free data retrieval call binding the contract method 0xe0956e0f.
//
// Solidity: function TOKEN_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerSession) TOKENMANAGER() ([32]byte, error) {
	return _Rolemanager.Contract.TOKENMANAGER(&_Rolemanager.CallOpts)
}

// TOKENMANAGER is a free data retrieval call binding the contract method 0xe0956e0f.
//
// Solidity: function TOKEN_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerCallerSession) TOKENMANAGER() ([32]byte, error) {
	return _Rolemanager.Contract.TOKENMANAGER(&_Rolemanager.CallOpts)
}

// TOKENREGISTRYMANAGER is a free data retrieval call binding the contract method 0x6e3673bb.
//
// Solidity: function TOKEN_REGISTRY_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerCaller) TOKENREGISTRYMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "TOKEN_REGISTRY_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENREGISTRYMANAGER is a free data retrieval call binding the contract method 0x6e3673bb.
//
// Solidity: function TOKEN_REGISTRY_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerSession) TOKENREGISTRYMANAGER() ([32]byte, error) {
	return _Rolemanager.Contract.TOKENREGISTRYMANAGER(&_Rolemanager.CallOpts)
}

// TOKENREGISTRYMANAGER is a free data retrieval call binding the contract method 0x6e3673bb.
//
// Solidity: function TOKEN_REGISTRY_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerCallerSession) TOKENREGISTRYMANAGER() ([32]byte, error) {
	return _Rolemanager.Contract.TOKENREGISTRYMANAGER(&_Rolemanager.CallOpts)
}

// WITHDRAWALMANAGER is a free data retrieval call binding the contract method 0xd2f03194.
//
// Solidity: function WITHDRAWAL_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerCaller) WITHDRAWALMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "WITHDRAWAL_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWALMANAGER is a free data retrieval call binding the contract method 0xd2f03194.
//
// Solidity: function WITHDRAWAL_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerSession) WITHDRAWALMANAGER() ([32]byte, error) {
	return _Rolemanager.Contract.WITHDRAWALMANAGER(&_Rolemanager.CallOpts)
}

// WITHDRAWALMANAGER is a free data retrieval call binding the contract method 0xd2f03194.
//
// Solidity: function WITHDRAWAL_MANAGER() view returns(bytes32)
func (_Rolemanager *RolemanagerCallerSession) WITHDRAWALMANAGER() ([32]byte, error) {
	return _Rolemanager.Contract.WITHDRAWALMANAGER(&_Rolemanager.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Rolemanager *RolemanagerCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Rolemanager *RolemanagerSession) GetInjector() (common.Address, error) {
	return _Rolemanager.Contract.GetInjector(&_Rolemanager.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Rolemanager *RolemanagerCallerSession) GetInjector() (common.Address, error) {
	return _Rolemanager.Contract.GetInjector(&_Rolemanager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Rolemanager *RolemanagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Rolemanager *RolemanagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Rolemanager.Contract.GetRoleAdmin(&_Rolemanager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Rolemanager *RolemanagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Rolemanager.Contract.GetRoleAdmin(&_Rolemanager.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Rolemanager *RolemanagerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Rolemanager *RolemanagerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Rolemanager.Contract.GetRoleMember(&_Rolemanager.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Rolemanager *RolemanagerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Rolemanager.Contract.GetRoleMember(&_Rolemanager.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Rolemanager *RolemanagerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Rolemanager *RolemanagerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Rolemanager.Contract.GetRoleMemberCount(&_Rolemanager.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Rolemanager *RolemanagerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Rolemanager.Contract.GetRoleMemberCount(&_Rolemanager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Rolemanager *RolemanagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Rolemanager *RolemanagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Rolemanager.Contract.HasRole(&_Rolemanager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Rolemanager *RolemanagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Rolemanager.Contract.HasRole(&_Rolemanager.CallOpts, role, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin_) view returns(bool)
func (_Rolemanager *RolemanagerCaller) IsAdmin(opts *bind.CallOpts, admin_ common.Address) (bool, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "isAdmin", admin_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin_) view returns(bool)
func (_Rolemanager *RolemanagerSession) IsAdmin(admin_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsAdmin(&_Rolemanager.CallOpts, admin_)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address admin_) view returns(bool)
func (_Rolemanager *RolemanagerCallerSession) IsAdmin(admin_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsAdmin(&_Rolemanager.CallOpts, admin_)
}

// IsMarketplaceManager is a free data retrieval call binding the contract method 0x2019e9cb.
//
// Solidity: function isMarketplaceManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerCaller) IsMarketplaceManager(opts *bind.CallOpts, manager_ common.Address) (bool, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "isMarketplaceManager", manager_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketplaceManager is a free data retrieval call binding the contract method 0x2019e9cb.
//
// Solidity: function isMarketplaceManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerSession) IsMarketplaceManager(manager_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsMarketplaceManager(&_Rolemanager.CallOpts, manager_)
}

// IsMarketplaceManager is a free data retrieval call binding the contract method 0x2019e9cb.
//
// Solidity: function isMarketplaceManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerCallerSession) IsMarketplaceManager(manager_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsMarketplaceManager(&_Rolemanager.CallOpts, manager_)
}

// IsRoleSupervisor is a free data retrieval call binding the contract method 0xe3e941b3.
//
// Solidity: function isRoleSupervisor(address supervisor_) view returns(bool)
func (_Rolemanager *RolemanagerCaller) IsRoleSupervisor(opts *bind.CallOpts, supervisor_ common.Address) (bool, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "isRoleSupervisor", supervisor_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRoleSupervisor is a free data retrieval call binding the contract method 0xe3e941b3.
//
// Solidity: function isRoleSupervisor(address supervisor_) view returns(bool)
func (_Rolemanager *RolemanagerSession) IsRoleSupervisor(supervisor_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsRoleSupervisor(&_Rolemanager.CallOpts, supervisor_)
}

// IsRoleSupervisor is a free data retrieval call binding the contract method 0xe3e941b3.
//
// Solidity: function isRoleSupervisor(address supervisor_) view returns(bool)
func (_Rolemanager *RolemanagerCallerSession) IsRoleSupervisor(supervisor_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsRoleSupervisor(&_Rolemanager.CallOpts, supervisor_)
}

// IsTokenFactoryManager is a free data retrieval call binding the contract method 0xde94b229.
//
// Solidity: function isTokenFactoryManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerCaller) IsTokenFactoryManager(opts *bind.CallOpts, manager_ common.Address) (bool, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "isTokenFactoryManager", manager_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenFactoryManager is a free data retrieval call binding the contract method 0xde94b229.
//
// Solidity: function isTokenFactoryManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerSession) IsTokenFactoryManager(manager_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsTokenFactoryManager(&_Rolemanager.CallOpts, manager_)
}

// IsTokenFactoryManager is a free data retrieval call binding the contract method 0xde94b229.
//
// Solidity: function isTokenFactoryManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerCallerSession) IsTokenFactoryManager(manager_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsTokenFactoryManager(&_Rolemanager.CallOpts, manager_)
}

// IsTokenManager is a free data retrieval call binding the contract method 0xebc26119.
//
// Solidity: function isTokenManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerCaller) IsTokenManager(opts *bind.CallOpts, manager_ common.Address) (bool, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "isTokenManager", manager_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenManager is a free data retrieval call binding the contract method 0xebc26119.
//
// Solidity: function isTokenManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerSession) IsTokenManager(manager_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsTokenManager(&_Rolemanager.CallOpts, manager_)
}

// IsTokenManager is a free data retrieval call binding the contract method 0xebc26119.
//
// Solidity: function isTokenManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerCallerSession) IsTokenManager(manager_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsTokenManager(&_Rolemanager.CallOpts, manager_)
}

// IsTokenRegistryManager is a free data retrieval call binding the contract method 0xccc8843e.
//
// Solidity: function isTokenRegistryManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerCaller) IsTokenRegistryManager(opts *bind.CallOpts, manager_ common.Address) (bool, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "isTokenRegistryManager", manager_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenRegistryManager is a free data retrieval call binding the contract method 0xccc8843e.
//
// Solidity: function isTokenRegistryManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerSession) IsTokenRegistryManager(manager_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsTokenRegistryManager(&_Rolemanager.CallOpts, manager_)
}

// IsTokenRegistryManager is a free data retrieval call binding the contract method 0xccc8843e.
//
// Solidity: function isTokenRegistryManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerCallerSession) IsTokenRegistryManager(manager_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsTokenRegistryManager(&_Rolemanager.CallOpts, manager_)
}

// IsWithdrawalManager is a free data retrieval call binding the contract method 0x5cea8645.
//
// Solidity: function isWithdrawalManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerCaller) IsWithdrawalManager(opts *bind.CallOpts, manager_ common.Address) (bool, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "isWithdrawalManager", manager_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalManager is a free data retrieval call binding the contract method 0x5cea8645.
//
// Solidity: function isWithdrawalManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerSession) IsWithdrawalManager(manager_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsWithdrawalManager(&_Rolemanager.CallOpts, manager_)
}

// IsWithdrawalManager is a free data retrieval call binding the contract method 0x5cea8645.
//
// Solidity: function isWithdrawalManager(address manager_) view returns(bool)
func (_Rolemanager *RolemanagerCallerSession) IsWithdrawalManager(manager_ common.Address) (bool, error) {
	return _Rolemanager.Contract.IsWithdrawalManager(&_Rolemanager.CallOpts, manager_)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Rolemanager *RolemanagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Rolemanager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Rolemanager *RolemanagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Rolemanager.Contract.SupportsInterface(&_Rolemanager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Rolemanager *RolemanagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Rolemanager.Contract.SupportsInterface(&_Rolemanager.CallOpts, interfaceId)
}

// RoleManagerInit is a paid mutator transaction binding the contract method 0x3af4a5e4.
//
// Solidity: function __RoleManager_init() returns()
func (_Rolemanager *RolemanagerTransactor) RoleManagerInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rolemanager.contract.Transact(opts, "__RoleManager_init")
}

// RoleManagerInit is a paid mutator transaction binding the contract method 0x3af4a5e4.
//
// Solidity: function __RoleManager_init() returns()
func (_Rolemanager *RolemanagerSession) RoleManagerInit() (*types.Transaction, error) {
	return _Rolemanager.Contract.RoleManagerInit(&_Rolemanager.TransactOpts)
}

// RoleManagerInit is a paid mutator transaction binding the contract method 0x3af4a5e4.
//
// Solidity: function __RoleManager_init() returns()
func (_Rolemanager *RolemanagerTransactorSession) RoleManagerInit() (*types.Transaction, error) {
	return _Rolemanager.Contract.RoleManagerInit(&_Rolemanager.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Rolemanager *RolemanagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Rolemanager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Rolemanager *RolemanagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Rolemanager.Contract.GrantRole(&_Rolemanager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Rolemanager *RolemanagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Rolemanager.Contract.GrantRole(&_Rolemanager.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles_, address[] accounts_) returns()
func (_Rolemanager *RolemanagerTransactor) GrantRoleBatch(opts *bind.TransactOpts, roles_ [][32]byte, accounts_ []common.Address) (*types.Transaction, error) {
	return _Rolemanager.contract.Transact(opts, "grantRoleBatch", roles_, accounts_)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles_, address[] accounts_) returns()
func (_Rolemanager *RolemanagerSession) GrantRoleBatch(roles_ [][32]byte, accounts_ []common.Address) (*types.Transaction, error) {
	return _Rolemanager.Contract.GrantRoleBatch(&_Rolemanager.TransactOpts, roles_, accounts_)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles_, address[] accounts_) returns()
func (_Rolemanager *RolemanagerTransactorSession) GrantRoleBatch(roles_ [][32]byte, accounts_ []common.Address) (*types.Transaction, error) {
	return _Rolemanager.Contract.GrantRoleBatch(&_Rolemanager.TransactOpts, roles_, accounts_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Rolemanager *RolemanagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Rolemanager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Rolemanager *RolemanagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Rolemanager.Contract.RenounceRole(&_Rolemanager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Rolemanager *RolemanagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Rolemanager.Contract.RenounceRole(&_Rolemanager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role_, address account_) returns()
func (_Rolemanager *RolemanagerTransactor) RevokeRole(opts *bind.TransactOpts, role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Rolemanager.contract.Transact(opts, "revokeRole", role_, account_)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role_, address account_) returns()
func (_Rolemanager *RolemanagerSession) RevokeRole(role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Rolemanager.Contract.RevokeRole(&_Rolemanager.TransactOpts, role_, account_)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role_, address account_) returns()
func (_Rolemanager *RolemanagerTransactorSession) RevokeRole(role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Rolemanager.Contract.RevokeRole(&_Rolemanager.TransactOpts, role_, account_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address contractsRegistry_, bytes data_) returns()
func (_Rolemanager *RolemanagerTransactor) SetDependencies(opts *bind.TransactOpts, contractsRegistry_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Rolemanager.contract.Transact(opts, "setDependencies", contractsRegistry_, data_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address contractsRegistry_, bytes data_) returns()
func (_Rolemanager *RolemanagerSession) SetDependencies(contractsRegistry_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Rolemanager.Contract.SetDependencies(&_Rolemanager.TransactOpts, contractsRegistry_, data_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address contractsRegistry_, bytes data_) returns()
func (_Rolemanager *RolemanagerTransactorSession) SetDependencies(contractsRegistry_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Rolemanager.Contract.SetDependencies(&_Rolemanager.TransactOpts, contractsRegistry_, data_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Rolemanager *RolemanagerTransactor) SetInjector(opts *bind.TransactOpts, injector_ common.Address) (*types.Transaction, error) {
	return _Rolemanager.contract.Transact(opts, "setInjector", injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Rolemanager *RolemanagerSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Rolemanager.Contract.SetInjector(&_Rolemanager.TransactOpts, injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Rolemanager *RolemanagerTransactorSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Rolemanager.Contract.SetInjector(&_Rolemanager.TransactOpts, injector_)
}

// RolemanagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Rolemanager contract.
type RolemanagerRoleAdminChangedIterator struct {
	Event *RolemanagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RolemanagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RolemanagerRoleAdminChanged)
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
		it.Event = new(RolemanagerRoleAdminChanged)
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
func (it *RolemanagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RolemanagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RolemanagerRoleAdminChanged represents a RoleAdminChanged event raised by the Rolemanager contract.
type RolemanagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Rolemanager *RolemanagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RolemanagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Rolemanager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RolemanagerRoleAdminChangedIterator{contract: _Rolemanager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Rolemanager *RolemanagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RolemanagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Rolemanager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RolemanagerRoleAdminChanged)
				if err := _Rolemanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Rolemanager *RolemanagerFilterer) ParseRoleAdminChanged(log types.Log) (*RolemanagerRoleAdminChanged, error) {
	event := new(RolemanagerRoleAdminChanged)
	if err := _Rolemanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RolemanagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Rolemanager contract.
type RolemanagerRoleGrantedIterator struct {
	Event *RolemanagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *RolemanagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RolemanagerRoleGranted)
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
		it.Event = new(RolemanagerRoleGranted)
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
func (it *RolemanagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RolemanagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RolemanagerRoleGranted represents a RoleGranted event raised by the Rolemanager contract.
type RolemanagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Rolemanager *RolemanagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RolemanagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Rolemanager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RolemanagerRoleGrantedIterator{contract: _Rolemanager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Rolemanager *RolemanagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RolemanagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Rolemanager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RolemanagerRoleGranted)
				if err := _Rolemanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Rolemanager *RolemanagerFilterer) ParseRoleGranted(log types.Log) (*RolemanagerRoleGranted, error) {
	event := new(RolemanagerRoleGranted)
	if err := _Rolemanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RolemanagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Rolemanager contract.
type RolemanagerRoleRevokedIterator struct {
	Event *RolemanagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RolemanagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RolemanagerRoleRevoked)
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
		it.Event = new(RolemanagerRoleRevoked)
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
func (it *RolemanagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RolemanagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RolemanagerRoleRevoked represents a RoleRevoked event raised by the Rolemanager contract.
type RolemanagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Rolemanager *RolemanagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RolemanagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Rolemanager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RolemanagerRoleRevokedIterator{contract: _Rolemanager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Rolemanager *RolemanagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RolemanagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Rolemanager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RolemanagerRoleRevoked)
				if err := _Rolemanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Rolemanager *RolemanagerFilterer) ParseRoleRevoked(log types.Log) (*RolemanagerRoleRevoked, error) {
	event := new(RolemanagerRoleRevoked)
	if err := _Rolemanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
