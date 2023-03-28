// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package irolemanager

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

// IrolemanagerMetaData contains all meta data concerning the Irolemanager contract.
var IrolemanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ADMINISTRATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKETPLACE_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_SUPERVISOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_FACTORY_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_REGISTRY_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__RoleManager_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles_\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts_\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"isMarketplaceManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supervisor\",\"type\":\"address\"}],\"name\":\"isRoleSupervisor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"isTokenFactoryManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"isTokenManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"isTokenRegistryManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"isWithdrawalManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IrolemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IrolemanagerMetaData.ABI instead.
var IrolemanagerABI = IrolemanagerMetaData.ABI

// Irolemanager is an auto generated Go binding around an Ethereum contract.
type Irolemanager struct {
	IrolemanagerCaller     // Read-only binding to the contract
	IrolemanagerTransactor // Write-only binding to the contract
	IrolemanagerFilterer   // Log filterer for contract events
}

// IrolemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IrolemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IrolemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IrolemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IrolemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IrolemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IrolemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IrolemanagerSession struct {
	Contract     *Irolemanager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IrolemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IrolemanagerCallerSession struct {
	Contract *IrolemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IrolemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IrolemanagerTransactorSession struct {
	Contract     *IrolemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IrolemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IrolemanagerRaw struct {
	Contract *Irolemanager // Generic contract binding to access the raw methods on
}

// IrolemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IrolemanagerCallerRaw struct {
	Contract *IrolemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// IrolemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IrolemanagerTransactorRaw struct {
	Contract *IrolemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIrolemanager creates a new instance of Irolemanager, bound to a specific deployed contract.
func NewIrolemanager(address common.Address, backend bind.ContractBackend) (*Irolemanager, error) {
	contract, err := bindIrolemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Irolemanager{IrolemanagerCaller: IrolemanagerCaller{contract: contract}, IrolemanagerTransactor: IrolemanagerTransactor{contract: contract}, IrolemanagerFilterer: IrolemanagerFilterer{contract: contract}}, nil
}

// NewIrolemanagerCaller creates a new read-only instance of Irolemanager, bound to a specific deployed contract.
func NewIrolemanagerCaller(address common.Address, caller bind.ContractCaller) (*IrolemanagerCaller, error) {
	contract, err := bindIrolemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IrolemanagerCaller{contract: contract}, nil
}

// NewIrolemanagerTransactor creates a new write-only instance of Irolemanager, bound to a specific deployed contract.
func NewIrolemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IrolemanagerTransactor, error) {
	contract, err := bindIrolemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IrolemanagerTransactor{contract: contract}, nil
}

// NewIrolemanagerFilterer creates a new log filterer instance of Irolemanager, bound to a specific deployed contract.
func NewIrolemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IrolemanagerFilterer, error) {
	contract, err := bindIrolemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IrolemanagerFilterer{contract: contract}, nil
}

// bindIrolemanager binds a generic wrapper to an already deployed contract.
func bindIrolemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IrolemanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Irolemanager *IrolemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Irolemanager.Contract.IrolemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Irolemanager *IrolemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Irolemanager.Contract.IrolemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Irolemanager *IrolemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Irolemanager.Contract.IrolemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Irolemanager *IrolemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Irolemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Irolemanager *IrolemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Irolemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Irolemanager *IrolemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Irolemanager.Contract.contract.Transact(opts, method, params...)
}

// ADMINISTRATORROLE is a free data retrieval call binding the contract method 0xf45edb5f.
//
// Solidity: function ADMINISTRATOR_ROLE() view returns(bytes32)
func (_Irolemanager *IrolemanagerCaller) ADMINISTRATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "ADMINISTRATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINISTRATORROLE is a free data retrieval call binding the contract method 0xf45edb5f.
//
// Solidity: function ADMINISTRATOR_ROLE() view returns(bytes32)
func (_Irolemanager *IrolemanagerSession) ADMINISTRATORROLE() ([32]byte, error) {
	return _Irolemanager.Contract.ADMINISTRATORROLE(&_Irolemanager.CallOpts)
}

// ADMINISTRATORROLE is a free data retrieval call binding the contract method 0xf45edb5f.
//
// Solidity: function ADMINISTRATOR_ROLE() view returns(bytes32)
func (_Irolemanager *IrolemanagerCallerSession) ADMINISTRATORROLE() ([32]byte, error) {
	return _Irolemanager.Contract.ADMINISTRATORROLE(&_Irolemanager.CallOpts)
}

// MARKETPLACEMANAGER is a free data retrieval call binding the contract method 0xce2c2940.
//
// Solidity: function MARKETPLACE_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerCaller) MARKETPLACEMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "MARKETPLACE_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MARKETPLACEMANAGER is a free data retrieval call binding the contract method 0xce2c2940.
//
// Solidity: function MARKETPLACE_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerSession) MARKETPLACEMANAGER() ([32]byte, error) {
	return _Irolemanager.Contract.MARKETPLACEMANAGER(&_Irolemanager.CallOpts)
}

// MARKETPLACEMANAGER is a free data retrieval call binding the contract method 0xce2c2940.
//
// Solidity: function MARKETPLACE_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerCallerSession) MARKETPLACEMANAGER() ([32]byte, error) {
	return _Irolemanager.Contract.MARKETPLACEMANAGER(&_Irolemanager.CallOpts)
}

// ROLESUPERVISOR is a free data retrieval call binding the contract method 0x0d80af9b.
//
// Solidity: function ROLE_SUPERVISOR() view returns(bytes32)
func (_Irolemanager *IrolemanagerCaller) ROLESUPERVISOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "ROLE_SUPERVISOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLESUPERVISOR is a free data retrieval call binding the contract method 0x0d80af9b.
//
// Solidity: function ROLE_SUPERVISOR() view returns(bytes32)
func (_Irolemanager *IrolemanagerSession) ROLESUPERVISOR() ([32]byte, error) {
	return _Irolemanager.Contract.ROLESUPERVISOR(&_Irolemanager.CallOpts)
}

// ROLESUPERVISOR is a free data retrieval call binding the contract method 0x0d80af9b.
//
// Solidity: function ROLE_SUPERVISOR() view returns(bytes32)
func (_Irolemanager *IrolemanagerCallerSession) ROLESUPERVISOR() ([32]byte, error) {
	return _Irolemanager.Contract.ROLESUPERVISOR(&_Irolemanager.CallOpts)
}

// TOKENFACTORYMANAGER is a free data retrieval call binding the contract method 0xfb70f02f.
//
// Solidity: function TOKEN_FACTORY_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerCaller) TOKENFACTORYMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "TOKEN_FACTORY_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENFACTORYMANAGER is a free data retrieval call binding the contract method 0xfb70f02f.
//
// Solidity: function TOKEN_FACTORY_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerSession) TOKENFACTORYMANAGER() ([32]byte, error) {
	return _Irolemanager.Contract.TOKENFACTORYMANAGER(&_Irolemanager.CallOpts)
}

// TOKENFACTORYMANAGER is a free data retrieval call binding the contract method 0xfb70f02f.
//
// Solidity: function TOKEN_FACTORY_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerCallerSession) TOKENFACTORYMANAGER() ([32]byte, error) {
	return _Irolemanager.Contract.TOKENFACTORYMANAGER(&_Irolemanager.CallOpts)
}

// TOKENMANAGER is a free data retrieval call binding the contract method 0xe0956e0f.
//
// Solidity: function TOKEN_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerCaller) TOKENMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "TOKEN_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENMANAGER is a free data retrieval call binding the contract method 0xe0956e0f.
//
// Solidity: function TOKEN_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerSession) TOKENMANAGER() ([32]byte, error) {
	return _Irolemanager.Contract.TOKENMANAGER(&_Irolemanager.CallOpts)
}

// TOKENMANAGER is a free data retrieval call binding the contract method 0xe0956e0f.
//
// Solidity: function TOKEN_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerCallerSession) TOKENMANAGER() ([32]byte, error) {
	return _Irolemanager.Contract.TOKENMANAGER(&_Irolemanager.CallOpts)
}

// TOKENREGISTRYMANAGER is a free data retrieval call binding the contract method 0x6e3673bb.
//
// Solidity: function TOKEN_REGISTRY_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerCaller) TOKENREGISTRYMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "TOKEN_REGISTRY_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENREGISTRYMANAGER is a free data retrieval call binding the contract method 0x6e3673bb.
//
// Solidity: function TOKEN_REGISTRY_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerSession) TOKENREGISTRYMANAGER() ([32]byte, error) {
	return _Irolemanager.Contract.TOKENREGISTRYMANAGER(&_Irolemanager.CallOpts)
}

// TOKENREGISTRYMANAGER is a free data retrieval call binding the contract method 0x6e3673bb.
//
// Solidity: function TOKEN_REGISTRY_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerCallerSession) TOKENREGISTRYMANAGER() ([32]byte, error) {
	return _Irolemanager.Contract.TOKENREGISTRYMANAGER(&_Irolemanager.CallOpts)
}

// WITHDRAWALMANAGER is a free data retrieval call binding the contract method 0xd2f03194.
//
// Solidity: function WITHDRAWAL_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerCaller) WITHDRAWALMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "WITHDRAWAL_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWALMANAGER is a free data retrieval call binding the contract method 0xd2f03194.
//
// Solidity: function WITHDRAWAL_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerSession) WITHDRAWALMANAGER() ([32]byte, error) {
	return _Irolemanager.Contract.WITHDRAWALMANAGER(&_Irolemanager.CallOpts)
}

// WITHDRAWALMANAGER is a free data retrieval call binding the contract method 0xd2f03194.
//
// Solidity: function WITHDRAWAL_MANAGER() view returns(bytes32)
func (_Irolemanager *IrolemanagerCallerSession) WITHDRAWALMANAGER() ([32]byte, error) {
	return _Irolemanager.Contract.WITHDRAWALMANAGER(&_Irolemanager.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin) view returns(bool)
func (_Irolemanager *IrolemanagerCaller) IsAdmin(opts *bind.CallOpts, _admin common.Address) (bool, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "isAdmin", _admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin) view returns(bool)
func (_Irolemanager *IrolemanagerSession) IsAdmin(_admin common.Address) (bool, error) {
	return _Irolemanager.Contract.IsAdmin(&_Irolemanager.CallOpts, _admin)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin) view returns(bool)
func (_Irolemanager *IrolemanagerCallerSession) IsAdmin(_admin common.Address) (bool, error) {
	return _Irolemanager.Contract.IsAdmin(&_Irolemanager.CallOpts, _admin)
}

// IsMarketplaceManager is a free data retrieval call binding the contract method 0x2019e9cb.
//
// Solidity: function isMarketplaceManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerCaller) IsMarketplaceManager(opts *bind.CallOpts, _manager common.Address) (bool, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "isMarketplaceManager", _manager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketplaceManager is a free data retrieval call binding the contract method 0x2019e9cb.
//
// Solidity: function isMarketplaceManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerSession) IsMarketplaceManager(_manager common.Address) (bool, error) {
	return _Irolemanager.Contract.IsMarketplaceManager(&_Irolemanager.CallOpts, _manager)
}

// IsMarketplaceManager is a free data retrieval call binding the contract method 0x2019e9cb.
//
// Solidity: function isMarketplaceManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerCallerSession) IsMarketplaceManager(_manager common.Address) (bool, error) {
	return _Irolemanager.Contract.IsMarketplaceManager(&_Irolemanager.CallOpts, _manager)
}

// IsRoleSupervisor is a free data retrieval call binding the contract method 0xe3e941b3.
//
// Solidity: function isRoleSupervisor(address _supervisor) view returns(bool)
func (_Irolemanager *IrolemanagerCaller) IsRoleSupervisor(opts *bind.CallOpts, _supervisor common.Address) (bool, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "isRoleSupervisor", _supervisor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRoleSupervisor is a free data retrieval call binding the contract method 0xe3e941b3.
//
// Solidity: function isRoleSupervisor(address _supervisor) view returns(bool)
func (_Irolemanager *IrolemanagerSession) IsRoleSupervisor(_supervisor common.Address) (bool, error) {
	return _Irolemanager.Contract.IsRoleSupervisor(&_Irolemanager.CallOpts, _supervisor)
}

// IsRoleSupervisor is a free data retrieval call binding the contract method 0xe3e941b3.
//
// Solidity: function isRoleSupervisor(address _supervisor) view returns(bool)
func (_Irolemanager *IrolemanagerCallerSession) IsRoleSupervisor(_supervisor common.Address) (bool, error) {
	return _Irolemanager.Contract.IsRoleSupervisor(&_Irolemanager.CallOpts, _supervisor)
}

// IsTokenFactoryManager is a free data retrieval call binding the contract method 0xde94b229.
//
// Solidity: function isTokenFactoryManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerCaller) IsTokenFactoryManager(opts *bind.CallOpts, _manager common.Address) (bool, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "isTokenFactoryManager", _manager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenFactoryManager is a free data retrieval call binding the contract method 0xde94b229.
//
// Solidity: function isTokenFactoryManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerSession) IsTokenFactoryManager(_manager common.Address) (bool, error) {
	return _Irolemanager.Contract.IsTokenFactoryManager(&_Irolemanager.CallOpts, _manager)
}

// IsTokenFactoryManager is a free data retrieval call binding the contract method 0xde94b229.
//
// Solidity: function isTokenFactoryManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerCallerSession) IsTokenFactoryManager(_manager common.Address) (bool, error) {
	return _Irolemanager.Contract.IsTokenFactoryManager(&_Irolemanager.CallOpts, _manager)
}

// IsTokenManager is a free data retrieval call binding the contract method 0xebc26119.
//
// Solidity: function isTokenManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerCaller) IsTokenManager(opts *bind.CallOpts, _manager common.Address) (bool, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "isTokenManager", _manager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenManager is a free data retrieval call binding the contract method 0xebc26119.
//
// Solidity: function isTokenManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerSession) IsTokenManager(_manager common.Address) (bool, error) {
	return _Irolemanager.Contract.IsTokenManager(&_Irolemanager.CallOpts, _manager)
}

// IsTokenManager is a free data retrieval call binding the contract method 0xebc26119.
//
// Solidity: function isTokenManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerCallerSession) IsTokenManager(_manager common.Address) (bool, error) {
	return _Irolemanager.Contract.IsTokenManager(&_Irolemanager.CallOpts, _manager)
}

// IsTokenRegistryManager is a free data retrieval call binding the contract method 0xccc8843e.
//
// Solidity: function isTokenRegistryManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerCaller) IsTokenRegistryManager(opts *bind.CallOpts, _manager common.Address) (bool, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "isTokenRegistryManager", _manager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenRegistryManager is a free data retrieval call binding the contract method 0xccc8843e.
//
// Solidity: function isTokenRegistryManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerSession) IsTokenRegistryManager(_manager common.Address) (bool, error) {
	return _Irolemanager.Contract.IsTokenRegistryManager(&_Irolemanager.CallOpts, _manager)
}

// IsTokenRegistryManager is a free data retrieval call binding the contract method 0xccc8843e.
//
// Solidity: function isTokenRegistryManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerCallerSession) IsTokenRegistryManager(_manager common.Address) (bool, error) {
	return _Irolemanager.Contract.IsTokenRegistryManager(&_Irolemanager.CallOpts, _manager)
}

// IsWithdrawalManager is a free data retrieval call binding the contract method 0x5cea8645.
//
// Solidity: function isWithdrawalManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerCaller) IsWithdrawalManager(opts *bind.CallOpts, _manager common.Address) (bool, error) {
	var out []interface{}
	err := _Irolemanager.contract.Call(opts, &out, "isWithdrawalManager", _manager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalManager is a free data retrieval call binding the contract method 0x5cea8645.
//
// Solidity: function isWithdrawalManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerSession) IsWithdrawalManager(_manager common.Address) (bool, error) {
	return _Irolemanager.Contract.IsWithdrawalManager(&_Irolemanager.CallOpts, _manager)
}

// IsWithdrawalManager is a free data retrieval call binding the contract method 0x5cea8645.
//
// Solidity: function isWithdrawalManager(address _manager) view returns(bool)
func (_Irolemanager *IrolemanagerCallerSession) IsWithdrawalManager(_manager common.Address) (bool, error) {
	return _Irolemanager.Contract.IsWithdrawalManager(&_Irolemanager.CallOpts, _manager)
}

// RoleManagerInit is a paid mutator transaction binding the contract method 0x3af4a5e4.
//
// Solidity: function __RoleManager_init() returns()
func (_Irolemanager *IrolemanagerTransactor) RoleManagerInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Irolemanager.contract.Transact(opts, "__RoleManager_init")
}

// RoleManagerInit is a paid mutator transaction binding the contract method 0x3af4a5e4.
//
// Solidity: function __RoleManager_init() returns()
func (_Irolemanager *IrolemanagerSession) RoleManagerInit() (*types.Transaction, error) {
	return _Irolemanager.Contract.RoleManagerInit(&_Irolemanager.TransactOpts)
}

// RoleManagerInit is a paid mutator transaction binding the contract method 0x3af4a5e4.
//
// Solidity: function __RoleManager_init() returns()
func (_Irolemanager *IrolemanagerTransactorSession) RoleManagerInit() (*types.Transaction, error) {
	return _Irolemanager.Contract.RoleManagerInit(&_Irolemanager.TransactOpts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles_, address[] accounts_) returns()
func (_Irolemanager *IrolemanagerTransactor) GrantRoleBatch(opts *bind.TransactOpts, roles_ [][32]byte, accounts_ []common.Address) (*types.Transaction, error) {
	return _Irolemanager.contract.Transact(opts, "grantRoleBatch", roles_, accounts_)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles_, address[] accounts_) returns()
func (_Irolemanager *IrolemanagerSession) GrantRoleBatch(roles_ [][32]byte, accounts_ []common.Address) (*types.Transaction, error) {
	return _Irolemanager.Contract.GrantRoleBatch(&_Irolemanager.TransactOpts, roles_, accounts_)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles_, address[] accounts_) returns()
func (_Irolemanager *IrolemanagerTransactorSession) GrantRoleBatch(roles_ [][32]byte, accounts_ []common.Address) (*types.Transaction, error) {
	return _Irolemanager.Contract.GrantRoleBatch(&_Irolemanager.TransactOpts, roles_, accounts_)
}
