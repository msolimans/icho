package servers

import (
	"context"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"net"
)

func StartgRPC(ctx context.Context) error {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	icho.RegisterIchoServiceServer(grpcServer, icho.NewIchoServerImp())

	go func() {
		<-ctx.Done()
		log.Info("Shutting down gRPC service")
		grpcServer.GracefulStop()
	}()
	return grpcServer.Serve(listen)

}
