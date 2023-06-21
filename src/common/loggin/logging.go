package loggin

import (
	"context"

	"github.com/sirupsen/logrus"
)

const contextKeyLogger = "LOGGER"

func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, contextKeyLogger, logger)
}

func GetLogger(ctx context.Context) *logrus.Entry {
	if logger, ok := ctx.Value(contextKeyLogger).(*logrus.Entry); ok {
		return logger
	}
	return logrus.NewEntry(logrus.StandardLogger())
}

func WithLoggerFields(ctx context.Context, logger *logrus.Entry) (context.Context, *logrus.Entry) {
	return WithLogger(ctx, logger), logger
}
