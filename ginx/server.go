package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/larry-dev/plugins/ginx/middleware"
	"github.com/spf13/viper"
)

// 服务端引擎
func NewServerEngine() *gin.Engine {
	gin.SetMode(viper.GetString("run_mode"))
	engine := gin.New()
	engine.Use(middleware.Recovery())
	if viper.GetBool("cors.enable") {
		engine.Use(middleware.CORS())
	}

	return engine
}
