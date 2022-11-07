package api

import (
	"context"
	"net"
	"net/http"

	"gitlab.com/tokend/nft-books/book-svc/internal/service/runners"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/tokend/nft-books/book-svc/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	cfg                 config.Config
	log                 *logan.Entry
	copus               types.Copus
	listener            net.Listener
	db                  *pgdb.DB
	mimeTypes           *config.MimeTypes
	deploySignatureConf *config.DeploySignatureConfig
}

func (s *service) run(cfg config.Config) error {
	s.log.Info("Service started")
	r := s.router(cfg)

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	ctx := context.Background()
	runners.Run(s.cfg, ctx)

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		cfg:                 cfg,
		log:                 cfg.Log(),
		copus:               cfg.Copus(),
		listener:            cfg.Listener(),
		db:                  cfg.DB(),
		mimeTypes:           cfg.MimeTypes(),
		deploySignatureConf: cfg.DeploySignatureConfig(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(cfg); err != nil {
		panic(errors.Wrap(err, "failed to initialize a service"))
	}
}
