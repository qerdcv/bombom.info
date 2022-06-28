package server

import "github.com/gin-gonic/gin"

type Server struct {
	*gin.Engine
}

func New() *Server {
	s := &Server{
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
