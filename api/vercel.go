package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsymphony/telegram-gemini-bot/application"
)

var (
	gateway *gin.Engine
)

func init() {
	gateway = gin.Default()
	gateway.LoadHTMLGlob("application/template/*")
	application.Register(gateway)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	gateway.ServeHTTP(w, r)
}
