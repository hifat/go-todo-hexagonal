package ginhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/internal/handler"
	"github.com/hifat/go-todo-hexagonal/internal/service"
	"github.com/hifat/go-todo-hexagonal/internal/token"
)

type taskHandler struct {
	taskSrv service.TaskService
}

func NewTaskHandler(taskSrv service.TaskService) taskHandler {
	return taskHandler{taskSrv}
}

func (t taskHandler) Get(ctx *gin.Context) {
	tasks, err := t.taskSrv.Get()
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func (t taskHandler) Create(ctx *gin.Context) {
	var reqTask handler.NewTask
	err := ctx.ShouldBind(&reqTask)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	user, _ := ctx.MustGet("user").(*token.Payload)

	newTask := service.NewTask{
		UserID: user.UserID,
		Detail: reqTask.Detail,
	}

	task, err := t.taskSrv.Create(newTask)
	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (t taskHandler) Show(ctx *gin.Context) {
	taskID := ctx.Param("task")

	task, err := t.taskSrv.Show(taskID)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (t taskHandler) Update(ctx *gin.Context) {
	taskID := ctx.Param("task")

	var taskReq handler.EditTask
	err := ctx.ShouldBind(&taskReq)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	editTask := service.EditTask{
		Detail: taskReq.Detail,
		Done:   taskReq.Done,
	}

	task, err := t.taskSrv.Update(taskID, editTask)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (t taskHandler) Delete(ctx *gin.Context) {
	taskID := ctx.Param("task")

	err := t.taskSrv.Delete(taskID)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}
