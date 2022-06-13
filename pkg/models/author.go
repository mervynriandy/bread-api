package models

import (
	"database/sql"
	"strconv"

	"victoria-falls/pkg/config"
)

type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var db *sql.DB = config.Connect()

func AuthorRowMapper(rows *sql.Rows) ([]Author, error) {
	var authors []Author

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var author Author
		id, _ := strconv.Atoi(string(values[0]))
		author.ID = id
		author.FirstName = string(values[1])
		author.LastName = string(values[2])

		authors = append(authors, author)
	}

	return authors, nil
}

func GetAllAuthor() ([]Author, error) {
	rows, err := db.Query("SELECT * FROM authors;")
	if err != nil {
		return nil, err
	}

	authors, err := AuthorRowMapper(rows)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func GetAuthorById(id int) (Author, error) {
	rows, err := db.Query("SELECT * FROM authors WHERE id = ?;", id)
	if err != nil {
		return Author{}, err
	}

	authors, err := AuthorRowMapper(rows)
	if err != nil {
		return Author{}, err
	}

	return authors[0], nil
}

func CreateAuthor(author Author) error {
	query := "INSERT INTO authors(first_name, last_name) VALUES(?, ?);"
	_, err := db.Query(query, author.FirstName, author.LastName)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAuthor(author Author, id int) error {
	query := "UPDATE authors SET first_name = ?, last_name = ? WHERE id = ?;"
	_, err := db.Query(query, author.FirstName, author.LastName, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAuthor(id int) error {
	query := "DELETE FROM authors WHERE id = ?;"
	_, err := db.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

func AuthorExists(id int) (bool, error) {
	rows, err := db.Query("SELECT * FROM authors WHERE id = ?;", id)
	if err != nil {
		return false, err
	}

	authors, err := AuthorRowMapper(rows)
	if err != nil {
		return false, err
	}

	if len(authors) < 1 {
		return false, nil
	}

	return true, nil
}
