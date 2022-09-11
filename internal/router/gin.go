package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/configs"
	"github.com/hifat/go-todo-hexagonal/helper/zlog"
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

func ExecGinRouter() {
	db := configs.GormDB()
	r := gin.Default()

	routeApi := r.Group("/api")

	newTaskGorm := gormrepo.NewTaskGorm(db)
	newTaskSrv := service.NewTaskService(newTaskGorm)
	newTaskHandler := ginhandler.NewTaskHandler(newTaskSrv)

	routeTasks := routeApi.Group("/tasks")
	{
		routeTasks.GET("/", newTaskHandler.Get)
		routeTasks.POST("/", newTaskHandler.Create)
		routeTasks.GET("/:task", newTaskHandler.Show)
		routeTasks.PUT("/:task", newTaskHandler.Update)
		routeTasks.DELETE("/:task", newTaskHandler.Delete)
	}

	newAuthGorm := gormrepo.NewAuthGorm(db)
	newAuthSrv := service.NewAuthService(newAuthGorm)
	newAuthHandler := ginhandler.NewAuthHandler(newAuthSrv)

	routeAuths := routeApi.Group("/auth")
	{
		routeAuths.POST("/register", newAuthHandler.Register)
	}

	zlog.Info("Server listening on port " + os.Getenv("APP_PORT"))
	r.Run(fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
}
