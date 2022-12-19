package v1

import (
	"database/sql"
	"net/http"

	"github.com/svbnbyrk/wallet/internal/entity"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(getStatusCode(err), ResponseError{err.Error()})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case entity.ErrInternalServerError:
		return http.StatusInternalServerError
	case entity.ErrNotFound:
		return http.StatusNotFound
	case entity.ErrConflict:
		return http.StatusConflict
	case sql.ErrNoRows:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
