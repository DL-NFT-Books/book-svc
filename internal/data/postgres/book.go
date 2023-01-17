package postgres

import (
	"database/sql"
	"strings"

	"gitlab.com/tokend/nft-books/book-svc/resources"

	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
)

const (
	booksTableName           = "book"
	idColumn                 = "id"
	tokenIdColumn            = "token_id"
	priceColumn              = "price"
	deletedColumn            = "deleted"
	contractNameColumn       = "contract_name"
	contractAddressColumn    = "contract_address"
	deployStatusColumn       = "deploy_status"
	contractSymbolColumn     = "contract_symbol"
	bannerColumn             = "banner"
	fileColumn               = "file"
	titleColumn              = "title"
	lastBlockColumn          = "last_block"
	descriptionColumn        = "description"
	voucherTokenColumn       = "voucher_token"
	voucherTokenAmountColumn = "voucher_token_amount"
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

func (b *BooksQ) Insert(data data.Book) (id int64, err error) {
	statement := squirrel.Insert(booksTableName).SetMap(structs.Map(data)).Suffix("returning id")
	err = b.db.Get(&id, statement)
	return
}

func (b *BooksQ) Count(title *string) (uint64, error) {
	var res uint64
	selStmt := squirrel.Select("COUNT(id)").
		From(booksTableName)

	if title != nil {
		selStmt = selStmt.Where(squirrel.Like{`LOWER(title)`: "%" + strings.ToLower(*title) + "%"})
	}

	err := b.db.Get(&res, selStmt)

	return res, err
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

func (b *BooksQ) FilterByID(id ...int64) data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Eq{idColumn: id})
	return b
}

func (b *BooksQ) FilterByTitle(title string) data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Like{`LOWER(title)`: "%" + strings.ToLower(title) + "%"})
	return b
}

func (b *BooksQ) FilterByTokenId(tokenId ...int64) data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Eq{tokenIdColumn: tokenId})
	return b
}

func (b *BooksQ) FilterByDeployStatus(status ...resources.DeployStatus) data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Eq{deployStatusColumn: status})
	return b
}

func (b *BooksQ) FilterByContractAddress(address ...string) data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Eq{contractAddressColumn: address})
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

func (b *BooksQ) Update(updater data.BookUpdateParams, id int64) error {
	return b.db.Exec(
		b.applyUpdateParams(b.updateBuilder, updater).
			Where(squirrel.Eq{
				idColumn: id,
			}))
}

func (b *BooksQ) applyUpdateParams(sql squirrel.UpdateBuilder, updater data.BookUpdateParams) squirrel.UpdateBuilder {
	if updater.File != nil {
		sql = sql.Set(fileColumn, *updater.File)
	}
	if updater.Banner != nil {
		sql = sql.Set(bannerColumn, *updater.Banner)
	}
	if updater.Title != nil {
		sql = sql.Set(titleColumn, *updater.Title)
	}
	if updater.Description != nil {
		sql = sql.Set(descriptionColumn, *updater.Description)
	}
	if updater.Contract != nil {
		sql = sql.Set(contractAddressColumn, *updater.Contract)
	}
	if updater.ContractName != nil {
		sql = sql.Set(contractNameColumn, *updater.ContractName)
	}
	if updater.DeployStatus != nil {
		sql = sql.Set(deployStatusColumn, *updater.DeployStatus)
	}
	if updater.Symbol != nil {
		sql = sql.Set(contractSymbolColumn, *updater.Symbol)
	}
	if updater.Price != nil {
		sql = sql.Set(priceColumn, *updater.Price)
	}
	if updater.VoucherToken != nil {
		sql = sql.Set(voucherTokenColumn, *updater.VoucherToken)
	}
	if updater.VoucherTokenAmount != nil {
		sql = sql.Set(voucherTokenAmountColumn, *updater.VoucherTokenAmount)
	}

	return sql
}

func (b *BooksQ) UpdatePrice(price string, id int64) error {
	return b.db.Exec(b.updateBuilder.Set(priceColumn, price).Where(squirrel.Eq{idColumn: id}))
}

func (b *BooksQ) UpdateContractName(name string, id int64) error {
	return b.db.Exec(b.updateBuilder.Set(contractNameColumn, name).Where(squirrel.Eq{idColumn: id}))
}

func (b *BooksQ) UpdateDeployStatus(newStatus resources.DeployStatus, id int64) error {
	return b.db.Exec(b.updateBuilder.Set(deployStatusColumn, newStatus).Where(squirrel.Eq{idColumn: id}))
}

func (b *BooksQ) UpdateContractAddress(newAddress string, id int64) error {
	return b.db.Exec(b.updateBuilder.Set(contractAddressColumn, newAddress).Where(squirrel.Eq{idColumn: id}))
}

func (b *BooksQ) UpdateLastBlock(newLastBlock uint64, id int64) error {
	return b.db.Exec(b.updateBuilder.Set(lastBlockColumn, newLastBlock).Where(squirrel.Eq{idColumn: id}))
}

func (b *BooksQ) UpdateSymbol(newSymbol string, id int64) error {
	return b.db.Exec(b.updateBuilder.Set(contractSymbolColumn, newSymbol).Where(squirrel.Eq{idColumn: id}))
}

func (b *BooksQ) UpdateContractParams(name, symbol, price string, id int64) error {
	return b.db.Exec(b.updateBuilder.
		Set(contractNameColumn, name).
		Set(contractSymbolColumn, symbol).
		Set(priceColumn, price).
		Where(squirrel.Eq{
			idColumn: id,
		}),
	)
}
