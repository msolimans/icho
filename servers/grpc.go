package servers

import (
	"context"
	"fmt"
	"github.com/msolimans/icho/icho"
	"google.golang.org/grpc"
	"net"
)

func startGRPC(ctx context.Context) error {
	listen, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	icho.RegisterIchoServiceServer(grpcServer, icho.NewIchoServerImp())

	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down gRPC service")
		grpcServer.GracefulStop()
	}()
	return grpcServer.Serve(listen)

}
