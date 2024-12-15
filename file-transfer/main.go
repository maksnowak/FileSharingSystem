package main

import (
	"context"
	"file-transfer/app"
	"flag"
	// "os"
	// "os/signal"
	"time"
)

func main() {
	port := flag.String("port", "8080", "Port to listen on.")
	flag.Parse()

	a := app.App{}
	a.Initialize()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.Run(&ctx, ":"+*port)
	if err := a.Close(ctx); err != nil {
		panic(err)
	}
}
