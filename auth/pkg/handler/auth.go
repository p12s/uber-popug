package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/p12s/uber-popug/auth/pkg/models"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.Account

	input.PublicId = uuid.New()
	input.Role = models.ROLE_EMPLOYEE
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.CreateAccount(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.Password = ""
	go h.broker.Event(models.EVENT_ACCOUNT_CREATED, h.broker.TopicAccountCUD, input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// TODO здесь и в других сервисах в токене юзера храню account_id (int) -
// первичный ключ из БД сервиса auth
// по-хорошему надо хранить public_id (uuid.UUID) - он одинаковый во всех сервисах
// в то время как account_id (int) может отличаться
func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accountToken, err := h.services.Authorizer.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	go h.broker.Event(models.EVENT_ACCOUNT_TOKEN_UPDATED, h.broker.TopicAccountCUD, accountToken)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": accountToken.Token,
	})
}

func (h *Handler) token(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	account, err := h.services.GetAccountById(accountId)
	// TODO если "no rows in result set" - возвращать осмысленный текст
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}
