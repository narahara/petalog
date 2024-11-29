package server

import "app/pkg/infrastructure/config"

type Server interface {
}

type server struct {
	config *config.Config
}

func NewServer(config *config.Config) Server {
	return &server{
		config: config,
	}
}
