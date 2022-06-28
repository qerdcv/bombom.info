package server

import (
	"github.com/gin-gonic/gin"
	"github.com/qerdcv/bombom.info/config"
)

type Server struct {
	conf config.Config

	*gin.Engine
}

func New(conf config.Config) *Server {
	if conf.IsProdEnv() {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &Server{
		conf,
		gin.Default(),
	}

	//templ := template.Must(template.New("").ParseFS(templates.Templates, "*.gohtml"))

	s.LoadHTMLGlob("templates/*.gohtml")
	s.Static("/assets", "assets")

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	s.GET("", s.index)
}
