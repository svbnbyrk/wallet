package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/internal/usecase"
	"github.com/svbnbyrk/wallet/pkg/logger"
)

type transactionRoutes struct {
	uc usecase.TransactionUsecase
	l  logger.Interface
}

func newTransactionRoutes(handler *gin.RouterGroup, t usecase.TransactionUsecase, l logger.Interface) {

	r := &transactionRoutes{t, l}

	h := handler.Group("/transaction")
	{
		h.GET("/history", r.history)
		h.POST("", r.post)
	}
}

type historyResponse struct {
	History []entity.Transaction `json:"history"`
}

func (tr *transactionRoutes) history(c *gin.Context) {
	transactions, err := tr.uc.History(c.Request.Context())
	if err != nil {
		tr.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "Unexpected Error")

		return
	}

	c.JSON(http.StatusOK, historyResponse{transactions})
}

type postRequest struct {
	WalletId        int64                  `json:"wallet_id" binding:"required"`
	TransactionType entity.TransactionType `json:"transactionType" binding:"required"`
	Currency        string                 `json:"currency" binding:"required"`
	Amount          float64                `json:"amount" binding:"required"`
}

func (tr *transactionRoutes) post(c *gin.Context) {
	var pr postRequest
	if err := c.ShouldBind(&pr); err != nil {
		tr.l.Error(err, "http - v1 - history")
		c.String(http.StatusBadRequest, "")
	}

	transaction := entity.NewTransaction(pr.WalletId, pr.TransactionType, pr.Currency, pr.Amount)
	err := tr.uc.Post(c, transaction)
	if err != nil {
		tr.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "Unexpected Error")

		return
	}
}
