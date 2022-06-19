package models

import (
	"database/sql"
	"strconv"
)

type Author struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

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
