package service

import (
	"github.com/qerdcv/bombom.info/bot"
	"github.com/qerdcv/bombom.info/client"
)

type Service struct {
	client *client.BS
	bot    *bot.Bot
}

func New(c *client.BS, b *bot.Bot) *Service {
	return &Service{
		c,
		b,
	}
}
