package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p12s/uber-popug/billing/pkg/models"
)

func (h *Handler) createTask(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	_ = accountId // TODO возможно будет использоваться

	var input models.Task
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateTask(input)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
