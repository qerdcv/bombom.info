package server

import (
	"net/http"

	"github.com/qerdcv/bombom.info/service"

	"github.com/gin-gonic/gin"
	"github.com/qerdcv/bombom.info/config"
)

type Server struct {
	conf config.Config

	service *service.Service
	*gin.Engine
}

func New(conf config.Config, service *service.Service) *Server {
	if conf.IsProdEnv() {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &Server{
		conf,
		service,
		gin.Default(),
	}

	s.NoRoute(func(c *gin.Context) { c.JSON(http.StatusNotFound, ErrorResponse{Message: "not found"}) })
	s.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, ErrorResponse{Message: "method not allowed"})
	})

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	r := s.Group("/api/v1")
	{
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, SuccessResponse{Message: "pong"})
		})
		r.GET("/clubs", s.getClubInfo)
		r.Use(rateLimit(s.conf.RateLimit)).POST("/join", s.requestToJoin)
	}
}
