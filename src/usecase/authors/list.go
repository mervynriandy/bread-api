package author_usecase

import (
	"context"
	"victoria-falls/src/models"
)

func (a AuthorUsecase) ListAuthor(context context.Context) []models.Author {
	author := &[]models.Author{}
	a.authorRepository.GetAll(author)
	return *author
}
