package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/qerdcv/bombom.info/config"
	"github.com/qerdcv/bombom.info/domain"
)

type Bot struct {
	bApi *tgbotapi.BotAPI
	conf config.Bot
}

func New(conf config.Config) (*Bot, error) {
	b, err := tgbotapi.NewBotAPI(conf.Bot.Token)
	if err != nil {
		return nil, fmt.Errorf("tgbotapi new bot: %w", err)
	}

	b.Debug = conf.IsDevEnv()

	return &Bot{b, conf.Bot}, nil
}

func (b *Bot) SendJoinRequest(p domain.Player) error {
	poll := tgbotapi.NewPoll(b.conf.ChatID, p.String(), "accept", "deny")
	poll.IsAnonymous = false

	_, err := b.bApi.Send(poll)
	if err != nil {
		return fmt.Errorf("bot api send: %w", err)
	}

	return nil
}
