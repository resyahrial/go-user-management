package config

import "fmt"

type DatabaseConfig struct {
	DbName          string `yaml:"name" env:"DB_NAME"`
	Host            string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port            string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Username        string `yaml:"username" env:"DB_USERNAME"`
	Password        string `yaml:"password" env:"DB_PASSWORD"`
	MaxIddleConn    int    `yaml:"maxidleconn" env:"DB_MAX_IDDLE_CONN"`
	MaxOpenConn     int    `yaml:"maxopenconn" env:"DB_MAX_OPEN_CONN"`
	ConnMaxLifetime int    `yaml:"connmaxlifetime" env:"DB_CONN_MAX_LIFETIME"`
}

func (config DatabaseConfig) GetDatabaseConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", config.Host, config.Username, config.Password, config.DbName, config.Port)
}
