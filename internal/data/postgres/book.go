package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
)

const booksTableName = "book"

func NewBooksQ(db *pgdb.DB) data.BookQ {
	return &BooksQ{
		db:  db.Clone(),
		sql: squirrel.Select("b.*").From(fmt.Sprintf("%s as b", booksTableName)),
	}
}

type BooksQ struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func (b *BooksQ) New() data.BookQ {
	return NewBooksQ(b.db)
}

func (b *BooksQ) Insert(data data.Book) (int64, error) {
	clauses := structs.Map(data)
	var id int64

	stmt := squirrel.Insert(booksTableName).SetMap(clauses).Suffix("returning id")
	err := b.db.Get(&id, stmt)

	return id, err
}

func (b *BooksQ) Get() (*data.Book, error) {
	var result data.Book

	err := b.db.Get(&result, b.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (b *BooksQ) Select() ([]data.Book, error) {
	var result []data.Book

	err := b.db.Select(&result, b.sql)

	return result, err
}

func (b *BooksQ) FilterByID(id int64) data.BookQ {
	b.sql = b.sql.Where(squirrel.Eq{"b.id": id})
	return b
}

func (b *BooksQ) Page(params pgdb.OffsetPageParams) data.BookQ {
	b.sql = params.ApplyTo(b.sql, "id")
	return b
}
