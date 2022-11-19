package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/network-svc/connector/api"
)

type NetworkConfigurator interface {
	NetworkConnector() *api.Connector
}

type networkConfigurator struct {
	once   comfig.Once
	getter kv.Getter
}

type NetworkConnectorConfig struct {
	URL   string `fig:"url,required"`
	Token string `fig:"token,required"`
}

func NewNetworkConfigurator(getter kv.Getter) NetworkConfigurator {
	return &networkConfigurator{getter: getter}
}

func (c *networkConfigurator) NetworkConnector() *api.Connector {
	return c.once.Do(func() interface{} {
		config := NetworkConnectorConfig{}

		raw := kv.MustGetStringMap(c.getter, "networker")

		if err := figure.
			Out(&config).
			From(raw).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return api.NewConnector(config.Token, config.URL)
	}).(*api.Connector)
}
