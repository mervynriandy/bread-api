package route

import (
	"bread-api/helper"
	"bread-api/internal/middleware"
	handlers "bread-api/src/handler"
	author "bread-api/src/usecase/authors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func InitializeRoutes(a author.AuthorUsecase) *http.Server {
	r := mux.NewRouter()
	am := middleware.AuthenticationMiddleware{}
	am.Populate()

	r.Use(am.Middleware)

	r.Handle("/api/v1/authors", handlers.GetAll(r, a)).Methods("GET")
	r.Handle("/api/v1/authors/{id:[0-9]+}", handlers.GetDetail(r, a)).Methods("GET")
	r.Handle("/api/v1/authors", handlers.Create(r, a)).Methods("POST")

	r.MethodNotAllowedHandler = MethodNotAllowedHandler()

	host := helper.GetEnv("SERVICE_HOST", "localhost")
	port := helper.GetEnv("SERVICE_ENV", "8080")
	return &http.Server{
		Addr:         host + ":" + port,
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
