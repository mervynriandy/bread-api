package author_usecase

import (
	"bread-api/src/models"
	"context"
)

func (a AuthorUsecase) CreateAuthor(context context.Context, author models.Author) {
	a.authorRepository.Create(&author)
}
