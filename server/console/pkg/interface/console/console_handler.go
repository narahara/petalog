package console_handler

import "common/pkg/service/console"

type ConsoleHandlerInterface interface {
	console.ConsoleServiceServer
}

type ConsoleHandler struct {
	console.UnimplementedConsoleServiceServer
}

func NewConsoleHandler() ConsoleHandlerInterface {
	return &ConsoleHandler{}
}
