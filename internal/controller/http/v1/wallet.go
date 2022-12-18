package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/internal/usecase"
	"github.com/svbnbyrk/wallet/pkg/logger"
)

type walletRoutes struct {
	uc usecase.WalletUsecase
	l  logger.Interface
}

func newWalletRoutes(handler *gin.RouterGroup, t usecase.WalletUsecase, l logger.Interface) {

	r := &walletRoutes{t, l}
	handler.GET("user/:id/wallet", r.get)
	h := handler.Group("/wallet")
	{
		h.POST("", r.post)
	}

}

func (ur *walletRoutes) post(c *gin.Context) {
	var w entity.Wallet
	if err := c.ShouldBind(&w); err != nil {
		ur.l.Error(err, "http - v1 - shouldbind")
		c.String(http.StatusBadRequest, "")
	}
	err := ur.uc.Store(c, w)
	if err != nil {
		ur.l.Error(err, "http - v1 - store")
		errorResponse(c, http.StatusInternalServerError, "Unexpected Error")

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
		errorResponse(c, http.StatusBadRequest, "Url parameter cannot parse integer")

		return
	}
	wallets, err := ur.uc.GetWalletsbyUser(c, int64(id))
	if err != nil {
		ur.l.Error(err, "http - v1 - store")
		errorResponse(c, http.StatusInternalServerError, "Unexpected Error")

		return
	}
	c.JSON(http.StatusOK, walletsResponse{wallets})
}
