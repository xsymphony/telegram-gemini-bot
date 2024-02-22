package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func welcome(c *gin.Context) {
	resp, err := Bot().client.GetWebhookInfo()
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func setWebhook(c *gin.Context) {
	if err := Bot().StartWebhook(); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success."})
}

func stopWebhook(c *gin.Context) {
	if err := Bot().StopWebhook(); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success."})
}

func recvMessage(c *gin.Context) {
	var update tgbotapi.Update
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	if err := Bot().RecvMessage(&update); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success."})
}

func Register(g *gin.Engine) {
	g.GET("/", welcome)
	g.POST("/message", recvMessage)
	g.POST("/webhook", setWebhook)
	g.DELETE("/webhook", stopWebhook)
}
