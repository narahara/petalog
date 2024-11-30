package app_handler

import (
	"common/pkg/service/app"
	"context"
)

func (a *AppHandler) Ping(ctx context.Context, req *app.PingRequest) (*app.PingResponse, error) {
	res := &app.PingResponse{
		Message: req.Message,
		Name:    req.Name,
	}
	return res, nil
}
