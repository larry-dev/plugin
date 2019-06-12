package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回json数据
func RespJson(c *gin.Context, err interface{}, data interface{}) {
	code, message := DecodeException(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Result{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
