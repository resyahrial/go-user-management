package config

type AuthConfig struct {
	AccessTimeDuration int    `yaml:"accesstimeduration" env:"ACCESS_TIME_DURATION"`
	AccessSecretKey    string `yaml:"accesssecretkey" env:"ACCESS_SECRET_KEY"`
}
