package main

import (
	"app/pkg/infrastructure/config"
	"console/pkg/infrastructure/server"
	console_handler "console/pkg/interface/console"
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

	handler := console_handler.NewConsoleHandler()
	server := server.NewServer(config, handler)

	err = server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
