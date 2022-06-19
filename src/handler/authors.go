package handlers

import (
	"context"
	"net/http"
	"victoria-falls/helper"
	appRes "victoria-falls/pkg/response"
	"victoria-falls/src/models"

	authors "victoria-falls/src/usecase/authors"
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
			List:     aut,
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
