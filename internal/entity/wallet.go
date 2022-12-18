package entity

type Wallet struct {
	ID       int64   `json:"id"`
	UserId   int64   `json:"user_id"`
	Balance  float64 `json:"balance" binding:"required,numeric"`
	Currency string  `json:"currency" binding:"required,iso4217"`
}

func NewWallet(userId int64, currency string, balance float64) Wallet {
	return Wallet{
		UserId:   userId,
		Currency: currency,
		Balance:  balance,
	}
}
