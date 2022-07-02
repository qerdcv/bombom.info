package server

import (
	"html/template"

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

	s.Use(rateLimit(s.conf.RateLimit))
	// templ := template.Must(template.New("").ParseFS(templates.Templates, "*.gohtml"))

	s.SetFuncMap(template.FuncMap{
		"inc": func(a int) int {
			return a + 1
		},
	})

	s.LoadHTMLGlob("templates/*.gohtml")
	s.Static("/assets", "assets")

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	s.GET("", s.index)
	s.POST("/join", s.requestToJoin)
}
