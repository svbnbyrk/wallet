package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/internal/usecase"
	"github.com/svbnbyrk/wallet/pkg/logger"
)

type userRoutes struct {
	uc usecase.UserUsecase
	l  logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.UserUsecase, l logger.Interface) {

	r := &userRoutes{t, l}

	h := handler.Group("/user")
	{
		h.POST("", r.post)
	}
}

func (ur *userRoutes) post(c *gin.Context) {
	var us entity.User
	if err := c.ShouldBind(&us); err != nil {
		ur.l.Error(err, "http - v1 - shouldbind")
		c.String(http.StatusBadRequest, "")
	}
	err := ur.uc.Store(c, us)
	if err != nil {
		ur.l.Error(err, "http - v1 - store")
		errorResponse(c, http.StatusInternalServerError, "Unexpected Error")

		return
	}
}
