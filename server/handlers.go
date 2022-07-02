package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/qerdcv/bombom.info/service"
)

func (s *Server) index(c *gin.Context) {
	club, err := s.service.GetClubByTag(s.conf.ClubTag)
	if err != nil {
		if errors.Is(err, service.ErrClubNotFound) {
			c.HTML(http.StatusNotFound, "notFound", gin.H{})

			return
		}

		_ = c.Error(err)

		return
	}

	c.HTML(http.StatusOK, "index", club)
}
