package middlewares

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LogsMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next()

	buf := new(strings.Builder)
	for _, err := range c.Errors {
		buf.Write([]byte(err.Error()))
		buf.WriteRune('\n')
	}

	log.Infof(
		"%v | Status code: %d | Method: %s | Path: %s | Response time: %s | Errors: %s",
		time.Now().Format("02-01-2006 15:04:05"),
		c.Writer.Status(),
		c.Request.Method,
		c.Request.RequestURI,
		time.Since(start),
		buf.String(),
	)
}
