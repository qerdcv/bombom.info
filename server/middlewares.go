package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func rateLimit(rpm int) gin.HandlerFunc {
	rl := new(sync.Map)

	return func(c *gin.Context) {
		l, ok := rl.LoadOrStore(c.ClientIP(), rate.NewLimiter(rate.Every(5*time.Minute), rpm))
		if ok {
			if !l.(*rate.Limiter).Allow() {
				c.JSON(http.StatusTooManyRequests, ErrorResponse{http.StatusText(http.StatusTooManyRequests)})
				c.Abort()

				return
			}
		}
	}
}
