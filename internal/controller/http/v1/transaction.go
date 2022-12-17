package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/internal/usecase"
	"github.com/svbnbyrk/wallet/pkg/logger"
)

type transactionRoutes struct {
	t usecase.TransactionUsecase
	l logger.Interface
}

func newTransactionRoutes(handler *gin.RouterGroup, t usecase.TransactionUsecase, l logger.Interface) {
	r := &transactionRoutes{t, l}

	h := handler.Group("/transaction")
	{
		h.GET("/history", r.history)
		h.POST("/", r.post)
	}
}

type historyResponse struct {
	History []entity.Transaction `json:"history"`
}

func (r *transactionRoutes) history(c *gin.Context) {
	transactions, err := r.t.History(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "Unexpected Error")

		return
	}

	c.JSON(http.StatusOK, historyResponse{transactions})
}

func (r *transactionRoutes) post(c *gin.Context) {
	var transaction entity.Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		r.l.Error(err, "http - v1 - history")
		c.String(http.StatusBadRequest, "")
	}

	err := r.t.Post(c, transaction)
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "Unexpected Error")

		return
	}
}
