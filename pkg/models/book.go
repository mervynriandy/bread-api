package models

import (
	"database/sql"
	"strconv"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      Author `json:"author"`
}

type RequestBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorId    int    `json:"author_id"`
}

func BookRowMapper(rows *sql.Rows) ([]Book, error) {
	var books []Book

	columns, err := rows.Columns()
	if err != nil {
		return []Book{}, err
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

		var book Book
		id, _ := strconv.Atoi(string(values[0]))
		authorId, _ := strconv.Atoi(string(values[3]))
		author, _ := GetAuthorById(authorId)

		book.ID = id
		book.Title = string(values[1])
		book.Description = string(values[2])
		book.Author = author

		books = append(books, book)
	}

	return books, nil
}

func GetAllBooks() ([]Book, error) {
	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		return []Book{}, err
	}

	books, err := BookRowMapper(rows)
	if err != nil {
		return []Book{}, err
	}

	return books, nil
}

func GetBookById(id int) (Book, error) {
	rows, err := db.Query("SELECT * FROM books WHERE id = ?;", id)
	if err != nil {
		return Book{}, err
	}

	book, err := BookRowMapper(rows)
	if err != nil {
		return Book{}, err
	}

	return book[0], nil
}

func CreateBook(book Book) error {
	query := "INSERT INTO books(title, description, author_id) VALUES (?, ?, ?);"
	_, err := db.Query(query, book.Title, book.Description, book.Author.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateBook(book Book, id int) error {
	query := "UPDATE books SET title = ?, description = ?, author_id = ? WHERE id = ?;"
	_, err := db.Query(query, book.Title, book.Description, book.Author.ID, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(id int) error {
	_, err := db.Query("DELETE FROM books WHERE id = ?;", id)
	if err != nil {
		return err
	}

	return nil
}

func BookExists(id int) (bool, error) {
	rows, err := db.Query("SELECT * FROM books WHERE id = ?;", id)
	if err != nil {
		return false, err
	}

	books, err := BookRowMapper(rows)
	if err != nil {
		return false, err
	}

	if len(books) < 1 {
		return false, nil
	}

	return true, nil
}
