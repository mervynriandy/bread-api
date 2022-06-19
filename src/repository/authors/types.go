package author_repo

import "victoria-falls/src/models"

type AuthorRepository interface {
	Create(author *models.Author)
	GetDetail()
	GetAll(author *[]models.Author)
}
