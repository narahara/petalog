package server

import (
	"app/pkg/infrastructure/config"
	console_handler "console/pkg/interface/console"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"common/pkg/service/console"
	utilNet "common/pkg/util/net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	myMiddleware "common/pkg/interface/grpc/middleware"
)

type Server interface {
	Run() (err error)
}

type server struct {
	config  *config.Config
	handler console_handler.ConsoleHandlerInterface
}

func NewServer(config *config.Config, handler console_handler.ConsoleHandlerInterface) Server {
	return &server{
		config:  config,
		handler: handler,
	}
}

func (s *server) Run() (err error) {
	ec := echo.New()
	ec.HideBanner = true

	appHome := os.Getenv("APP_HOME")
	uiHome := fmt.Sprintf("%s/server/console/ui", appHome)

	ec.Static("/js", fmt.Sprintf("%s/dist/js", uiHome))
	ec.Static("/css", fmt.Sprintf("%s/dist/css", uiHome))
	ec.Static("/img", fmt.Sprintf("%s/dist/img", uiHome))
	ec.Static("/fonts", fmt.Sprintf("%s/dist/fonts", uiHome))
	ec.GET("/*", func(c echo.Context) error {
		return c.File(fmt.Sprintf("%s/dist/index.html", uiHome))
	})

	ipAddr, err := utilNet.GetLocalAddr()
	if err != nil {
		log.Println(err)
	}

	mux, err := initGrpcGateway()
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf(":%d", 50000)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*20), grpc.MaxCallSendMsgSize(1024*1024*20))}
	console.RegisterConsoleServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts)

	go func() {
		grpcServer := initGrpc()
		console.RegisterConsoleServiceServer(grpcServer, s.handler)
		fmt.Printf("grpc server is running on %s:%d\n", ipAddr, 50000)
		grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", 50000))
		if err != nil {
			log.Fatalln(err)
		}
		err = grpcServer.Serve(grpcListener)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	ec.Any("/api/*", func(c echo.Context) error {
		w := c.Response().Writer
		r := c.Request()
		r.URL.Path = strings.Replace(r.URL.Path, "/api", "", 1)
		fmt.Printf("grpc-gw:%s (%d)\n", r.URL.Path, len(r.URL.Path))
		mux.ServeHTTP(w, r)
		return nil
	})

	ec.Start(fmt.Sprintf(":%d", s.config.Port))

	return nil
}

func initGrpcGateway() (*runtime.ServeMux, error) {
	runtime.DefaultContextTimeout = time.Second * 120
	mux := runtime.NewServeMux(
		//runtime.WithErrorHandler(interceptor.GrpcErrorHandler),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   false,
				Multiline:       true,
				Indent:          "  ",
				EmitUnpopulated: true, // デフォルト値を出力する。これがないとomitemptyされる
			},
		}),
	)
	return mux, nil
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
