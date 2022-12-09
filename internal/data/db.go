package data

type DB interface {
	New() DB

	KeyValue() KeyValueQ
	Books() BookQ

	Transaction(func() error) error
}
