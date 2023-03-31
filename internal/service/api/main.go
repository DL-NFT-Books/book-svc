package api

import (
	documenter "github.com/dl-nft-books/blob-svc/connector/api"
	"github.com/dl-nft-books/book-svc/internal/config"
	doorman "github.com/dl-nft-books/doorman/connector"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net"
	"net/http"
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

func (s *service) run() error {
	s.log.Info("Service started")

	r := s.router()
	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
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
	if err := newService(cfg).run(); err != nil {
		return errors.Wrap(err, "failed to initialize a service")
	}

	return nil
}
