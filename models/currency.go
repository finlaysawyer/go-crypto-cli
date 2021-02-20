package models

type Cryptocurrency struct {
	Data struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"data"`
}
