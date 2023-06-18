package http

import (
	"fmt"
	"net/http"

	"github.com/apm-dev/eth_getBalance-proxy/src/common"
	"github.com/apm-dev/eth_getBalance-proxy/src/domain"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ProxyHandler struct {
	rpcProxyService domain.RpcProxyService
	validator       *validator.Validate
}

func NewProxyHandler(e *echo.Echo, rps domain.RpcProxyService) {
	handler := &ProxyHandler{
		rpcProxyService: rps,
		validator:       validator.New(),
	}
	e.GET("/eth/balance/:address", handler.eth_getBalance)
}

func (h *ProxyHandler) eth_getBalance(c echo.Context) error {
	address := c.Param("address")
	err := h.validator.Var(address, "required,eth_addr")
	if err != nil {
		return c.JSON(http.StatusBadRequest, &GetBalanceResponse{Error: "'address' path param should be a valid eth address"})
	}
	req, err := domain.NewJsonRpcRequest(1, "eth_getBalance", []string{address, "latest"})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &GetBalanceResponse{Error: err.Error()})
	}
	resp, err := h.rpcProxyService.SendRequest(c.Request().Context(), "eth", req)
	if err != nil {
		logrus.Errorf("%s: %s", "ProxyHandler.eth_getBalance", err.Error())
		code, msg := common.ErrToHttpCodeAndMessage(err, "eth_getBalance")
		return c.JSON(code, &GetBalanceResponse{Error: msg})
	}
	if resp == nil {
		return c.JSON(http.StatusOK, &GetBalanceResponse{Error: "no balance data"})
	}
	if resp.Error != nil {
		return c.JSON(http.StatusOK, &GetBalanceResponse{Error: fmt.Sprintf("%d : %s", resp.Error.Code, resp.Error.Message)})
	}
	balance, err := common.HexToInt(resp.Result)
	if err != nil {
		logrus.Errorf("%s: %s", "ProxyHandler.eth_getBalance", err.Error())
		return c.JSON(http.StatusInternalServerError, "failed to parse the balance")
	}
	return c.JSON(http.StatusOK, &GetBalanceResponse{Balance: balance})
}
