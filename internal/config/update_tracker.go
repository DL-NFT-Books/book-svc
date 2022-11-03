package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

const updateTrackerYamlKey = "update_tracker"

type UpdateTracker struct {
	Name          string `fig:"name"`
	Capacity      int64  `fig:"capacity"`
	IterationSize uint64 `fig:"iteration_size"`
	Runner        Runner `fig:"runner"`
}

var defaultMintTracker = UpdateTracker{
	Name:          "update_tracker",
	Capacity:      1,
	IterationSize: 1000,
	Runner:        defaultRunner,
}

func (c *config) UpdateTracker() UpdateTracker {
	return c.updateTrackerOnce.Do(func() interface{} {
		cfg := defaultMintTracker

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, updateTrackerYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out update tracker config"))
		}

		return cfg
	}).(UpdateTracker)
}
