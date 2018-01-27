package mem

import (
	_ "fmt"
	"time"

	"github.com/go-kit/kit/log"
)

var ()

type loggingService struct {
	logger log.Logger
	next   Member
}

func NewLoggingService(logger log.Logger, s Member) Member {
	return &loggingService{logger, s}
}
func (s *loggingService) Login(uid, pwd string) (bool, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "YOUR_METHOD_NAME",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.next.Login(uid, pwd)
}
func (s *loggingService) Logout(uid string) (bool, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "YOUR_METHOD_NAME",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.next.Logout(uid)
}
func (s *loggingService) Remark() error {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "YOUR_METHOD_NAME",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.next.Remark()
}
