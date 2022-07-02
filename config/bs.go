package config

// BS is for brawl stars
type BS struct {
	BaseURL string `split_words:"true" required:"true"`
	Token   string `split_words:"true" required:"true"`
}
