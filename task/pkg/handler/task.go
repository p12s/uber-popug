package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/p12s/uber-popug/task/internal/tools"
	"github.com/p12s/uber-popug/task/pkg/models"
)

func (h *Handler) getTask(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	_ = accountId // TODO возможно будут проверки, принадлежит ли этот таск текущему юзеру

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	task, err := h.services.GetTaskById(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) createTask(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var input models.Task
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// TODO достаю по account_id (int) - костыль. Надо по public_id (uuid.UUID)
	account, err := h.services.GetAccountByPrimaryId(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	// пусть пока таск ассайнится на текущего пользователя (чей токен инициировал создание)
	input.AssignedAccountId = account.PublicId

	input.Description = tools.GetPureTitle(input.Description)
	input.JiraId = tools.GetPureTaskKey(input.Description)
	input.PublicId = uuid.New()
	input.Status = models.TASK_BIRD_IN_CAGE

	id, err := h.services.CreateTask(input)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	// CUD-событие на создание таска
	go h.broker.Event(models.EVENT_TASK_CREATED, h.broker.TopicTaskCUD,
		&models.Task{
			PublicId:    input.PublicId,
			Description: input.Description,
			Status:      input.Status,
		})

	// BE-событие на ассайн таска
	go h.broker.Event(models.EVENT_TASK_BIRD_CAGED, h.broker.TopicTaskBE,
		&models.Task{
			PublicId:          input.PublicId,
			AssignedAccountId: account.PublicId,
		})

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) birdCageTask(c *gin.Context) {
	var input models.BirdCageTask
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.BirdCageTask(input.PublicId, input.AccountId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	go h.broker.Event(models.EVENT_TASK_BIRD_CAGED, h.broker.TopicTaskBE,
		&input)

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) milletBowlTask(c *gin.Context) {
	var input models.MilletBowlTask
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.MilletBowlTask(input.PublicId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	go h.broker.Event(models.EVENT_TASK_MILLET_BOWLED, h.broker.TopicTaskBE,
		&input)

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) getAllTask(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	_ = accountId // TODO возможно будут проверки, принадлежит ли этот таск текущему юзеру
	// employee - отдавать только свои таски,
	// manager, accountant, ... - другие правила

	tasks, err := h.services.GetAllTasksByAssignedAccountId(accountId)
	if err != nil {
		fmt.Println("error - task handler", err.Error())

		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	fmt.Println("NOT error - OK task handler", tasks)

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) birdCageAllTasks(c *gin.Context) {
	// достаем все таски
	// достаем всех работников
	// переассайниваем рандомно, в цикле отправляя события
	fmt.Println("resaasigne all tasks")
	c.JSON(http.StatusOK, nil)
}
