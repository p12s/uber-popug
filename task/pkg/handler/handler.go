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
	router.Use(CORSMiddleware())

	router.GET("/health", h.health)

	task := router.Group("/task", h.userIdentity)
	{
		task.GET("/:id", h.getTask)
		task.GET("/", h.getAllTask)
		task.POST("/", h.createTask)
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
