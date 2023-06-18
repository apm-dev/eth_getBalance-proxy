package domain

import "time"

type PrometheusMetrics interface {
	AddRpsCount(op string)
	AddErrCount(op, err string)
	AggregateResponseTimeDeferred(op string, start *time.Time)
}
