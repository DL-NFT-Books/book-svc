package postgres

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
)

type db struct {
	raw *pgdb.DB
}

func NewDB(rawDB *pgdb.DB) data.DB {
	return &db{
		raw: rawDB,
	}
}

func (db *db) New() data.DB {
	return NewDB(db.raw.Clone())
}

func (db *db) KeyValue() data.KeyValueQ {
	return NewKeyValueQ(db.raw)
}

func (db *db) Books() data.BookQ {
	return NewBooksQ(db.raw)
}

func (db *db) Transaction(fn func() error) error {
	return db.raw.Transaction(func() error {
		return fn()
	})
}
