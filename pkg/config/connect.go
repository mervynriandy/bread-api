package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:tectagon@/bookstore")
	if err != nil {
		fmt.Println(err)
	}

	return db
}
