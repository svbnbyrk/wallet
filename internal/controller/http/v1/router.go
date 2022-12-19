package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/pkg/logger"
)

// Create router
func NewRouter(handler *gin.Engine, l logger.Interface, tuc entity.TransactionUseCase, uuc entity.UserUseCase, wuc entity.WalletUseCase) {
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
