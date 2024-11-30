package my_middleware

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func DebugInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		path := info.FullMethod
		fmt.Println("debug interceptor:", path)
		return handler(ctx, req)
	}
}
