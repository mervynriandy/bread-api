package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"victoria-falls/helper"
	route "victoria-falls/internal/routes"
	"victoria-falls/pkg/db/mysql"
	"victoria-falls/pkg/logger"
	ar "victoria-falls/src/repository/authors"
	auc "victoria-falls/src/usecase/authors"

	"github.com/joho/godotenv"

	"go.uber.org/zap"
)

const PORT = 8080

func main() {

	// Initiate default variables
	var (
		env   = helper.EnvString("SERVICE_ENV", "development")
		srcnm = helper.EnvString("SERVICE_NAME", "victoria-falls")
	)
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Initiate logger using zap
	// Accepts env level and service name
	// to be used for logging

	logger := logger.InitZapLogger(env, srcnm)
	s := godotenv.Load()
	if s != nil {
		logger.Fatal(fmt.Sprintf("error: loading .env file: %s", s.Error()))
	}

	errs := make(chan error, 3)

	// Initiate database connection
	mysqlConfig, err := mysql.NewConfig(logger)
	if err != nil {
		logger.Error("error: initiate database config",
			zap.Error(err))
		errs <- err
	}

	db_conn, err := mysql.NewConnection(logger, mysqlConfig)
	if err != nil {
		logger.Error("error: database connection",
			zap.Error(err))
		errs <- err
	}

	authorRepo := ar.AuthorRepo(db_conn, logger)
	authorUsecase := auc.AuthorCase(authorRepo, logger)

	srv := route.InitializeRoutes(*authorUsecase)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
