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

func (bot *TgBot) RecvMessage(update *tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}
	if err := bot.handleCommand(update); err != nil {
		log.Println(err)
		return err
	}
	if err := bot.handleText(update); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (bot *TgBot) handleCommand(update *tgbotapi.Update) error {
	return nil
}

func (bot *TgBot) handleText(update *tgbotapi.Update) error {
	if update.Message.IsCommand() {
		return nil
	}

	if update.Message == nil || update.Message.Text == "" {
		return nil
	}

	log.Printf("[telegram]%s: %s", update.Message.From.UserName, update.Message.Text)

	return bot.reply(update.Message.MessageID, update.Message.Chat.ID, "收到消息")
}

func (bot *TgBot) typing(chatID int64) error {
	_, err := bot.client.Send(tgbotapi.NewChatAction(chatID, "typing"))
	return err
}

func (bot *TgBot) reply(replyMessageID int, chatID int64, content string) error {
	msg := tgbotapi.NewMessage(chatID, content)
	msg.ReplyToMessageID = replyMessageID
	msg.ParseMode = "MarkdownV2"
	_, err := bot.client.Send(msg)
	return err
}
