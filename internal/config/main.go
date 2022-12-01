package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	documenter "gitlab.com/tokend/nft-books/blob-svc/connector/config"
	doormaner "gitlab.com/tokend/nft-books/doorman/connector/config"
)

type Config interface {
	// base
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer

	// connectors
	doormaner.DoormanConfiger
	documenter.Documenter

	// additional configs
	MimeTypesConfigurator
	DeploySignatureConfigurator
}

type config struct {
	// base
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer

	// connectors
	doormaner.DoormanConfiger
	documenter.Documenter

	// additional configs
	MimeTypesConfigurator
	DeploySignatureConfigurator

	getter kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:                      getter,
		Databaser:                   pgdb.NewDatabaser(getter),
		Copuser:                     copus.NewCopuser(getter),
		Listenerer:                  comfig.NewListenerer(getter),
		Logger:                      comfig.NewLogger(getter, comfig.LoggerOpts{}),
		MimeTypesConfigurator:       NewMimeTypesConfigurator(getter),
		DoormanConfiger:             doormaner.NewDoormanConfiger(getter),
		Documenter:                  documenter.NewDocumenter(getter),
		DeploySignatureConfigurator: NewDeploySignatureConfigurator(getter),
	}
}
