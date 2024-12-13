package main

import (
	"context"
	"file-transfer/app"
	"flag"
	"os"
	"os/signal"
	"time"
)

func main() {
	port := flag.String("port", "8080", "Port to listen on.")
	host := flag.String("address", "localhost", "Address of the API endpoint")
	flag.Parse()

	a := app.App{}
	a.Initialize()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.Run(&ctx, *host+":"+*port)
	if err := a.Close(ctx); err != nil {
		panic(err)
	}
}
