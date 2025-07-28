package main

import (
	"log"
	"scopegg2/internal/config"
	"scopegg2/internal/di"
)

func main() {
	cfg := config.InitConfig()

	app, err := di.InitializeApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
