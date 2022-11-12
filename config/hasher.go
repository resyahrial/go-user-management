package config

type HasherConfig struct {
	Cost int `yaml:"cost" env:"HASHER_COST"`
}
