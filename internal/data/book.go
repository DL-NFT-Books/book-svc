package data

import "gitlab.com/distributed_lab/kit/pgdb"

type BookQ interface {
	New() BookQ
	Insert(data Book) (int64, error)
	Get() (*Book, error)
	Select() ([]Book, error)
	FilterByID(id int64) BookQ
	Page(params pgdb.OffsetPageParams) BookQ
}

type Book struct {
	ID          int64  `db:"id" structs:"-"`
	Title       string `db:"id" structs:"title"`
	Description string `db:"description" structs:"description"`
	Price       int32  `db:"price" structs:"price"`
	Banner      string `db:"banner" structs:"banner"`
	File        string `db:"file" structs:"file"`
}
