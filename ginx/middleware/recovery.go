package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/larry-dev/plugins/eglog"
	"github.com/larry-dev/plugins/ginx/response"
	"runtime"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				eglog.Error().Str("stack", stack(3)).Msgf("%v", err)
				response.RespJson(c, err, nil)
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
