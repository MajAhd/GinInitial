package v1

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// PingController holds dependencies for ping handler
type PingController struct {
	logger *slog.Logger
}

// NewPingController creates a new instance of PingController
func NewPingController(logger *slog.Logger) *PingController {
	return &PingController{
		logger: logger,
	}
}

// RegisterRoutes registers endpoints to a router group
func (pc *PingController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/ping", pc.pingHandler)
	rg.GET("/liveness", pc.livenessHandler)
}

func (pc *PingController) pingHandler(c *gin.Context) {
	pc.logger.Debug("Ping endpoint hit", slog.String("endpoint", "/api/v1/ping"))
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"version": "v1", // Explicitly define version returned
		"domain":  c.Request.Host,
		"ip":      c.ClientIP(),
	})
}

func (pc *PingController) livenessHandler(c *gin.Context) {
	pc.logger.Debug("Liveness endpoint hit", slog.String("endpoint", "/api/v1/liveness"))
	c.JSON(http.StatusOK, gin.H{
		"message": "live",
		"version": "v1",
		"time":    time.Now(),
	})
}
