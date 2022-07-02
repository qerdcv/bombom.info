package config

type Bot struct {
	Token  string `split_words:"true" required:"true"`
	ChatID int64  `split_words:"true" required:"true" default:"-677364817"`
}
