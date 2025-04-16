package config

type SystemConfig struct {
	Env Env `mapstructure:"env" yaml:"env" json:"env"`
}

type Env string

const (
	EnvLocal Env = "local"
	EnvDev   Env = "dev"
	EnvProd  Env = "prod"
)
