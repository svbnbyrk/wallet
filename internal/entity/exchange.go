package entity

type Exchange struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

func NewExchange(currency string, rate float64) Exchange {
	return Exchange{
		Currency: currency,
		Rate:     rate,
	}
}
