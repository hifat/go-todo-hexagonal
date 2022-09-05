package ginhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/internal/handler"
	"github.com/hifat/go-todo-hexagonal/internal/service"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	newTask := service.NewTask{
		UserID: 1,
		Detail: reqTask.Detail,
		Done:   false,
	}

	task, err := t.taskSrv.Create(newTask)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (t taskHandler) Show(ctx *gin.Context) {
	taskID, err := strconv.ParseUint(ctx.Param("task"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	task, err := t.taskSrv.Show(uint(taskID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (t taskHandler) Update(ctx *gin.Context) {
	taskID, err := strconv.ParseUint(ctx.Param("task"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	var taskReq handler.EditTask
	err = ctx.ShouldBind(&taskReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	editTask := service.EditTask{
		Detail: taskReq.Detail,
		Done:   taskReq.Done,
	}

	task, err := t.taskSrv.Update(uint(taskID), editTask)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (t taskHandler) Delete(ctx *gin.Context) {
	taskID, err := strconv.ParseUint(ctx.Param("task"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	err = t.taskSrv.Delete(uint(taskID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}
