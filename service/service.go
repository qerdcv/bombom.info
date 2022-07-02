package service

import "github.com/qerdcv/bombom.info/client"

type Service struct {
	client *client.BS
}

func New(client *client.BS) *Service {
	return &Service{
		client,
	}
}
