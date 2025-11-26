package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		log.Println("Executing Middleware for : ", c.Request.URL.Path)

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		log.Printf("[%d] %s %s (%s)\n", status, c.Request.Method, c.Request.URL.Path, latency)
	}
}
