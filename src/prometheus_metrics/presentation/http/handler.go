package http

import (
	"github.com/apm-dev/eth_getBalance-proxy/src/domain"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

type PrometheusHandler struct {
}

func RegisterPrometheusHandlers(
	e *echo.Echo,
	prometheus domain.PrometheusMetrics,
) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			prometheus.AddRpsCount(c.Request().Method + " " + c.Request().URL.Path)
			return next(c)
		}
	})
	e.GET("/metrics", echoprometheus.NewHandler())
}
