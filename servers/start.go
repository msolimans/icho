package servers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Start() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT)
		<-sigs
		cancel()
	}()

	//launch grpc
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("Starting gRPC service %s \n", grpcEndpoint)
		err := startGRPC(ctx)
		if err != nil {
			fmt.Println("gRPC server: ", err)
			cancel()
		}
	}()

	//launch rest
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("Starting REST service %s \n", httpEndpoint)
		err := startRest(ctx)
		if err != nil && err != http.ErrServerClosed {
			fmt.Print("REST service: ", err.Error())
			cancel()
		}
	}()

	wg.Wait()
}
