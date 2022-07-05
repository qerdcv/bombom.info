package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qerdcv/bombom.info/domain"
	"github.com/qerdcv/bombom.info/service"
)

func (s *Server) getClubInfo(c *gin.Context) {
	club, err := s.service.GetClubByTag(s.conf.ClubTag)
	if err != nil {
		if errors.Is(err, service.ErrClubNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{err.Error()})

			return
		}

		c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})

		return
	}

	c.JSON(http.StatusOK, club)
}

func (s *Server) requestToJoin(c *gin.Context) {
	var req domain.JoinRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})

		return
	}

	if err := s.service.RequestToJoin(req); err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidJoinRequest):
			c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		case errors.Is(err, service.ErrPlayerNotFound):
			c.JSON(
				http.StatusBadRequest, ErrorResponse{fmt.Sprintf("%s: %s", err.Error(), req.UserTag)},
			)
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
		}

		return
	}

	c.JSON(http.StatusOK, SuccessResponse{"OK"})
}
