package author_repo

import "bread-api/src/models"

type AuthorRepository interface {
	Create(author *models.Author)
	GetDetail(author *models.Author, id string)
	GetAll(author *[]models.Author)
}
