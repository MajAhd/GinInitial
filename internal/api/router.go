package api

import (
	"log/slog"

	"gininitial/internal/api/graphql"
	v1 "gininitial/internal/api/rest/v1"
	"gininitial/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter strictly forms the route structure of our service
func SetupRouter(logger *slog.Logger) *gin.Engine {
	r := gin.New()

	// Global Middleware
	r.Use(gin.Recovery())
	r.Use(middleware.SlogMiddleware(logger))

	// Base API grouping
	api := r.Group("/api")

	// --- REST API v1 ---
	restV1 := api.Group("/v1")
	{
		pingController := v1.NewPingController(logger)
		pingController.RegisterRoutes(restV1)
		// Add more v1 controllers here
	}

	// --- REST API v2 (Placeholder for future versioning) ---
	// restV2 := api.Group("/v2")
	// {
	//    Add v2 controllers here to avoid breaking v1 contracts
	// }

	// --- GraphQL API ---
	gqlGroup := api.Group("/graphql")
	{
		graphql.RegisterRoutes(gqlGroup, logger)
	}

	return r
}
