package main

import (
	"app/pkg/infrastructure/config"
	"app/pkg/infrastructure/server"
	app_handler "app/pkg/interface/app"
	"flag"
	"log"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "config file path")
	flag.Parse()

	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	handler := app_handler.NewAppHandler()
	server := server.NewServer(config, handler)

	err = server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
