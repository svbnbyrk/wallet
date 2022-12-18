package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/internal/usecase"
	"github.com/svbnbyrk/wallet/pkg/logger"
)

type userRoutes struct {
	uc usecase.User
	l  logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.User, l logger.Interface) {

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
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errorResponse(c, http.StatusBadRequest, fmt.Sprint(fieldErr))
		}
		return
	}
	err := ur.uc.Store(c, us)
	if err != nil {
		ur.l.Error(err, "http - v1 - store")
		errorResponse(c, http.StatusInternalServerError, "Unexpected Error")

		return
	}
}
