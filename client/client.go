package client

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/qerdcv/bombom.info/config"
	"github.com/qerdcv/bombom.info/domain"
)

var (
	ErrNotFound = errors.New("no club found")
)

type BS struct {
	*resty.Client
}

func New(conf config.BS) *BS {
	return &BS{
		resty.New().
			SetHeader("Content-Type", "application/json").
			SetAuthToken(conf.Token).
			SetBaseURL(conf.BaseURL),
	}
}

func (bs *BS) GetClubByTag(tag string) (club domain.Club, err error) {
	resp, err := bs.R().SetResult(&club).SetPathParam("clubTag", tag).Get("/clubs/{clubTag}")
	if err != nil {
		return club, fmt.Errorf("get clubs by tag: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return club, ErrNotFound
	}

	return club, nil
}

func (bs *BS) GetPlayerByTag(tag string) (player domain.Player, err error) {
	resp, err := bs.R().SetResult(&player).SetPathParam("playerTag", tag).Get("/players/{playerTag}")
	if err != nil {
		return player, fmt.Errorf("get clubs by tag: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return player, ErrNotFound
	}

	return player, nil
}
