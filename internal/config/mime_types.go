package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const mimeTypesYamlKey = "mime_types"

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

		if err := figure.Out(&config).
			From(kv.MustGetStringMap(c.getter, mimeTypesYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mime types config"))
		}

		return &config
	}).(*MimeTypes)
}
