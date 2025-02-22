package main

import (
	"context"
	"file-transfer/app"
	"flag"
	"os"
	"os/signal"
	"syscall"

	_ "file-transfer/docs"
	"time"
)

//	@title			File transfer API
//	@version		0.2
//	@description	Webserver providing saving and retrieval of files from MongoDB

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// @BasePath	/
func main() {
	port := flag.String("port", "8080", "Port to listen on.")
	flag.Parse()

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	a := app.App{}

	go func() {
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := a.Close(ctx); err != nil {
			panic(err)
		}
		close(done)
	}()

	a.Initialize(&ctx)
	a.Run(&ctx, ":"+*port)
	<-done
}
