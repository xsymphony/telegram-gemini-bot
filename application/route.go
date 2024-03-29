package application

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func index(c *gin.Context) {
	webhookInfo, err := tgBot().client.GetWebhookInfo()
	if err != nil {
		c.HTML(200, "templates/index.html", gin.H{})
		return
	}
	v, _ := json.MarshalIndent(webhookInfo, "", "    ")
	c.HTML(200, "templates/index.html", gin.H{"WebhookInfo": string(v)})
}

func setWebhook(c *gin.Context) {
	if err := tgBot().StartWebhook(); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success."})
}

func stopWebhook(c *gin.Context) {
	if err := tgBot().StopWebhook(); err != nil {
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
	if err := tgBot().RecvMessage(&update); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success."})
}

func Register(g *gin.Engine) {
	g.GET("/", index)
	g.POST("/message", recvMessage)
	g.POST("/webhook", setWebhook)
	g.DELETE("/webhook", stopWebhook)
}
