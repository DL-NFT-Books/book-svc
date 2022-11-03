package data

//go:generate mockery --case=underscore --name=DB
type DB interface {
	New() DB

	KeyValue() KeyValueQ
	Books() BookQ

	Transaction(func() error) error
}
