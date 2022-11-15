package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/adapters"
	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/adapters/httpserver"
	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8081"
		dockerFilePath = "./cmd/httpserver/Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}
