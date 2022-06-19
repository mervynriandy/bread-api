package author_repo

import (
	"fmt"
	"log"
	"strings"
	"victoria-falls/src/models"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

type AuthorRepositoryDB struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func AuthorRepo(db *sqlx.DB, logger *zap.Logger) AuthorRepository {
	if db == nil {
		logger.Fatal("error: db cannot be nil")
	}

	return AuthorRepositoryDB{
		db:     db,
		logger: logger,
	}
}

func (ar AuthorRepositoryDB) Create(author *models.Author) {
	query := "INSERT INTO authors(first_name, last_name) VALUES(?, ?);"
	rows, err := ar.db.Query(query, author.FirstName, author.LastName)
	if err != nil {
		log.Printf(`Error v2 with: %s`, err)
	} else {
		fmt.Println("author: ", rows)
	}
}

func (ar AuthorRepositoryDB) GetDetail() {

}

func (ar AuthorRepositoryDB) GetAll(author *[]models.Author) {
	ar.db.MapperFunc(strings.ToUpper)
	err := ar.db.Select(author, "SELECT id, first_name, last_name FROM authors")
	if err != nil {
		log.Printf(`Error v2 with: %s`, err)
	} else {
		fmt.Println("author: ", author)
	}
}
