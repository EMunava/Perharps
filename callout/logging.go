package callout

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s loggingService) InvokeCallout(ctx context.Context, chatId uint32, title, message string, variables map[string]string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "InvokeCallout",
			"title", title,
			"message", message,
			"chat", chatId,
			"variables", variables,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.InvokeCallout(ctx, chatId, title, message, variables)
}
