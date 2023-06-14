package middleware

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func ResponseLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{body: new(bytes.Buffer), ResponseWriter: c.Writer}
		c.Writer = w
		if w.body.Len() == 0 {
			logger.Debug("[RESPONSE]",
				zap.Int("status code", c.Writer.Status()),
			)
			c.Next()
			return
		}
		logger.Debug("[RESPONSE]",
			zap.Int("status code", c.Writer.Status()),
		)
		c.Next()
		return
	}
}
