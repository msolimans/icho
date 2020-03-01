package main

import (
	"context"
	"fmt"
	"github.com/msolimans/icho/servers"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
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
		fmt.Println("Starting gRPC service")
		err := servers.StartgRPC(ctx)
		if err != nil {
			fmt.Println("gRPC server: ", err)
			cancel()
		}
	}()

	//launch rest
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Starting REST service")
		err := servers.StartRest(ctx)
		if err != nil && err != http.ErrServerClosed {
			fmt.Print("REST service: ", err.Error())
			cancel()
		}
	}()

	wg.Wait()
}
