package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api     *tgbotapi.BotAPI
	groupID int64
}

func NewBot(token string, groupID int64) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:     bot,
		groupID: groupID,
	}, nil
}

func (b *Bot) SendNotification(message string) error {
	msg := tgbotapi.NewMessage(b.groupID, message)
	msg.ParseMode = "HTML"
	_, err := b.api.Send(msg)
	return err
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "notify":
				// Проверяем, что команда отправлена в нужной группе
				if update.Message.Chat.ID == b.groupID {
					// Отправляем уведомление всем участникам группы
					msg := tgbotapi.NewMessage(b.groupID, "🔔 Уведомление от "+update.Message.From.FirstName+":\n"+update.Message.CommandArguments())
					msg.ParseMode = "HTML"
					b.api.Send(msg)
				}
			case "id":
				// Отправляем ID чата
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					fmt.Sprintf("ID этой группы: %d", update.Message.Chat.ID))
				b.api.Send(msg)
			}
		}
	}

	return nil
}
