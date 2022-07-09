package main

import (
	"rest-api/config"
	"rest-api/handler"
	"rest-api/repository"
	"rest-api/router"
	"rest-api/service"

	"github.com/sirupsen/logrus"
)

func main() {
	c := config.GetConfig()
	db, err := repository.PostgresConnect(c)
	if err != nil {
		logrus.Fatal(err)
	}
	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	router := router.NewRouter(h)
	router.RouterStart()
}
