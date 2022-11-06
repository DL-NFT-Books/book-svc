package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
)

const deployTrackerYamlKey = "deploy_tracker"

type DeployTracker struct {
	Name          string         `fig:"name"`
	Address       common.Address `fig:"address"`
	Runner        Runner         `fig:"runner"`
	FirstBlock    uint64         `fig:"first_block"`
	IterationSize uint64         `fig:"iteration_size"`
}

var defaultDeployTracker = DeployTracker{
	Name:          "deploy_tracker",
	Address:       common.Address{},
	Runner:        defaultRunner,
	FirstBlock:    0,
	IterationSize: 100,
}

func (c *config) DeployTracker() DeployTracker {
	return c.deployTrackerOnce.Do(func() interface{} {
		cfg := defaultDeployTracker

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks, contractHook).
			From(kv.MustGetStringMap(c.getter, deployTrackerYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mint tracker config"))
		}

		return cfg
	}).(DeployTracker)
}

var contractHook = figure.Hooks{
	"common.Address": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			address := common.HexToAddress(v)
			return reflect.ValueOf(address), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},
}
