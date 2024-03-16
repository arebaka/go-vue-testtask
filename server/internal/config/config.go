package config

import (
	"errors"

	"are.moe/testtask/internal/domain"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type DBInfo struct {
	Host     string `env:"PGHOST,required"`
	Port     int    `env:"PGPORT" envDefault:"5432"`
	User     string `env:"PGUSER,required"`
	Password string `env:"PGPASSWORD,required"`
	DBname   string `env:"PGDATABASE,required"`
}

type Config struct {
	Port     uint16                 `env:"PORT" envDefault:"3000"`
	PageSize domain.ProductPageSize `env:"PAGE_SIZE" envDefault:"20"`
	DBInfo   DBInfo
}

var config *Config

func Init() error {
	godotenv.Load(".env")
	config = &Config{}
	if err := env.Parse(config); err != nil {
		return errors.New("Failed to configure: " + err.Error())
	}
	return nil
}

func Get() *Config {
	return config
}
