package route

import (
	"fmt"
	"net/http"
	"time"
	"victoria-falls/internal/middleware"
	handlers "victoria-falls/src/handler"
	author "victoria-falls/src/usecase/authors"

	"github.com/gorilla/mux"
)

func InitializeRoutes(a author.AuthorUsecase) *http.Server {
	r := mux.NewRouter()
	am := middleware.AuthenticationMiddleware{}
	am.Populate()

	r.Use(am.Middleware)

	r.Handle("/api/v1/authors", handlers.GetAll(r, a)).Methods("GET")
	r.Handle("/api/v1/authors", handlers.Create(r, a)).Methods("POST")

	r.MethodNotAllowedHandler = MethodNotAllowedHandler()

	return &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
}

func MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Method not allowed")
	})
}
