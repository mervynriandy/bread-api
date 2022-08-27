package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"bread-api/helper"
	route "bread-api/internal/routes"
	"bread-api/pkg/db/mysql"
	"bread-api/pkg/logger"
	ar "bread-api/src/repository/authors"
	auc "bread-api/src/usecase/authors"

	"github.com/joho/godotenv"

	"go.uber.org/zap"
)

const PORT = 8080

func main() {

	// Initiate default variables
	var (
		env   = helper.GetEnv("SERVICE_ENV", "development")
		srcnm = helper.GetEnv("SERVICE_NAME", "bread-api")
		wait  time.Duration
	)
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Initiate logger using zap
	// Accepts env level and service name
	// to be used for logging
	logger := logger.InitZapLogger(env, srcnm)
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

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

	db_conn, err := mysql.NewConnection(mysqlConfig)
	if err != nil {
		logger.Error("error: database connection",
			zap.Error(err))
		errs <- err
	}

	authorRepo := ar.AuthorRepo(db_conn)
	authorUsecase := auc.AuthorCase(authorRepo)

	srv := route.InitializeRoutes(*authorUsecase)
	logger.Info("Server is ready to handle requests at " + helper.GetEnv("SERVICE_PORT", "8080"))

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
