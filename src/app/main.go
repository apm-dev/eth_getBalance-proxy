package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ReneKroon/ttlcache"
	"github.com/apm-dev/eth_getBalance-proxy/src/common"
	"github.com/apm-dev/eth_getBalance-proxy/src/config"
	_nodeRepo "github.com/apm-dev/eth_getBalance-proxy/src/node/data/repo"
	prometheusmetrics "github.com/apm-dev/eth_getBalance-proxy/src/prometheus_metrics"
	_prometheusHttp "github.com/apm-dev/eth_getBalance-proxy/src/prometheus_metrics/presentation/http"
	_proxyService "github.com/apm-dev/eth_getBalance-proxy/src/proxy"
	"github.com/apm-dev/eth_getBalance-proxy/src/proxy/data/cache"
	_proxyHttp "github.com/apm-dev/eth_getBalance-proxy/src/proxy/presentation/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.NewConfig()

	logLevel, err := logrus.ParseLevel(config.App.LogLevel)
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(logLevel)

	prometheus := prometheusmetrics.NewService(strings.ReplaceAll(config.App.ServiceName, "-", "_") + "__")

	e := echo.New()
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// Healthz
	e.GET("/healthz", func(c echo.Context) error {
		// For liveness check, simply return HTTP 200 OK
		// indicating that the application is live and healthy.

		// For readiness check, we can perform additional checks
		// such as database connectivity, external service dependencies,
		// or other necessary components, and return HTTP 200 OK only
		// when the application is ready to serve traffic.
		return c.NoContent(http.StatusOK)
	})

	nodeRepo := _nodeRepo.NewNodeRepository()

	ttlCache := ttlcache.NewCache()
	ttlCache.SkipTtlExtensionOnHit(true)

	rpcProxyCache := cache.NewRpcProxyCache(ttlCache)

	rpcProxy := _proxyService.NewRpcProxyService(config, rpcProxyCache, nodeRepo)

	_prometheusHttp.RegisterPrometheusHandlers(e, prometheus)
	_proxyHttp.RegisterProxyHandlers(e, rpcProxy, prometheus)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", config.App.WebPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	common.WaitForSignal()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logrus.Fatal(err)
	}
}
