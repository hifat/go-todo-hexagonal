package main

import (
	"time"

	"github.com/hifat/go-todo-hexagonal/internal/router"
)

func main() {
	initTimezone()
	router.ExecGinRouter()
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
