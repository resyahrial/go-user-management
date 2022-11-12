package config

type HasherConfig struct {
	Cost string `yaml:"cost" env:"HASHER_COST"`
}
