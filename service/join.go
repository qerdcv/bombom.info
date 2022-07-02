package service

import (
	"errors"
	"fmt"

	"github.com/qerdcv/bombom.info/client"
	"github.com/qerdcv/bombom.info/domain"
)

var (
	ErrInvalidJoinRequest = errors.New("invalid join request")
	ErrPlayerNotFound     = errors.New("player not found")
)

func (s *Service) RequestToJoin(req domain.JoinRequest) error {
	if err := req.Validate(); err != nil {
		return ErrInvalidJoinRequest
	}

	player, err := s.client.GetPlayerByTag(req.UserTag)
	if err != nil {
		if errors.Is(err, client.ErrNotFound) {
			return ErrPlayerNotFound
		}

		return fmt.Errorf("client get player by tag: %w", err)
	}

	return s.bot.SendJoinRequest(player)
}
