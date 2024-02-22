package application

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	client *tgbotapi.BotAPI
}

func newTagBot() func() *TgBot {
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
			log.Printf("Authorized on account %s\n", client.Self.UserName)
			bot = &TgBot{
				client: client,
			}
		})
		return bot
	}
}

var tgBot = newTagBot()

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
	isUseful := bot.filterUselessGroupMessage(update)
	if !isUseful {
		return nil
	}
	if err := bot.handleText(update); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (bot *TgBot) handleCommand(update *tgbotapi.Update) error {
	if !update.Message.IsCommand() {
		return nil
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "clear":
		msg.Text = "会话记录已清除，开启一场新的对话吧。"
		sessions.delete(update.Message.From.ID)
	default:
		msg.Text = "不支持的命令，目前支持：/clear"
	}

	if _, err := bot.client.Send(msg); err != nil {
		return err
	}

	return nil
}

func (bot *TgBot) handleText(update *tgbotapi.Update) error {
	if update.Message.IsCommand() {
		return nil
	}

	if update.Message == nil || update.Message.Text == "" {
		return nil
	}

	log.Printf("[telegram]%s: %s\n", update.Message.From.UserName, update.Message.Text)

	_ = bot.typing(update.Message.Chat.ID)

	reply, err := sessions.Ask(update.Message.From.ID, update.Message.Text)
	if err != nil {
		return bot.reply(update.Message.MessageID, update.Message.Chat.ID, fmt.Sprintf("gemini响应错误:%s", err.Error()))
	}

	return bot.reply(update.Message.MessageID, update.Message.Chat.ID, reply)
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
	log.Printf("[telegram]%s: %s\n", bot.client.Self.UserName, content)
	return err
}

func (bot *TgBot) isAtMe(update *tgbotapi.Update) bool {
	for _, entity := range update.Message.Entities {
		if entity.Type == "mention" {
			if strings.Contains(update.Message.Text, bot.client.Self.UserName) {
				return true
			}
		}
	}
	return false
}

func (bot *TgBot) isReplyMe(update *tgbotapi.Update) bool {
	if update.Message != nil && update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.From.UserName == bot.client.Self.UserName {
		return true
	}
	return false
}

func (bot *TgBot) filterUselessGroupMessage(update *tgbotapi.Update) bool {
	// 私聊消息不做特殊处理
	if update.FromChat().IsPrivate() {
		return true
	}
	if bot.isAtMe(update) {
		update.Message.Text = strings.ReplaceAll(update.Message.Text, "@"+bot.client.Self.UserName, "")
		return true
	}
	if bot.isReplyMe(update) {
		return true
	}
	return false
}
