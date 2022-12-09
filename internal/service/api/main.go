package api

import (
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	documenter "gitlab.com/tokend/nft-books/blob-svc/connector/api"
	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	keyValue "gitlab.com/tokend/nft-books/book-svc/internal/data/key_value"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/postgres"
	doorman "gitlab.com/tokend/nft-books/doorman/connector"
	"net"
	"net/http"
	"strconv"
)

type service struct {
	// Base configs
	cfg      config.Config
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	db       *pgdb.DB

	// Custom configs
	mimeTypes          *config.MimeTypes
	deploySignatureCfg *config.DeploySignatureConfig

	// Connectors
	doorman    doorman.ConnectorI
	documenter *documenter.Connector
}

func (s *service) run(cfg config.Config) error {
	s.log.Info("Service started")

	// Update increment key
	if err := s.setInitialSubscribeOffset(); err != nil {
		return errors.Wrap(err, "failed to set initial offset", logan.F{
			"initial_offset": cfg.DeploySignatureConfig().InitialOffset,
		})
	}

	r := s.router()
	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

// setInitialSubscribeOffset is a function that setups the initial parameter of token id
// in the KV table that is needed for a book deployment flow
func (s *service) setInitialSubscribeOffset() error {
	keyValueQ := postgres.NewKeyValueQ(s.db)

	return keyValueQ.Upsert(data.KeyValue{
		Key:   keyValue.TokenIdIncrementKey,
		Value: strconv.FormatInt(s.cfg.DeploySignatureConfig().InitialOffset, 10),
	})
}

func newService(cfg config.Config) *service {
	return &service{
		cfg:                cfg,
		log:                cfg.Log(),
		copus:              cfg.Copus(),
		listener:           cfg.Listener(),
		db:                 cfg.DB(),
		mimeTypes:          cfg.MimeTypes(),
		deploySignatureCfg: cfg.DeploySignatureConfig(),

		doorman:    cfg.DoormanConnector(),
		documenter: cfg.DocumenterConnector(),
	}
}

func Run(cfg config.Config) error {
	if err := newService(cfg).run(cfg); err != nil {
		return errors.Wrap(err, "failed to initialize a service")
	}

	return nil
}
