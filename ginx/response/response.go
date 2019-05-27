package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendResponse(c *gin.Context, err interface{}, data interface{}) {
	code, message := DecodeException(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Result{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
