package v1

import (
	"log/slog"
	"net/http"

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
}

// @Summary      Ping endpoint
// @Description  Responds with a pong payload
// @Tags         Ping
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/ping [get]
func (pc *PingController) pingHandler(c *gin.Context) {
	pc.logger.Debug("Ping endpoint hit", slog.String("endpoint", "/api/v1/ping"))
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"version": "v1",
		"domain":  c.Request.Host,
		"ip":      c.ClientIP(),
	})
}
