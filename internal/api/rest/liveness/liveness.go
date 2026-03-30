package v1

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LivenessController holds dependencies for ping handler
type LivenessController struct {
	logger *slog.Logger
}

// NewLivenessController creates a new instance of LivenessController
func NewLivenessController(logger *slog.Logger) *LivenessController {
	return &LivenessController{
		logger: logger,
	}
}

// RegisterRoutes registers endpoints to a router group
func (pc *LivenessController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/liveness", pc.livenessHandler)
	rg.GET("/readiness", pc.readinessHandler)
}

// @Summary      Liveness Check
// @Description  Reports the vital status of the app
// @Tags         Health
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /health/liveness [get]
// @Router       /health/liveness [post]
func (pc *LivenessController) livenessHandler(c *gin.Context) {
	pc.logger.Debug("Liveness endpoint hit", slog.String("endpoint", "/health/liveness"))
	c.JSON(http.StatusOK, gin.H{
		"message": "live",
		"method":  c.Request.Method,
		"time":    time.Now(),
	})
}

// @Summary      Readiness Check
// @Description  Reports if the app is ready to serve traffic
// @Tags         Health
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /health/readiness [get]
// @Router       /health/readiness [post]
func (pc *LivenessController) readinessHandler(c *gin.Context) {
	pc.logger.Debug("Readiness endpoint hit", slog.String("endpoint", "/health/readiness"))
	c.JSON(http.StatusOK, gin.H{
		"message": "read",
		"method":  c.Request.Method,
		"time":    time.Now(),
	})
}
