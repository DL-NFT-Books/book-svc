package data

import (
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

type BookQ interface {
	New() BookQ

	Get() (*Book, error)
	Select() ([]Book, error)
	Count() (uint64, error)

	Insert(data Book) (int64, error)
	InsertNetwork(data ...BookNetwork) (err error)
	Update(updater BookUpdateParams, id int64) error

	FilterByID(id ...int64) BookQ
	FilterByContractAddress(address ...string) BookQ
	FilterByChainId(chainId ...int64) BookQ

	Page(params pgdb.OffsetPageParams) BookQ
}

// BookUpdateParams is a structure for applicable update params on bookQ `Update`
type BookUpdateParams struct {
	Banner      *string
	File        *string
	Description *string
}

type Book struct {
	ID              int64     `db:"id" structs:"-"`
	Description     string    `db:"description" structs:"description"`
	CreatedAt       time.Time `db:"created_at" structs:"created_at"`
	Banner          string    `db:"banner" structs:"banner"`
	File            string    `db:"file" structs:"file"`
	NetworkAsString string    `db:"network" structs:"network"`
}
type BookNetwork struct {
	BookId          int64  `db:"book_id" structs:"book_id"`
	ContractAddress string `db:"contract_address" structs:"contract_address"`
	ChainId         int64  `db:"chain_id" structs:"chain_id"`
}
