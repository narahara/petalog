package grpc_handler

import (
	"common/pkg/service/app"
)

type AppHandlerInterface interface {
	app.AppServiceServer
}

type AppHandler struct {
	app.UnimplementedAppServiceServer
}

func NewAppHandler() AppHandlerInterface {

	return &AppHandler{}
}
