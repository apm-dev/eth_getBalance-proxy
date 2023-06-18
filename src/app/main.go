package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ReneKroon/ttlcache"
	"github.com/apm-dev/eth_getBalance-proxy/src/common"
	"github.com/apm-dev/eth_getBalance-proxy/src/config"
	_nodeRepo "github.com/apm-dev/eth_getBalance-proxy/src/node/data/repo"
	_proxyService "github.com/apm-dev/eth_getBalance-proxy/src/proxy"
	"github.com/apm-dev/eth_getBalance-proxy/src/proxy/data/cache"
	_proxyHttp "github.com/apm-dev/eth_getBalance-proxy/src/proxy/presentation/http"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.NewConfig()

	logLevel, err := logrus.ParseLevel(config.App.LogLevel)
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(logLevel)

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			return next(c)
		}
	})

	nodeRepo := _nodeRepo.NewNodeRepository()

	ttlCache := ttlcache.NewCache()
	ttlCache.SkipTtlExtensionOnHit(true)

	rpcProxyCache := cache.NewRpcProxyCache(ttlCache)

	rpcProxy := _proxyService.NewRpcProxyService(config, rpcProxyCache, nodeRepo)

	_proxyHttp.NewProxyHandler(e, rpcProxy)

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
