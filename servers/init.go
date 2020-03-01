package servers

import (
	"flag"
	"fmt"
	"strings"
)

var httpEndpoint string
var grpcEndpoint string

func init() {
	flag.StringVar(&httpEndpoint, "http", ":8888", "Http Endpoint of Icho")
	flag.StringVar(&grpcEndpoint, "grpc", ":9999", "gRPC Endpoint of Icho")
	flag.Parse()

	if !strings.Contains(httpEndpoint, ":") {
		httpEndpoint = fmt.Sprintf(":%s", httpEndpoint)
	}

	if !strings.Contains(grpcEndpoint, ":") {
		grpcEndpoint = fmt.Sprintf(":%s", grpcEndpoint)
	}

}
