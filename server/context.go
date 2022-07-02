package server

import "github.com/qerdcv/bombom.info/domain"

type Context struct {
	Status string
	Club   domain.Club
}
