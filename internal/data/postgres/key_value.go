package postgres

import (
	"database/sql"

	"github.com/dl-nft-books/book-svc/internal/data"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	TokenIdIncrementKey = "token_id_increment"
)

const (
	keyValueTable = "key_value"
	keyColumn     = "key"
	valueColumn   = "value"
)

var (
	keyValueSelect = sq.Select("*").From(keyValueTable)
)

type keyValueQ struct {
	db *pgdb.DB
}

func NewKeyValueQ(db *pgdb.DB) data.KeyValueQ {
	return &keyValueQ{
		db: db,
	}
}

func (q *keyValueQ) Upsert(kv data.KeyValue) error {
	query := sq.Insert(keyValueTable).
		SetMap(structs.Map(kv)).
		Suffix("ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value")

	return q.db.Exec(query)
}

func (q *keyValueQ) New() data.KeyValueQ {
	return NewKeyValueQ(q.db.Clone())
}

func (q *keyValueQ) Get(key string) (*data.KeyValue, error) {
	return q.get(key, false)
}

func (q *keyValueQ) LockingGet(key string) (*data.KeyValue, error) {
	return q.get(key, true)
}
func (q *keyValueQ) get(key string, forUpdate bool) (*data.KeyValue, error) {
	stmt := keyValueSelect.Where(sq.Eq{keyColumn: key})
	if forUpdate {
		stmt = stmt.Suffix("FOR UPDATE")
	}
	var value data.KeyValue

	err := q.db.Get(&value, stmt)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &value, err
}
