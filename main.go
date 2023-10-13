package main

import (
	"log"

	"oec/internal"
	"oec/internal/config"
	"oec/internal/util"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := util.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	f := fiber.New()
	api := internal.NewAPI(cfg, f, db)

	log.Fatal(api.App.Listen(cfg.PORT))
}
