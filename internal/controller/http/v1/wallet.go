package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/internal/usecase"
	"github.com/svbnbyrk/wallet/pkg/logger"
)

type walletRoutes struct {
	uc usecase.Wallet
	l  logger.Interface
}

func newWalletRoutes(handler *gin.RouterGroup, t usecase.Wallet, l logger.Interface) {

	r := &walletRoutes{t, l}
	h := handler.Group("/wallet")
	{
		h.POST("", r.post)
	}
	handler.GET("user/:id/wallet", r.get)
}

type walletPostRequest struct {
	UserId   int64  `json:"user_id"`
	Currency string `json:"currency" binding:"required,iso4217"`
	Balance  float64
}

func (ur *walletRoutes) post(c *gin.Context) {
	var w walletPostRequest
	if err := c.ShouldBind(&w); err != nil {
		ur.l.Error(err, "http - v1 - shouldbind")
		for _, fieldErr := range err.(validator.ValidationErrors) {
			c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprint(fieldErr))
		}
	}

	wallet := entity.NewWallet(w.UserId, w.Currency, w.Balance)
	err := ur.uc.Store(c, wallet)
	if err != nil {
		ur.l.Error(err, "http - v1 - store")
		ErrorResponse(c, err)

		return
	}
}

type walletsResponse struct {
	Wallets []entity.Wallet `json:"wallets"`
}

func (ur *walletRoutes) get(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ur.l.Error(err, "http - v1 - store")
		ErrorResponse(c, err)

		return
	}
	wallets, err := ur.uc.GetWalletsbyUser(c, int64(id))
	if err != nil {
		ur.l.Error(err, "http - v1 - store")
		ErrorResponse(c, err)

		return
	}
	c.JSON(http.StatusOK, walletsResponse{wallets})
}
