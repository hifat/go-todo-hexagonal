package gormrepo

import (
	"strconv"
	"time"

	"github.com/hifat/go-todo-hexagonal/helper/errs"
	"github.com/hifat/go-todo-hexagonal/internal/repository"
	"gorm.io/gorm"
)

type taskRepositoryDB struct {
	db *gorm.DB
}

func NewTaskGorm(db *gorm.DB) repository.TaskRepository {
	return &taskRepositoryDB{db}
}

func (r *taskRepositoryDB) Get(userID string) ([]repository.Task, error) {
	var tasks []repository.Task
	tx := r.db.Where("user_id = ?", userID).Find(&tasks)
	if tx.Error != nil {
		return nil, errHandler(tx.Error)
	}

	return tasks, nil
}

type Task struct {
	ID        uint
	UserID    uint
	Detail    string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *taskRepositoryDB) Create(newTask repository.NewTask) (*repository.Task, error) {
	taskID, err := strconv.Atoi(newTask.UserID)
	if err != nil {
		return nil, errs.NaN("TaskID")
	}
	createdTask := Task{
		UserID: uint(taskID),
		Detail: newTask.Detail,
		Done:   newTask.Done,
	}

	tx := r.db.Create(&createdTask)
	if tx.Error != nil {
		return nil, errHandler(tx.Error)
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
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errs.NaN("TaskID")
	}

	var taskReciever Task
	tx := r.db.Find(&taskReciever, taskID)
	if tx.Error != nil {
		return nil, errHandler(tx.Error)
	}

	task := repository.Task{
		ID:        strconv.Itoa(int(taskReciever.ID)),
		UserID:    strconv.Itoa(int(taskReciever.UserID)),
		Detail:    taskReciever.Detail,
		Done:      taskReciever.Done,
		CreatedAt: taskReciever.CreatedAt,
		UpdatedAt: taskReciever.UpdatedAt,
	}

	return &task, nil
}

func (r *taskRepositoryDB) ToggleDone(id string) (*repository.Task, error) {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errs.NaN("taskID")
	}

	var task repository.Task
	tx := r.db.First(&task, taskID)
	if tx.Error != nil {
		return nil, errHandler(tx.Error)
	}

	tx = r.db.Model(&task).Updates(map[string]interface{}{
		"Done": !task.Done,
	})
	if tx.Error != nil {
		return nil, errHandler(tx.Error)
	}

	return &task, nil
}

func (r *taskRepositoryDB) Update(id string, editTask repository.EditTask) (*repository.Task, error) {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errs.NaN("taskID")
	}

	var task repository.Task
	tx := r.db.First(&task, taskID)
	if tx.Error != nil {
		return nil, errHandler(tx.Error)
	}

	tx = r.db.Model(&task).Updates(map[string]interface{}{
		"detail": editTask.Detail,
		"Done":   editTask.Done,
	})
	if tx.Error != nil {
		return nil, errHandler(tx.Error)
	}

	return &task, nil
}

func (r *taskRepositoryDB) Delete(id string) error {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return errs.NaN("TaskID")
	}

	tx := r.db.Delete(&repository.Task{}, taskID)
	if tx.Error != nil {
		return errHandler(tx.Error)
	}

	return nil
}
