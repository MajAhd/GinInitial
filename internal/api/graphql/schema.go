package graphql

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the GraphQL endpoint placeholder
func RegisterRoutes(rg *gin.RouterGroup, logger *slog.Logger) {
	rg.POST("/", func(c *gin.Context) {
		logger.Debug("GraphQL endpoint hit")
		c.JSON(http.StatusOK, gin.H{
			"message": "graphql endpoint placeholder",
		})
	})

	// You can add graphql playground via GET request as well
	rg.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "GraphQL Playground UI placeholder")
	})
}
