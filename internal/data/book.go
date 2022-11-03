package data

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type BookQ interface {
	New() BookQ

	Get() (*Book, error)
	Select() ([]Book, error)

	Insert(data Book) (int64, error)
	Update(data Book) error
	DeleteByID(id int64) error

	UpdatePrice(price string, id int64) error
	UpdateContractName(name string, id int64) error
	UpdateLastBlock(newLastBlock uint64, id int64) error
	UpdateSymbol(newSymbol string, id int64) error

	// do not include deleted books
	FilterActual() BookQ
	FilterByID(id int64) BookQ

	Page(params pgdb.OffsetPageParams) BookQ
}

type Book struct {
	ID              int64  `db:"id" structs:"-"`
	Title           string `db:"title" structs:"title"`
	Symbol          string `db:"symbol" structs:"symbol"`
	Description     string `db:"description" structs:"description"`
	Price           string `db:"price" structs:"price"`
	ContractAddress string `db:"contract_address" structs:"contract_address"`
	ContractName    string `db:"contract_name" structs:"contract_name"`
	ContractVersion string `db:"contract_version" structs:"contract_version"`
	Banner          string `db:"banner" structs:"banner"`
	File            string `db:"file" structs:"file"`
	Deleted         bool   `db:"deleted" structs:"-"`
	LastBlock       uint64 `db:"last_block" structs:"last_block"`
}

func (b *Book) Address() common.Address {
	return common.HexToAddress(b.ContractAddress)
}
