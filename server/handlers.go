package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{})
}
