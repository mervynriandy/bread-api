package mysql

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

type Config interface {
	Dsn() string
	DBName() string
}

type config struct {
	dbName string
	dsn    string
}

func NewConfig(logger *zap.Logger) (Config, error) {
	var cfg config
	var err error

	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASS")
	dbName := os.Getenv("MYSQL_NAME")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	cfg.dbName = os.Getenv("MYSQL_NAME")

	if err != nil {
		logger.Error(fmt.Sprintf("error: cannot convert mysql port from string to int: %s", err.Error()))
		return nil, err
	}

	cfg.dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	return &cfg, nil
}

func (c *config) Dsn() string {
	return c.dsn
}

func (c *config) DBName() string {
	return c.dbName
}
