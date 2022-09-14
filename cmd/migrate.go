package main

import (
	"github.com/hifat/go-todo-hexagonal/configs"
	"github.com/hifat/go-todo-hexagonal/migration"
)

func main() {
	db := configs.GormDB()
	db.AutoMigrate(
		&migration.User{},
		&migration.Task{},
		&migration.Session{},
	)
}
