package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	s3Cfg "gitlab.com/tokend/nft-books/blob-svc/connector/config"
	doormanCfg "gitlab.com/tokend/nft-books/doorman/connector/config"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	MimeTypesConfigurator
	doormanCfg.DoormanConfiger
	s3Cfg.Documenter

	UpdateTracker() UpdateTracker
	DeployTracker() DeployTracker
	EtherClient() EtherClient
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	MimeTypesConfigurator
	doormanCfg.DoormanConfiger
	s3Cfg.Documenter

	getter            kv.Getter
	updateTrackerOnce comfig.Once
	deployTrackerOnce comfig.Once
	ethererOnce       comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:                getter,
		Databaser:             pgdb.NewDatabaser(getter),
		Copuser:               copus.NewCopuser(getter),
		Listenerer:            comfig.NewListenerer(getter),
		Logger:                comfig.NewLogger(getter, comfig.LoggerOpts{}),
		MimeTypesConfigurator: NewMimeTypesConfigurator(getter),
		DoormanConfiger:       doormanCfg.NewDoormanConfiger(getter),
		Documenter:            s3Cfg.NewDocumenter(getter),
	}
}
