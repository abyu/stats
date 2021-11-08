package internal

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"os"
)

//Config for the app
type Config struct {
	DBHost string `yaml:"dbHost" env:"DB_HOST" env-default:"localhost"`
	DBPort string `yaml:"dbPort" env:"DB_PORT" env-default:"5432"`
	DBUserName string `yaml:"dbUsername" env:"DB_USER"`
	DBPassword string `yaml:"dbPassword" env:"DB_PASSWORD"`
	DB string `yaml:"db" env:"DB" env-default:"postgres"`
}

//AppConfig return config from config.yaml or env vars
func AppConfig() (*Config, error)  {
	var cfg Config
	reader := getConfigReader()

	err := reader(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getConfigReader() func(*Config) error {
	configFile := "config.yaml"
	_, err := os.Stat(configFile)
	if err != nil {
		log.Warnf("Couldn't find a config.yaml file, defaulting to env variables")
		return func(c *Config) error { return cleanenv.ReadEnv(c) }
	}
	return func(c *Config) error { return cleanenv.ReadConfig(configFile, c) }
}