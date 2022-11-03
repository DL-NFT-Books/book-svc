package config

import "time"

type Runner struct {
	NormalPeriod      time.Duration `fig:"normal_period"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period"`
	MaxAbnormalPeriod time.Duration `fig:"normal_period"`
}

var defaultRunner = Runner{
	NormalPeriod:      time.Minute,
	MinAbnormalPeriod: time.Minute,
	MaxAbnormalPeriod: time.Minute,
}
