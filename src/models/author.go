package models

import (
	"bread-api/helper"
	"database/sql"
	"time"
)

type Author struct {
	ID        int        `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
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
		helper.ConvertAssign(author, values)
		authors = append(authors, author)
	}

	return authors, nil
}
