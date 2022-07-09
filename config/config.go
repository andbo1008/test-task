package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DBDriver   string `enviroment:"DBDRIVER"`
	DBHost     string `enviroment:"DBHOST"`
	DBuser     string `enviroment:"DBUSER"`
	DBname     string `enviroment:"DBNAME"`
	DBpassword string `enviroment:"DBPASSWORD"`
	DBport     string `enviroment:"DBPORT"`
	DBsslmode  string `enviroment:"DBSSLMODE"`
}

func GetConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error load .env file\n", err)
	}

	var c Config
	if err := envconfig.Process("", &c); err != nil {
		logrus.Fatal(err)
	}
	return &c
}
