package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qerdcv/bombom.info/client"
)

func (s *Server) index(c *gin.Context) {
	club, err := s.client.GetClubByTag(s.conf.ClubTag)
	if err != nil {
		if errors.Is(err, client.ErrNoClubFound) {
			c.HTML(http.StatusNotFound, "notFound", gin.H{})

			return
		}

		_ = c.Error(err)

		return
	}

	c.HTML(http.StatusOK, "index", club)
}
