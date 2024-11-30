package server

import (
	"app/pkg/infrastructure/config"
	app_handler "app/pkg/interface/app"
	utilNet "common/pkg/util/net"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"

	myMiddleware "common/pkg/interface/grpc/middleware"
	"common/pkg/service/app"
)

type Server interface {
	Run() (err error)
}

type server struct {
	config  *config.Config
	handler app_handler.AppHandlerInterface
}

func NewServer(config *config.Config, handler app_handler.AppHandlerInterface) Server {
	return &server{
		config:  config,
		handler: handler,
	}
}

func (s *server) Run() (err error) {
	grpcServer := initGrpc()

	app.RegisterAppServiceServer(grpcServer, s.handler)

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return err
	}

	ipAddr, err := utilNet.GetLocalAddr()
	if err != nil {
		return err
	}

	fmt.Printf("server is running on %s:%d\n", ipAddr, s.config.Port)
	err = grpcServer.Serve(grpcListener)
	if err != nil {
		return err
	}
	return nil
}

// GRPCの初期化
func initGrpc() *grpc.Server {
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(myMiddleware.RecoveryFunc),
	}
	mid := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_recovery.UnaryServerInterceptor(opts...),
		myMiddleware.DebugInterceptor(),
	))
	s := grpc.NewServer(mid, grpc.MaxRecvMsgSize(1024*1024*20), grpc.MaxSendMsgSize(1024*1024*20))
	return s
}
