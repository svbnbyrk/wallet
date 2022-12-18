package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/svbnbyrk/wallet/internal/usecase"
	"github.com/svbnbyrk/wallet/pkg/logger"
)

func NewRouter(handler *gin.Engine, l logger.Interface, tuc usecase.TransactionUsecase, uuc usecase.UserUsecase, wuc usecase.WalletUsecase) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")
	{
		newTransactionRoutes(h, tuc, l)
		newUserRoutes(h, uuc, l)
		newWalletRoutes(h, wuc, l)
	}
}
