package config

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type AdminsConfig struct {
	Admins []common.Address
}

func (c *config) AdminsConfig() AdminsConfig {
	return c.admins.Do(func() interface{} {
		var conf = struct {
			List []string `fig:"list,required"`
		}{}

		err := figure.
			Out(&conf).
			From(kv.MustGetStringMap(c.getter, "admins")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		admins := make([]common.Address, 0, len(conf.List))
		for _, l := range conf.List {
			admins = append(admins, common.HexToAddress(l))
		}

		return AdminsConfig{
			Admins: admins,
		}
	}).(AdminsConfig)
}
