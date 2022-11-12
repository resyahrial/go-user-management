package config

type AppConfig struct {
	Name          string `yaml:"name"`
	Version       string `yaml:"version"`
	Environment   string `yaml:"-"`
	LogLevel      int    `yaml:"log_level"`
	DebugMode     bool   `yaml:"debug"`
	ServerAppHost string `yaml:"host"`
	ServerAppPort string `yaml:"port"`
}
