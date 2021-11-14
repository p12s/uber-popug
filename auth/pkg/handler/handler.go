package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/uber-popug/auth/pkg/broker"
	"github.com/p12s/uber-popug/auth/pkg/service"
)

// Handler - struct contains service
type Handler struct {
	services *service.Service
	broker   *broker.Kafka
}

// NewHandler - constructor
func NewHandler(services *service.Service, broker *broker.Kafka) *Handler {
	return &Handler{services: services, broker: broker}
}

// InitRoutes - routes
func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/health", h.health)
	router.POST("/sign-up", h.signUp)
	router.POST("/sign-in", h.signIn)

	api := router.Group("/", h.userIdentity)
	{
		api.POST("/token", h.token)
	}

	return router
}
