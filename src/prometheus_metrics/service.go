package prometheusmetrics

import (
	"time"

	"github.com/apm-dev/eth_getBalance-proxy/src/domain"
	"github.com/prometheus/client_golang/prometheus"
)

type Service struct {
	rpsCount       *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	responseTimeMs *prometheus.GaugeVec
}

func NewService(prefix string) domain.PrometheusMetrics {
	s := &Service{}
	s.rpsCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "rpc",
		},
		[]string{"method"},
	)
	// internal errors
	s.errorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "error",
		},
		[]string{"method", "error"},
	)
	// response time per method
	s.responseTimeMs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: prefix + "response_time_ms",
		},
		[]string{"method"},
	)
	prometheus.MustRegister(
		s.rpsCount,
		s.errorCount,
		s.responseTimeMs,
	)
	return s
}

func (s *Service) AddRpsCount(op string) {
	s.rpsCount.WithLabelValues(op).Add(1)
}

func (s *Service) AddErrCount(op string, err string) {
	s.errorCount.WithLabelValues(op, err).Add(1)
}

func (s *Service) AggregateResponseTimeDeferred(op string, start *time.Time) {
	s.responseTimeMs.WithLabelValues(op).Set(float64(time.Since(*start).Milliseconds()))
}
