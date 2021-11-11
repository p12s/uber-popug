package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/uber-popug/task/pkg/broker"
	"github.com/p12s/uber-popug/task/pkg/service"
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

	task := router.Group("/task", h.userIdentity)
	{
		task.GET("/:id", h.getTask)
		task.POST("/", h.createTask)
	}

	return router
}
