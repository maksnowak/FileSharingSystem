package main

import (
	_ "accounts/docs"
	"context"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//	@title			Accounts API
//	@version		0.1
//	@description	Webserver serving a complete implementation of account managing API endpoint.

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// @BasePath	/
func main() {
	// parse flags
	fmt.Println("Hello from `accounts` microservice!")
	port := flag.String("port", "2024", "Port to listen on.")
	addr := flag.String("address", "localhost", "Address of the API endpoint")
	flag.Parse()

	logger := log.New(os.Stdout, "Server: ", log.Flags())

	// define api endpoints
	rout := chi.NewRouter()
	{
		rout.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://%v:%v/swagger/doc.json", *addr, *port)),
		))
		rout.Get("/hello", Hello)
	}

	// start server
	serv := &http.Server{Addr: *addr + ":" + *port, Handler: rout}
	go func() {
		logger.Printf("http: Listening on %v:%v\n", *addr, *port)
		if err := serv.ListenAndServe(); err != nil {
			logger.Fatalln(err)
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
