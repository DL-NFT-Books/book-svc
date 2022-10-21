package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type MimeTypesConfigurator interface {
	MimeTypes() *MimeTypes
}

type MimeTypes struct {
	AllowedBannerMimeTypes []string `fig:"banner,required"`
	AllowedFileMimeTypes   []string `fig:"file,required"`
}

type mimeTypesConfigurator struct {
	getter kv.Getter
	once   comfig.Once
}

func NewMimeTypesConfigurator(getter kv.Getter) MimeTypesConfigurator {
	return &mimeTypesConfigurator{
		getter: getter,
	}
}

func (c *mimeTypesConfigurator) MimeTypes() *MimeTypes {
	return c.once.Do(func() interface{} {
		config := MimeTypes{}

		if err := figure.Out(&config).From(kv.MustGetStringMap(c.getter, "mime_types")).Please(); err != nil {
			panic(err)
		}

		return &config
	}).(*MimeTypes)
}
