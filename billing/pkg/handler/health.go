package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) health(c *gin.Context) { // TODO - доавить мидлвер - вывод в консоль время и тп - как в nginx логирование
	c.JSON(http.StatusOK, map[string]interface{}{
		"service": "billing",
		"status":  "OK",
	})
}
