package main

import (
	"context"
	_ "file-encryption/docs"
	"file-encryption/handler"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//	@title			File Encryption API
//	@version		0.1
//	@description	Webserver serving file encryption/decryption capabilities.

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// @BasePath
func main() {
	port := flag.String("port", "7780", "Port to listen on.")
	host := flag.String("address", "0.0.0.0", "Address of the API endpoint")
	flag.Parse()

	r := chi.NewRouter()
	{
		r.Use(middleware.Heartbeat("/ping"))
		r.Use(middleware.Recoverer)
		r.Use(middleware.Logger)
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"http://localhost*", "https://localhost*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		r.Post("/encrypt", handler.Encrypt)
		r.Post("/decrypt", handler.Decrypt)

		if os.Getenv("APP_ENV") != "prod" {
			addr := fmt.Sprintf("http://localhost:%v/swagger/", *port)
			r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(addr+"doc.json")))
			log.Printf("Swagger UI available at %v\n", addr+"index.html")
		}
	}

	serv := &http.Server{Addr: *host + ":" + *port, Handler: r}
	go func() {
		log.Printf("http: Listening on %v:%v\n", *host, *port)
		if err := serv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
