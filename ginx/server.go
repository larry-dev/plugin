package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/larry-dev/plugins/ginx/middleware"
	"github.com/spf13/viper"
)

func NewEngine() *gin.Engine {
	gin.SetMode(viper.GetString("run_mode"))
	engine := gin.New()
	engine.Use(middleware.Recovery())
	return engine
}
