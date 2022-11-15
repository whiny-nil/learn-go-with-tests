package main_test

import (
	"fmt"
	"testing"

	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/adapters"
	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/adapters/grpcserver"
	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerFilePath = "./cmd/grpcserver/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, &driver)
}
