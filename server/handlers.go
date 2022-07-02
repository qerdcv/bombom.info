package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qerdcv/bombom.info/domain"
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

	c.HTML(http.StatusOK, "index", Context{
		Status: c.Query("status"),
		Club:   club,
	})
}

func (s *Server) requestToJoin(c *gin.Context) {
	req := domain.JoinRequest{
		UserTag:      c.PostForm("tag"),
		TelegramName: c.PostForm("telegram"),
	}
	if err := s.service.RequestToJoin(req); err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidJoinRequest):
			c.HTML(http.StatusBadRequest, "index", err.Error())
		case errors.Is(err, service.ErrPlayerNotFound):
			c.HTML(http.StatusBadRequest, "error", fmt.Sprintf("%s: %s", err.Error(), req.UserTag))
		default:
			c.HTML(http.StatusInternalServerError, "error", http.StatusText(http.StatusInternalServerError))
		}

		return
	}

	c.Redirect(http.StatusFound, "?status=success")
}
