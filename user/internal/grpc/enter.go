package grpc

import "google.golang.org/grpc"

type Service interface {
	Register(server *grpc.Server)
}
