package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"victoria-falls/internal/middleware"
	"victoria-falls/pkg/handlers"

	"github.com/gorilla/mux"
)

const PORT = 8080

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	am := middleware.AuthenticationMiddleware{}
	am.Populate()

	r.Use(am.Middleware)

	r.HandleFunc("/api/v1/authors", handlers.GetAllAuthors).Methods("GET")
	r.HandleFunc("/api/v1/authors/{id}", handlers.GetAuthorById).Methods("GET")
	r.HandleFunc("/api/v1/authors", handlers.CreateAuthor).Methods("POST")
	r.HandleFunc("/api/v1/authors/{id}", handlers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/api/v1/authors/{id}", handlers.DeleteAuthor).Methods("DELETE")

	r.HandleFunc("/api/v1/books", handlers.GetAllBooks).Methods("GET")
	r.HandleFunc("/api/v1/books/{id}", handlers.GetBookById).Methods("GET")
	r.HandleFunc("/api/v1/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/api/v1/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/v1/books/{id}", handlers.DeleteBook).Methods("DELETE")

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
