package service

import (
	"errors"
	"fmt"

	"github.com/qerdcv/bombom.info/client"
	"github.com/qerdcv/bombom.info/domain"
)

var (
	ErrClubNotFound = errors.New("club not found")
)

func (s *Service) GetClubByTag(tag string) (club domain.Club, err error) {
	club, err = s.client.GetClubByTag(tag)
	if err != nil {
		if errors.Is(err, client.ErrNotFound) {
			return club, ErrClubNotFound
		}

		return club, fmt.Errorf("client get club by tag: %w", err)
	}

	return club, nil
}
