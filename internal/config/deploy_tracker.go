package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const deployTrackerYamlKey = "deploy_tracker"

type DeployTracker struct {
	Name          string `fig:"name"`
	Runner        Runner `fig:"runner"`
	IterationSize uint64 `fig:"iteration_size"`
}

var defaultDeployTracker = DeployTracker{
	Name:          "deploy_tracker",
	Runner:        defaultRunner,
	IterationSize: 100,
}

func (c *config) DeployTracker() DeployTracker {
	return c.deployTrackerOnce.Do(func() interface{} {
		cfg := defaultDeployTracker

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, deployTrackerYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mint tracker config"))
		}

		return cfg
	}).(DeployTracker)
}
