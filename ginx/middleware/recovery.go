package middleware

import (
	"fmt"
	"gitee.com.egcode.plugins/eglog"
	"gitee.com.egcode.plugins/ginx/response"
	"github.com/gin-gonic/gin"
	"runtime"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				eglog.Error().Str("stack", stack(3)).Msgf("%v", err)
				response.SendResponse(c, err, nil)
				c.Abort()
			}
		}()
		c.Next()
	}
}
func stack(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%s :%d", file, line)
}
