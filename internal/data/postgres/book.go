package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
)

const (
	booksTableName = "book"
	idColumn       = "id"
	priceColumn    = "price"
	deletedColumn  = "deleted"
)

func NewBooksQ(db *pgdb.DB) data.BookQ {
	return &BooksQ{
		db:       db.Clone(),
		selector: squirrel.Select("b.*").From(fmt.Sprintf("%s as b", booksTableName)),
	}
}

type BooksQ struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
}

func (b *BooksQ) New() data.BookQ {
	return NewBooksQ(b.db)
}

func (b *BooksQ) Insert(data data.Book) (int64, error) {
	clauses := structs.Map(data)
	var id int64

	stmt := squirrel.
		Insert(booksTableName).
		SetMap(clauses).
		Suffix("returning id")
	err := b.db.Get(&id, stmt)

	return id, err
}

func (b *BooksQ) Get() (*data.Book, error) {
	var result data.Book

	err := b.db.Get(&result, b.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (b *BooksQ) Select() ([]data.Book, error) {
	var result []data.Book

	err := b.db.Select(&result, b.selector)

	return result, err
}

func (b *BooksQ) FilterByID(id int64) data.BookQ {
	b.selector = b.selector.Where(squirrel.Eq{"b.id": id})
	return b
}

func (b *BooksQ) Page(params pgdb.OffsetPageParams) data.BookQ {
	b.selector = params.ApplyTo(b.selector, idColumn)
	return b
}

func (b *BooksQ) FilterActual() data.BookQ {
	b.selector = b.selector.Where(squirrel.Eq{"b.deleted": "f"})
	return b
}

func (b *BooksQ) DeleteByID(id int64) error {
	stmt := squirrel.
		Update(booksTableName).
		Set(deletedColumn, "t").
		Where(squirrel.Eq{"id": id})

	return b.db.Exec(stmt)
}

func (b *BooksQ) Update(data data.Book) error {
	clauses := structs.Map(data)
	stmt := squirrel.
		Update(booksTableName).
		SetMap(clauses).
		Where(squirrel.Eq{"id": data.ID})

	return b.db.Exec(stmt)
}

func (b *BooksQ) UpdatePriceByID(price string, id int64) error {
	return b.db.Exec(
		squirrel.
			Update(booksTableName).
			Set(priceColumn, price).
			Where(squirrel.Eq{"id": id}),
	)
}

func (b *BooksQ) UpdatePriceByAddress(price, address string) error {
	return b.db.Exec(
		squirrel.
			Update(booksTableName).
			Set(priceColumn, price).
			Where(squirrel.Eq{"address": address}),
	)
}
