package entity

import (
	"time"

	"github.com/google/uuid"
)

type Type int

const (
	Deposit Type = iota + 1
	Withdrawal
)

type Transaction struct {
	//TODO ID
	Id              uuid.UUID `json:"id"`
	WalletId        uuid.UUID `json:"wallet_id" binding:"required"`
	TransactionType Type      `json:"transactionType" binding:"required"`
	Currency        string    `json:"currency" binding:"required"`
	Amount          float64   `json:"amount" binding:"required"`
	Balance         float64   `json:"balance" binding:"required"`
	CreatedAt       time.Time `json:"created_at"`
}

func NewTransaction(walletId uuid.UUID, transactionType int, currency string, amount float64, balance float64) {

}
