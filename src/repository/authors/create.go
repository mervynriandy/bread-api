package author_repo

import (
	"bread-api/src/models"
	"fmt"

	"go.uber.org/zap"
)

func (ar AuthorRepositoryDB) Create(author *models.Author) {
	query := "INSERT INTO authors(name) VALUES(?);"
	rows, err := ar.db.Query(query, author.Name)
	if err != nil {
		zap.L().Error(`Error, `, zap.Error(err))
	} else {
		fmt.Println("author: ", rows)
	}
}
