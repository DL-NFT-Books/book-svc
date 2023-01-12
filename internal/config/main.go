package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	documenter "gitlab.com/tokend/nft-books/blob-svc/connector/config"
	doormaner "gitlab.com/tokend/nft-books/doorman/connector/config"
	networker "gitlab.com/tokend/nft-books/network-svc/connector"
)

type Config interface {
	// Base configs
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer

	// Connectors
	doormaner.DoormanConfiger
	documenter.Documenter
	networker.NetworkConfigurator

	// Custom configs
	MimeTypesConfigurator
	DeploySignatureConfigurator
}

type config struct {
	// Base configs
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer

	// Connectors
	doormaner.DoormanConfiger
	documenter.Documenter
	networker.NetworkConfigurator

	// Custom configs
	MimeTypesConfigurator
	DeploySignatureConfigurator

	getter kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		// Base configs
		Databaser:  pgdb.NewDatabaser(getter),
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),

		// Custom configs
		MimeTypesConfigurator:       NewMimeTypesConfigurator(getter),
		DeploySignatureConfigurator: NewDeploySignatureConfigurator(getter),

		// Connectors
		Documenter:          documenter.NewDocumenter(getter),
		DoormanConfiger:     doormaner.NewDoormanConfiger(getter),
		NetworkConfigurator: networker.NewNetworkConfigurator(getter),
		getter:              getter,
	}
}
