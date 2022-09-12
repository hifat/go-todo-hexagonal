package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/configs"
	"github.com/hifat/go-todo-hexagonal/helper/zlog"
	"github.com/hifat/go-todo-hexagonal/internal/handler/ginhandler"
	"github.com/hifat/go-todo-hexagonal/internal/handler/ginhandler/ginmiddleware"
	"github.com/hifat/go-todo-hexagonal/internal/repository/gormrepo"
	"github.com/hifat/go-todo-hexagonal/internal/service"
	"github.com/hifat/go-todo-hexagonal/internal/token"
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

	jwtMaker, err := token.NewJWTMaker(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		panic(err)
	}

	middlewareAuth := ginmiddleware.Auth(jwtMaker)

	routeApi := r.Group("/api")

	newTaskGorm := gormrepo.NewTaskGorm(db)
	newTaskSrv := service.NewTaskService(newTaskGorm)
	newTaskHandler := ginhandler.NewTaskHandler(newTaskSrv)

	routeTasks := routeApi.Group("/tasks", middlewareAuth)
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
		routeAuths.POST("/login", newAuthHandler.Login)
		routeAuths.GET("/me", middlewareAuth, newAuthHandler.Me)
	}

	zlog.Info("Server listening on port " + os.Getenv("APP_PORT"))
	r.Run(fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
}
