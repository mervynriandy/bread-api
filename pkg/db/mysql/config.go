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
	cfg.dbName = os.Getenv("MYSQL_NAME")

	if err != nil {
		logger.Error(fmt.Sprintf("error: cannot convert mysql port from string to int: %s", err.Error()))
		return nil, err
	}

	cfg.dsn = fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)

	return &cfg, nil
}

func (c *config) Dsn() string {
	return c.dsn
}

func (c *config) DBName() string {
	return c.dbName
}
