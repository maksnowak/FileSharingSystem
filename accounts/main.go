package main

import (
	"accounts/db"
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
	port := flag.String("port", "2024", "Port to listen on.")
	host := flag.String("address", "localhost", "Address of the API endpoint")
	flag.Parse()

	logger := log.New(os.Stdout, "server: ", log.Flags())

	r := chi.NewRouter()
	{
		if os.Getenv("APP_ENV") != "prod" {
			addr := fmt.Sprintf("http://localhost:%v/swagger/", *port)
			r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(addr+"doc.json")))
			logger.Printf("Swagger UI available at %v\n", addr+"index.html")
		}
		//r.Get("/hello", Hello)
	}

	serv := &http.Server{Addr: *host + ":" + *port, Handler: r}
	go func() {
		logger.Printf("http: Listening on %v:%v\n", *host, *port)
		if err := serv.ListenAndServe(); err != nil {
			logger.Fatalln(err)
		}
	}()

	// DB USAGE
	db.Connect()
	db.GetCollection("users")

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db.Disconnect()
	if err := serv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
