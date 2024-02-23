package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsymphony/telegram-gemini-bot/application"
	gin_extra "github.com/xsymphony/telegram-gemini-bot/pkg/gin-extra"
)

var (
	gateway *gin.Engine
)

func init() {
	gateway = gin.Default()
	gin_extra.LoadHTMLFromEmbedFS(gateway, application.EmbedFS, "templates/*html")
	application.Register(gateway)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	gateway.ServeHTTP(w, r)
}
