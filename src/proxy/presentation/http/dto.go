package http

type GetBalanceResponse struct {
	Balance string `json:"balance,omitempty"`
	Error   string `json:"error,omitempty"`
}
