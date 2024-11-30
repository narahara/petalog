package my_middleware

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryFunc(p interface{}) error {
	return status.Errorf(codes.Internal, "Unexpected error")
}
