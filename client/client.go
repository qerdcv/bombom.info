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
	ErrNoClubFound = errors.New("no club found")
)

type BS struct {
	*resty.Client
}

func New(conf config.Config) *BS {
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
		return club, ErrNoClubFound
	}

	return club, nil
}
