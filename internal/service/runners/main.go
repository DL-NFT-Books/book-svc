package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/book-svc/internal/config"
)

type Runner func(ctx context.Context)

func initializeRunners(cfg config.Config) (runners []Runner) {
	runners = append(runners, NewUpdateTracker(cfg).Run)
	runners = append(runners, NewDeployTracker(cfg).Run)

	return
}

func Run(cfg config.Config, ctx context.Context) {
	for _, runner := range initializeRunners(cfg) {
		go runner(ctx)
	}
}
