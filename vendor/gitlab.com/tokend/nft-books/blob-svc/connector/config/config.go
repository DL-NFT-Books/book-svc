package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/blob-svc/connector/api"
)

type Documenter interface {
	DocumenterConnector() *api.Connector
}

type documenter struct {
	once   comfig.Once
	getter kv.Getter
}

func NewDocumenter(getter kv.Getter) Documenter {
	return &documenter{
		getter: getter,
	}
}

type documenterConfig struct {
	URL   string `fig:"url,required"`
	Token string `fig:"token,required"`
}

func (c *documenter) DocumenterConnector() *api.Connector {
	return c.once.Do(func() interface{} {
		var conf documenterConfig

		if err := figure.
			Out(&conf).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "connector")).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out documenter"))
		}

		return api.NewConnector(conf.URL, conf.Token)
	}).(*api.Connector)
}
