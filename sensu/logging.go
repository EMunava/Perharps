package sensu

import (
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s loggingService) handleSensu(ctx context.Context, chatId uint32, sensu SensuMessageRequest) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "handleSensu",
			"message", sensu,
			"chat", chatId,
			"took", time.Since(begin),
		)
	}(time.Now())
	s.Service.handleSensu(ctx, chatId, sensu)
}
