package model

type Pool struct {
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Fee          string `json:"fee"`
	TickSpacing  string `json:"tick_spacing"`
	Address      string `json:"address"`
	CoinAAddress string `json:"coin_a_address"`
	CoinBAddress string `json:"coin_b_address"`
	IsClosed     bool   `json:"is_closed"`
	Price        string `json:"price"`
}

type Data struct {
	Pools []Pool `json:"lp_list"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    Data   `json:"data"`
}
