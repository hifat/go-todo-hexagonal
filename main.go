package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/configs"
	"github.com/hifat/go-todo-hexagonal/internal/handler/ginhandler"
	"github.com/hifat/go-todo-hexagonal/internal/repository/gormrepo"
	"github.com/hifat/go-todo-hexagonal/internal/service"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	db := configs.GormDB()
	r := gin.Default()

	newTaskGorm := gormrepo.NewTaskGorm(db)
	newTaskSrv := service.NewTaskService(newTaskGorm)
	newTaskHandler := ginhandler.NewTaskHandler(newTaskSrv)

	routeApi := r.Group("/api")
	routeTasks := routeApi.Group("/tasks")
	{
		routeTasks.GET("/", newTaskHandler.Get)
		routeTasks.POST("/", newTaskHandler.Create)
		routeTasks.GET("/:task", newTaskHandler.Show)
		routeTasks.PUT("/:task", newTaskHandler.Update)
		routeTasks.DELETE("/:task", newTaskHandler.Delete)
	}

	r.Run(fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_POST")))
}
