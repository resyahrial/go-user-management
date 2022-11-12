package config

type AppConfig struct {
	Name          string `yaml:"name" env:"APP_NAME"`
	DebugMode     bool   `yaml:"debug" env:"APP_DEBUG_MODE"`
	ServerAppHost string `yaml:"host" env:"APP_HOST" env-default:"localhost"`
	ServerAppPort string `yaml:"port" env:"APP_PORT" env-default:"8080"`
}
