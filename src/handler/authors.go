package handlers

import (
	"bread-api/helper"
	appRes "bread-api/pkg/response"
	"bread-api/src/models"
	"context"
	"net/http"

	"github.com/gorilla/mux"

	authors "bread-api/src/usecase/authors"
)

func Create(h http.Handler, author authors.AuthorUsecase) http.Handler {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		aut := &models.Author{}
		helper.GetRequest(w, r, aut)

		author.CreateAuthor(ctx, *aut)

		helper.SendResponse(w, &appRes.Response{
			Code:     http.StatusOK,
			Status:   true,
			Messages: "Success",
		})
	})
	return handlerFunc
}

func GetAll(h http.Handler, author authors.AuthorUsecase) http.Handler {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		author := author.ListAuthor(ctx)

		helper.SendResponse(w, &appRes.Response{
			Code:     http.StatusOK,
			Status:   true,
			Messages: "Success",
			List:     author,
		})
	})
	return handlerFunc
}

func GetDetail(h http.Handler, author authors.AuthorUsecase) http.Handler {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		id := mux.Vars(r)["id"]
		author := author.DetailAuthor(ctx, id)

		helper.SendResponse(w, &appRes.Response{
			Code:     http.StatusOK,
			Status:   true,
			Messages: "Success",
			Detail:   author,
		})
	})
	return handlerFunc
}
