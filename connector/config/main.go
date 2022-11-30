package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/connector/api"
)

type BooksConnectorConfigurator interface {
	BooksConnector() *api.Connector
}

type booksConnectorConfigurator struct {
	once   comfig.Once
	getter kv.Getter
}

type BooksConnectorConfig struct {
	URL   string `fig:"url,required"`
	Token string `fig:"token,required"`
}

func NewBooksConfigurator(getter kv.Getter) BooksConnectorConfigurator {
	return &booksConnectorConfigurator{getter: getter}
}

func (c *booksConnectorConfigurator) BooksConnector() *api.Connector {
	return c.once.Do(func() interface{} {
		config := BooksConnectorConfig{}

		raw := kv.MustGetStringMap(c.getter, "booker")

		if err := figure.
			Out(&config).
			From(raw).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return api.NewConnector(config.Token, config.URL)
	}).(*api.Connector)
}
