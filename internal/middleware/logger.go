package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

// SlogMiddleware logs a Gin HTTP request in JSON format using slog
func SlogMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		msg := "HTTP Request"
		logAttrs := []slog.Attr{
			slog.Int("status", status),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("ip", c.ClientIP()),
			slog.Duration("latency", latency),
			slog.String("user_agent", c.Request.UserAgent()),
		}

		if len(c.Errors) > 0 {
			logAttrs = append(logAttrs, slog.String("errors", c.Errors.String()))
			logger.LogAttrs(c.Request.Context(), slog.LevelError, msg, logAttrs...)
		} else if status >= 400 {
			logger.LogAttrs(c.Request.Context(), slog.LevelWarn, msg, logAttrs...)
		} else {
			logger.LogAttrs(c.Request.Context(), slog.LevelInfo, msg, logAttrs...)
		}
	}
}
