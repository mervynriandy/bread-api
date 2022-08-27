package author_usecase

import (
	"bread-api/src/models"
	"context"
)

func (a AuthorUsecase) ListAuthor(context context.Context) []models.Author {
	author := &[]models.Author{}
	a.authorRepository.GetAll(author)
	return *author
}
