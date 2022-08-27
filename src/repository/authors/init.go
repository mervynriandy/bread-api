package author_repo

import (
	"bread-api/src/models"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

type AuthorRepositoryDB struct {
	db *sqlx.DB
}

func AuthorRepo(db *sqlx.DB) AuthorRepository {
	if db == nil {
		zap.L().Fatal("error: db cannot be nil")
	}

	return AuthorRepositoryDB{
		db: db,
	}
}

func (ar AuthorRepositoryDB) GetDetail(author *models.Author, id string) {
	ar.db.MapperFunc(strings.ToUpper)
	err := ar.db.Select(author, "SELECT * FROM authors WHERE id=?", id)
	if err != nil {
		zap.L().Error(`Error, `, zap.Error(err))
	} else {
		fmt.Println("author: ", author)
	}
}

func (ar AuthorRepositoryDB) GetAll(author *[]models.Author) {
	ar.db.MapperFunc(strings.ToUpper)
	err := ar.db.Select(author, "SELECT * FROM authors")
	if err != nil {
		zap.L().Error(`Error, `, zap.Error(err))
	} else {
		fmt.Println("author: ", author)
	}
}
