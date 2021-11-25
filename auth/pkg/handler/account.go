package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p12s/uber-popug/auth/pkg/models"
)

func (h *Handler) updateAccount(c *gin.Context) {
	var input models.UpdateAccountInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.UpdateAccount(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO для простоты любое обновление account будет BE, хотя BE должно быть только обновление роли
	go h.broker.Event(models.EVENT_ACCOUNT_UPDATED, h.broker.TopicAccountBE, input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func (h *Handler) deleteAccount(c *gin.Context) {
	// TODO надо проверять что это 1) либо свой акк - свой могу удалить всегда 2) либо есть роль админа
	// иначе - нет прав
	var input models.DeleteAccountInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.DeleteAccountByPublicId(input.PublicId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	go h.broker.Event(models.EVENT_ACCOUNT_DELETED, h.broker.TopicAccountCUD, input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}
