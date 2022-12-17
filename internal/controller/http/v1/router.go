package v1

import (
	"github.com/svbnbyrk/wallet/pkg/logger"
	"github.com/svbnbyrk/wallet/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.TransactionUsecase) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")
	{
		newTransactionRoutes(h, t, l)
	}
}