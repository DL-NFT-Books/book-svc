package data

import "gitlab.com/distributed_lab/kit/pgdb"

type BookQ interface {
	New() BookQ
	Insert(data Book) (int64, error)
	Get() (*Book, error)
	Select() ([]Book, error)
	Update(data Book) error
	DeleteByID(id int64) error
	FilterByID(id int64) BookQ
	// do not include deleted books
	FilterActual() BookQ
	Page(params pgdb.OffsetPageParams) BookQ
}

type Book struct {
	ID              int64  `db:"id" structs:"-"`
	Title           string `db:"title" structs:"title"`
	Description     string `db:"description" structs:"description"`
	Price           string `db:"price" structs:"price"`
	ContractAddress string `db:"contract_address" structs:"contract_address"`
	ContractName    string `db:"contract_name" structs:"contract_name"`
	ContractVersion string `db:"contract_version" structs:"contract_version"`
	Banner          string `db:"banner" structs:"banner"`
	File            string `db:"file" structs:"file"`
	Deleted         bool   `db:"deleted" structs:"-"`
}
