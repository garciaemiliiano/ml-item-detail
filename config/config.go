package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment string

var instance *Config

const (
	Local   = Environment("local")
	Develop = Environment("develop")
)

type Config struct {
	ENV         Environment
	PORT        string
	SQLITE_PATH string
	MS_NAME     string
}

func newConfig() (*Config, error) {
	_ = godotenv.Load(".env.local")
	env := Environment(GetEnv("ENV", "develop"))

	conf := &Config{
		ENV:         env,
		PORT:        GetEnv("PORT", "5000"),
		SQLITE_PATH: GetEnv("SQLITE_PATH", ""),
		MS_NAME:     GetEnv("MS_NAME", "ml-item-detail"),
	}

	return conf, nil
}

func Get() (*Config, error) {
	if instance == nil {
		config, err := newConfig()
		if err != nil {
			return nil, err
		}
		instance = config
	}
	return instance, nil
}

func GetConfig() Config {
	config, err := Get()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	return *config
}

func GetEnv(key string, fallback ...string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return ""
}
