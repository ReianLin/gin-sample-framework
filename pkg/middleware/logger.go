package middleware

import (
	"gin-sample-framework/pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := getRelativePath(c.Request.URL.Path)
		query := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		log.WithContext(c.Request.Context()).WithFields(logger.Fields{
			"path":       path,
			"query":      query,
			"ip":         c.ClientIP(),
			"user-agent": c.Request.UserAgent(),
			"latency":    latency,
			"size":       c.Writer.Size(),
			"error":      c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}).Info("gin request")
		c.Next()
	}
}

func getRelativePath(fullPath string) string {
	path := strings.TrimPrefix(fullPath, "/")
	if path == "" {
		return "/"
	}
	return path
}
