package config

type Seed struct {
	Admin struct {
		Email    string `yaml:"email" env:"SEED_ADMIN_EMAIL"`
		Password string `yaml:"password" env:"SEED_ADMIN_PASSWORD"`
	} `yaml:"admin"`
}
