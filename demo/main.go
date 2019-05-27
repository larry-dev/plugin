package main

import (
	"gitee.com.egcode.plugins/eglog"
	"gitee.com.egcode.plugins/ginx"
	"gitee.com.egcode.plugins/http/httpserver"
	"github.com/gin-gonic/gin"
)
func init() {
	eglog.InitLog()
}
func main() {
	engine := ginx.NewEngine()
	engine.GET("/", func(context *gin.Context) {
		context.String(200, "hello world")
	})
	if err := httpserver.NewServer(httpserver.Config{
		ListenAddr: ":8080",
		Handler:    engine,
	}); err != nil {
		panic(err)
	}
}
