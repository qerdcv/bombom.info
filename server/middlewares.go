package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func rateLimit(rpm int) gin.HandlerFunc {
	rl := new(sync.Map)

	return func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			log.Println("not post")

			c.Next()

			return
		}

		fmt.Println(c.ClientIP(), " - client ip address")
		l, ok := rl.LoadOrStore(c.ClientIP(), rate.NewLimiter(rate.Every(time.Minute), rpm))
		if ok {
			limiter := l.(*rate.Limiter)
			if limiter.Allow() {
				//c.Next()
				fmt.Println("HEHE")
				return
			}

			c.HTML(http.StatusTooManyRequests, "error", http.StatusText(http.StatusTooManyRequests))
			return
		}
	}
}
