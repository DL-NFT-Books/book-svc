package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/dl-nft-books/book-svc/resources"

	"github.com/Masterminds/squirrel"
	"github.com/dl-nft-books/book-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	booksTableName    = "book"
	idColumn          = "id"
	bannerColumn      = "banner"
	fileColumn        = "file"
	descriptionColumn = "description"
	createdAtColumn   = "created_at"

	booksNetworksTableName = "book_network"
	bookIdColumn           = "book_id"
	tokenIdColumn          = "token_id"
	contractAddressColumn  = "contract_address"
	deployStatusColumn     = "deploy_status"
	chainIdColumn          = "chain_id"
)

func NewBooksQ(db *pgdb.DB) data.BookQ {
	return &BooksQ{
		db: db.Clone(),
		selectBuilder: squirrel.Select(booksTableName+".*",
			fmt.Sprintf("json_agg(json_build_object('%s', %s, '%s', %s, '%s', %s, '%s', %s)) as network",
				tokenIdColumn, tokenIdColumn,
				contractAddressColumn, contractAddressColumn,
				deployStatusColumn, deployStatusColumn,
				chainIdColumn, chainIdColumn)).From(booksTableName).
			Join(fmt.Sprintf("%s on %s.%s = %s.%s", booksNetworksTableName,
				booksNetworksTableName, bookIdColumn,
				booksTableName, idColumn)).GroupBy(idColumn),
		updateBuilder:        squirrel.Update(booksTableName),
		networkUpdateBuilder: squirrel.Update(booksNetworksTableName),
	}
}

type BooksQ struct {
	db                   *pgdb.DB
	selectBuilder        squirrel.SelectBuilder
	updateBuilder        squirrel.UpdateBuilder
	networkUpdateBuilder squirrel.UpdateBuilder
}

func (b *BooksQ) New() data.BookQ {
	return NewBooksQ(b.db)
}

func (b *BooksQ) Insert(data data.Book) (id int64, err error) {
	statement := squirrel.Insert(booksTableName).
		Columns(bannerColumn, fileColumn, descriptionColumn, createdAtColumn).
		Values(data.Banner, data.File, data.Description, data.CreatedAt).Suffix("returning id")
	fmt.Println(statement.ToSql())
	err = b.db.Get(&id, statement)
	return
}

func (b *BooksQ) InsertNetwork(data ...data.BookNetwork) (err error) {
	statement := squirrel.Insert(booksNetworksTableName).
		Columns(bookIdColumn, tokenIdColumn, contractAddressColumn, chainIdColumn, deployStatusColumn)
	for _, network := range data {
		statement = statement.Values(network.BookId, network.TokenId, network.ContractAddress, network.ChainId, network.DeployStatus)
	}

	return b.db.Exec(statement)
}

func (b *BooksQ) Count() (uint64, error) {
	var res uint64
	selStmt := squirrel.Select("COUNT(book)").
		FromSelect(b.selectBuilder, "book")
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

func (b *BooksQ) FilterByChainId(chainId ...int64) data.BookQ {
	b.selectBuilder = b.selectBuilder.Where(squirrel.Eq{chainIdColumn: chainId})
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
	if updater.Description != nil {
		sql = sql.Set(descriptionColumn, *updater.Description)
	}
	return sql
}

func (b *BooksQ) UpdateDeployStatus(newStatus resources.DeployStatus, bookId, chainId int64) error {
	return b.db.Exec(b.networkUpdateBuilder.Set(deployStatusColumn, newStatus).Where(squirrel.Eq{bookIdColumn: bookId, chainIdColumn: chainId}))
}

func (b *BooksQ) UpdateContractAddress(newAddress string, id int64) error {
	return b.db.Exec(b.updateBuilder.Set(contractAddressColumn, newAddress).Where(squirrel.Eq{idColumn: id}))
}
