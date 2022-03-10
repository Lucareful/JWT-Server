package middlewares

import (
	"bytes"
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)
		log.Debug(string(body))
		log.Debug(c.Request.Header)
		c.Next()
	}
}
