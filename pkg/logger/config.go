package logger

import (
	"context"

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
	cfg := zap.NewDevelopmentConfig()
	if env == "production" {
		cfg = zap.NewProductionConfig()
	}

	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, _ := cfg.Build(options...)

	logger = logger.With(
		zap.String("server_service_name", service),
	)

	return logger
}

func NewRequest(ctx context.Context, logger *zap.Logger) context.Context {
	requestID, _ := uuid.NewV4()
	logger = logger.With(
		zap.String("server_request_id", requestID.String()),
	)

	return context.WithValue(ctx, loggerKey, logger)
}
