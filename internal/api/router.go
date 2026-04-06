package api

import (
	"log/slog"

	_ "gininitial/docs"
	"gininitial/internal/api/graphql"
	liveness "gininitial/internal/api/rest/liveness"
	v1 "gininitial/internal/api/rest/v1"
	ws "gininitial/internal/api/ws"
	"gininitial/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uptrace/bun"
)

// RouterDependencies contains all shared app dependencies like logging, DB connections, etc.
type RouterDependencies struct {
	Logger *slog.Logger
	DB     *bun.DB
	// Config *Config    (example for the future)
}

// SetupRouter strictly forms the route structure of our service
func SetupRouter(deps RouterDependencies) *gin.Engine {
	r := gin.New()

	// Global Middleware
	r.Use(gin.Recovery())
	r.Use(middleware.SlogMiddleware(deps.Logger))
	r.Use(middleware.ErrorHandler())
	r.TrustedPlatform = "X-CDN-Client-IP"

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// WebSocket endpoint
	r.GET("/ws", ws.HandleWebSocket)

	// Break down the routing logic into cleanly decoupled modules
	setupHealthRoutes(r, deps)
	setupAPIRoutes(r, deps)
	setupGraphQLRoutes(r, deps)

	return r
}

func setupHealthRoutes(r *gin.Engine, deps RouterDependencies) {
	health := r.Group("/health")
	{
		livenessController := liveness.NewLivenessController(deps.Logger)
		livenessController.RegisterRoutes(health)
	}
}

func setupAPIRoutes(r *gin.Engine, deps RouterDependencies) {
	api := r.Group("/api")

	// --- REST API v1 ---
	restV1 := api.Group("/v1")
	{
		pingController := v1.NewPingController(deps.Logger)
		pingController.RegisterRoutes(restV1)
		// Add more v1 controllers here dynamically
	}

	// --- REST API v2 (Placeholder for future versioning) ---
	// restV2 := api.Group("/v2")
	// { ... }
}

func setupGraphQLRoutes(r *gin.Engine, deps RouterDependencies) {
	gqlGroup := r.Group("/api/graphql")
	{
		graphql.RegisterRoutes(gqlGroup, deps.Logger)
	}
}
