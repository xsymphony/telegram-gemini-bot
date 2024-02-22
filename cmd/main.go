package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsymphony/telegram-gemini-bot/application"
)

func main() {
	gin.SetMode(gin.DebugMode)
	g := gin.Default()
	// 加载静态文件
	g.LoadHTMLGlob("../application/template/*")
	application.Register(g)
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404")
	})
	g.Run(":3333")
}
