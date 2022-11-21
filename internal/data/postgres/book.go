package postgres

import (
	"database/sql"

	"gitlab.com/tokend/nft-books/book-svc/resources"

	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
)

const (
	booksTableName        = "book"
	idColumn              = "id"
	tokenIdColumn         = "token_id"
	priceColumn           = "price"
	deletedColumn         = "deleted"
	contractNameColumn    = "contract_name"
	contractAddressColumn = "contract_address"
	deployStatusColumn    = "deploy_status"
	contractSymbolColumn  = "contract_symbol"
	bannerColumn          = "banner"
	fileColumn            = "file"
	titleColumn           = "title"
	lastBlockColumn       = "last_block"
	descriptionColumn     = "description"
	chainIDColumn         = "chain_id"
)

func NewBooksQ(db *pgdb.DB) data.BookQ {
	return &BooksQ{
		db:            db.Clone(),
		selectBuilder: squirrel.Select("*").From(booksTableName),
		updateBuilder: squirrel.Update(booksTableName),
	}
}

type BooksQ struct {
	db            *pgdb.DB
	selectBuilder squirrel.SelectBuilder
	updateBuilder squirrel.UpdateBuilder
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

	err := b.db.Get(&result, b.selectBuilder)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (b *BooksQ) Select() ([]data.Book, error) {
	var result []data.Book

	err := b.db.Select(&result, b.selectBuilder)

	return result, err
}

func (b *BooksQ) FilterByID(id int64) data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Eq{
		idColumn: id,
	})

	return b
}

func (b *BooksQ) FilterByTokenId(tokenId int64) data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Eq{
		tokenIdColumn: tokenId,
	})

	return b
}

func (b *BooksQ) FilterByDeployStatus(status resources.DeployStatus) data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Eq{
		deployStatusColumn: status,
	})

	return b
}

func (b *BooksQ) FilterByChainID(chainID int64) data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Eq{
		chainIDColumn: chainID,
	})

	return b
}

func (b *BooksQ) Page(params pgdb.OffsetPageParams) data.BookQ {
	b.selectBuilder = params.ApplyTo(b.selectBuilder, idColumn)

	return b
}

func (b *BooksQ) FilterActual() data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Eq{
		deletedColumn: "f",
	})

	return b
}

func (b *BooksQ) DeleteByID(id int64) error {
	stmt := b.updateBuilder.
		Set(deletedColumn, "t").
		Where(squirrel.Eq{
			idColumn: id,
		})

	return b.db.Exec(stmt)
}

func (b *BooksQ) UpdateFile(file string, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(fileColumn, file).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}

func (b *BooksQ) UpdateBanner(banner string, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(bannerColumn, banner).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}

func (b *BooksQ) UpdateTitle(title string, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(titleColumn, title).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}

func (b *BooksQ) UpdateDescription(desc string, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(descriptionColumn, desc).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}

func (b *BooksQ) UpdatePrice(price string, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(priceColumn, price).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}

func (b *BooksQ) UpdateContractName(name string, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(contractNameColumn, name).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}

func (b *BooksQ) UpdateDeployStatus(newStatus resources.DeployStatus, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(deployStatusColumn, newStatus).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}

func (b *BooksQ) UpdateContractAddress(newAddress string, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(contractAddressColumn, newAddress).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}

func (b *BooksQ) UpdateLastBlock(newLastBlock uint64, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(lastBlockColumn, newLastBlock).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}

func (b *BooksQ) UpdateSymbol(newSymbol string, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(contractSymbolColumn, newSymbol).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}

func (b *BooksQ) UpdateContractParams(name, symbol, price string, id int64) error {
	return b.db.Exec(
		b.updateBuilder.
			Set(contractNameColumn, name).
			Set(contractSymbolColumn, symbol).
			Set(priceColumn, price).
			Where(squirrel.Eq{
				idColumn: id,
			}),
	)
}
