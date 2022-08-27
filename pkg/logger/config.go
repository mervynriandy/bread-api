package logger

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerKeyType struct{}

var loggerKey loggerKeyType

func GetLogger(ctx context.Context) *zap.Logger {
	return ctx.Value(loggerKey).(*zap.Logger)
}

func InitZapLogger(env, service string, options ...zap.Option) *zap.Logger {
	var logger *zap.Logger

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	var zapCore []zapcore.Core

	if env == "production" {
		if _, err := os.Stat("log"); os.IsNotExist(err) {
			err := os.Mkdir("log", os.ModePerm)
			if err != nil {
				fmt.Println("Error Create Folder")
			}
		}

		logFile, _ := os.OpenFile("./log/log-"+dayLog()+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		writer := zapcore.AddSync(logFile)

		fileEncoder := zapcore.NewJSONEncoder(config)
		defaultLogLevel := zapcore.InfoLevel
		zapCore = append(zapCore, zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
	} else {
		defaultLogLevel := zapcore.DebugLevel
		zapCore = append(zapCore, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel))
	}

	logger = zap.New(zapcore.NewTee(
		zapCore...,
	), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}

func dayLog() string {
	year := time.Now().Year()
	month := time.Now().Month().String()
	day := time.Now().Day()

	return strconv.Itoa(day) + "-" + month + "-" + strconv.Itoa(year)
}

func NewRequest(ctx context.Context, logger *zap.Logger) context.Context {
	requestID, _ := uuid.NewV4()
	logger = logger.With(
		zap.String("server_request_id", requestID.String()),
	)

	return context.WithValue(ctx, loggerKey, logger)
}
