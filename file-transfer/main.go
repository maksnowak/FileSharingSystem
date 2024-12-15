package main

import (
	"context"
	"file-transfer/app"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"time"
)

func main() {
	port := flag.String("port", "8080", "Port to listen on.")
	flag.Parse()

	a := app.App{}
	a.Initialize()

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := a.Close(ctx); err != nil {
			panic(err)
		}
		close(done)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.Run(&ctx, ":"+*port)
	<-done
}
