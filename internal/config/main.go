package config

import (
	documenter "github.com/dl-nft-books/blob-svc/connector/config"
	doormaner "github.com/dl-nft-books/doorman/connector/config"
	networker "github.com/dl-nft-books/network-svc/connector"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
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
		MimeTypesConfigurator: NewMimeTypesConfigurator(getter),

		// Connectors
		Documenter:          documenter.NewDocumenter(getter),
		DoormanConfiger:     doormaner.NewDoormanConfiger(getter),
		NetworkConfigurator: networker.NewNetworkConfigurator(getter),
		getter:              getter,
	}
}
