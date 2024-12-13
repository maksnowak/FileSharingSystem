package main

import (
	"accounts/db"
	_ "accounts/docs"
	"accounts/handlers"
	"context"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		if os.Getenv("APP_ENV") != "prod" {
			addr := fmt.Sprintf("http://localhost:%v/swagger/", *port)
			r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(addr+"doc.json")))
			logger.Printf("Swagger UI available at %v\n", addr+"index.html")
		}
	}

	r.Route("/accounts", func(r chi.Router) {
		r.Post("/", handlers.Register)                 // POST register account
		r.Get("/", handlers.GetAllAccounts)            // GET retrieve all accounts
		r.Get("/{user_id}", handlers.GetAccountByID)   // GET account by ID
		r.Put("/{user_id}", handlers.UpdateAccount)    // PUT update an account
		r.Delete("/{user_id}", handlers.DeleteAccount) // DELETE an account
	})

	r.Route("/login", func(r chi.Router) {
		r.Get("/{username}", handlers.GetPasswordSalt) // GET the password salt of the user
		r.Get("/", handlers.Login)                     // GET try to log in the user
	})

	serv := &http.Server{Addr: *host + ":" + *port, Handler: r}
	go func() {
		logger.Printf("http: Listening on %v:%v\n", *host, *port)
		if err := serv.ListenAndServe(); err != nil {
			logger.Fatalln(err)
		}
	}()

	// DB USAGE
	db.Connect()
	defer db.Disconnect()

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
