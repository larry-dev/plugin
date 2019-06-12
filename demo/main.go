package main

import (
	"gitee.com/egcode/plugins/eglog"
	"gitee.com/egcode/plugins/ginx"
	"gitee.com/egcode/plugins/http/httpserver"
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
	//生成pem和key
	//openssl req -new -nodes -x509 -out server.pem -keyout server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.egcode.cn/emailAddress=jclarry@qq.com"
	if err := httpserver.NewServer(httpserver.Config{
		ListenAddr: ":8080",
		Handler:    engine,
		EnableHttps:true,
		KeyFile:"./server.key",
		CertFile:"./server.pem",
	}); err != nil {
		panic(err)
	}
}
