package mem

import (
	_ "fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

var ()

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram

	next Member
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Member) Member {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           s,
	}
}
func (s *instrumentingService) Remark() error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "book").Add(1)
		s.requestLatency.With("method", "book").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Remark()
}
func (s *instrumentingService) Login(uid, pwd string) (bool, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "book").Add(1)
		s.requestLatency.With("method", "book").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Login(uid, pwd)
}
func (s *instrumentingService) Logout(uid string) (bool, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "book").Add(1)
		s.requestLatency.With("method", "book").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Logout(uid)
}
