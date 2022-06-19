package author_usecase

import (
	"context"
	"victoria-falls/src/models"
)

func (a AuthorUsecase) CreateAuthor(context context.Context, author models.Author) {
	a.authorRepository.Create(&author)
}
