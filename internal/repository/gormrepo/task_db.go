package gormrepo

import (
	"strconv"
	"time"

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

type createTask struct {
	ID        uint
	UserID    uint
	Detail    string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *taskRepositoryDB) Create(newTask repository.NewTask) (*repository.Task, error) {
	userID, err := strconv.Atoi(newTask.UserID)
	if err != nil {
		return nil, err
	}
	createdTask := createTask{
		UserID: uint(userID),
		Detail: newTask.Detail,
		Done:   newTask.Done,
	}

	tx := r.db.Create(&createdTask)
	if tx.Error != nil {
		return nil, tx.Error
	}

	task := repository.Task{
		ID:        strconv.Itoa(int(createdTask.ID)),
		UserID:    strconv.Itoa(int(createdTask.UserID)),
		Detail:    createdTask.Detail,
		Done:      createdTask.Done,
		CreatedAt: createdTask.CreatedAt,
		UpdatedAt: createdTask.UpdatedAt,
	}

	return &task, nil
}

func (r *taskRepositoryDB) Show(id string) (*repository.Task, error) {
	var task repository.Task
	tx := r.db.Find(&task, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &task, nil
}

func (r *taskRepositoryDB) Update(id string, editTask repository.EditTask) (*repository.Task, error) {
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

func (r *taskRepositoryDB) Delete(id string) error {
	tx := r.db.Delete(&repository.Task{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
