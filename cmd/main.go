package main

import (
	"log"

	"github.com/SenyashaGo/jwt-auth/pkg/config"
	"github.com/SenyashaGo/jwt-auth/pkg/srv"
	"github.com/SenyashaGo/jwt-auth/pkg/storage"
)

func main() {
	cfg, err := config.Parse("../config/config.json")
	if err != nil {
		panic("Can`t parse a config file.")
	}

	newStorage, err := storage.NewStorage(cfg.Dsn)
	if err != nil {
		log.Println(err)
	}

	srv.Run(newStorage, cfg)
}
