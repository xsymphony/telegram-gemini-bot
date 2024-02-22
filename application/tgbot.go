package application

import (
	"log"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	client *tgbotapi.BotAPI
}

func newBot() func() *TgBot {
	var (
		once sync.Once
		bot  *TgBot
	)
	return func() *TgBot {
		once.Do(func() {
			token := os.Getenv("TGBOT_TOKEN")
			client, err := tgbotapi.NewBotAPI(token)
			if err != nil {
				log.Fatal(err)
			}
			client.Debug = true
			log.Printf("Authorized on account %s", client.Self.UserName)
			bot = &TgBot{
				client: client,
			}
		})
		return bot
	}
}

var Bot = newBot()

func (bot *TgBot) StartWebhook() error {
	domain := os.Getenv("DOMAIN")
	link := domain + "/message"
	wh, _ := tgbotapi.NewWebhook(link)
	_, err := bot.client.Request(wh)
	return err
}

func (bot *TgBot) StopWebhook() error {
	_, err := bot.client.Request(&tgbotapi.DeleteWebhookConfig{
		DropPendingUpdates: true,
	})
	return err
}
