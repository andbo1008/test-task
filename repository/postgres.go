package repository

import (
	"fmt"
	"rest-api/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func PostgresConnect(c *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(c.DBDriver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.DBHost, c.DBport, c.DBuser, c.DBname, c.DBpassword, c.DBsslmode))
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	return db, nil
}
