package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type JWTConfigurator interface {
	JWT() *JWT
}

type JWT struct {
	SignatureKey string `fig:"signature_key,required"`
}

type jwtConfigurator struct {
	getter kv.Getter
	once   comfig.Once
}

func NewJWTConfigurator(getter kv.Getter) JWTConfigurator {
	return &jwtConfigurator{
		getter: getter,
	}
}

func (c *jwtConfigurator) JWT() *JWT {
	return c.once.Do(func() interface{} {
		config := JWT{}

		if err := figure.Out(&config).From(kv.MustGetStringMap(c.getter, "jwt")).Please(); err != nil {
			panic(err)
		}

		return &config
	}).(*JWT)
}
