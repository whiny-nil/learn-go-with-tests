package main

import (
	"context"
	"log"
	"net"

	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/adapters/grpcserver"
	go_specs_greet "github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/domain/interactions"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &GreetServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type GreetServer struct {
	grpcserver.UnimplementedGreeterServer
}

func (g GreetServer) Greet(ctx context.Context, request *grpcserver.GreetRequest) (*grpcserver.GreetReply, error) {
	return &grpcserver.GreetReply{Message: go_specs_greet.Greet(request.Name)}, nil
}
