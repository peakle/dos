package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
)

var url *string
var port *int
var method *string

func main() {
	var workersCount = flag.Int("workers", 1, "workers count")
	url = flag.String("url", "localhost", "url for dos")
	port = flag.Int("method", 80, "port for dos")
	method = flag.String("method", "GET", "method for dos")

	flag.Parse()

	var ctx, cancel = context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 0; i < *workersCount; i++ {
		wg.Add(1)
		worker(ctx, &wg)
	}

	var sigs = make(chan os.Signal)
	signal.Notify(sigs, os.Kill, os.Interrupt)

	go func() {
		select {
		case <-sigs:
			cancel()
		}

		return
	}()

	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			dos()
		}
	}
}

func dos() {

}
