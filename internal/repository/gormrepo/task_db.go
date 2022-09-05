package gormrepo

import (
	"github.com/hifat/go-todo-hexagonal/internal/repository"
	"gorm.io/gorm"
)

type taskRepositoryDB struct {
	db *gorm.DB
}

func NewTaskGorm(db *gorm.DB) repository.TaskRepository {
	return &taskRepositoryDB{db}
}

func (r *taskRepositoryDB) Get() ([]repository.Task, error) {
	var tasks []repository.Task
	tx := r.db.Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tasks, nil
}

func (r *taskRepositoryDB) Create(newTask repository.NewTask) (*repository.Task, error) {
	task := repository.Task{
		UserID: newTask.UserID,
		Detail: newTask.Detail,
		Done:   newTask.Done,
	}

	tx := r.db.Create(&task)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &task, nil
}

func (r *taskRepositoryDB) Show(id uint) (*repository.Task, error) {
	var task repository.Task
	tx := r.db.Find(&task, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &task, nil
}

func (r *taskRepositoryDB) Update(id uint, editTask repository.EditTask) (*repository.Task, error) {
	var task repository.Task

	tx := r.db.First(&task, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = r.db.Model(&task).Updates(map[string]interface{}{
		"detail": editTask.Detail,
		"Done":   editTask.Done,
	})
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &task, nil
}

func (r *taskRepositoryDB) Delete(id uint) error {
	tx := r.db.Delete(&repository.Task{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
