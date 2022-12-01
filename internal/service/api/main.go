package api

import (
	"net"
	"net/http"
	"strconv"

	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/postgres"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/tokend/nft-books/book-svc/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const tokenIdIncrementKey = "token_id_increment"

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
	// Update increment key
	if err := s.setInitialSubscribeOffset(); err != nil {
		return errors.Wrap(err, "failed to set initial offset", logan.F{
			"initial_offset": cfg.DeploySignatureConfig().InitialOffset,
		})
	}

	r := s.router(cfg)

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func (s *service) setInitialSubscribeOffset() error {
	keyValueQ := postgres.NewKeyValueQ(s.db)

	return keyValueQ.Upsert(data.KeyValue{
		Key:   tokenIdIncrementKey,
		Value: strconv.FormatInt(s.cfg.DeploySignatureConfig().InitialOffset, 10),
	})
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
