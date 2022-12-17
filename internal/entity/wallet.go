package entity

import (
	"github.com/google/uuid"
)

type Wallet struct {
	Id       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	Currency string `json:"currency"`
	Balance  float64 `json:"balance"`
}
