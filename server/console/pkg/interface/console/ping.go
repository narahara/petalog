package console_handler

import (
	"common/pkg/service/console"
	"context"
)

func (a *ConsoleHandler) Ping(ctx context.Context, req *console.PingRequest) (*console.PingResponse, error) {
	res := &console.PingResponse{
		Message: req.Message,
	}
	return res, nil
}
