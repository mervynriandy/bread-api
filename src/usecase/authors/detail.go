package author_usecase

import (
	"bread-api/src/models"
	"context"
)

func (a AuthorUsecase) DetailAuthor(context context.Context, id string) models.Author {
	author := &models.Author{}
	a.authorRepository.GetDetail(author, id)
	return *author
}
