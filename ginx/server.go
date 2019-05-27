package ginx

import (
	"gitee.com.egcode.plugins/ginx/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewEngine() *gin.Engine {
	if !viper.GetBool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(middleware.Recovery())
	return engine
}
