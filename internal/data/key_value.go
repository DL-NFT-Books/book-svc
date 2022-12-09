package data

type KeyValue struct {
	Key   string `db:"key" structs:"key"`
	Value string `db:"value" structs:"value"`
}

type KeyValueQ interface {
	New() KeyValueQ

	Get(key string) (*KeyValue, error)
	// Upsert updates value if there is one, insert if no
	Upsert(KeyValue) error
	// LockingGet reads row and locks the row for reading and updating
	// until the end of the current transaction
	LockingGet(key string) (*KeyValue, error)
}
