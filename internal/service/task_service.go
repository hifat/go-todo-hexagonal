package service

import (
	"github.com/hifat/go-todo-hexagonal/helper/errs"
	"github.com/hifat/go-todo-hexagonal/helper/zlog"
	"github.com/hifat/go-todo-hexagonal/internal/repository"
)

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{taskRepo}
}

func (t taskService) Get(userID string) ([]Task, error) {
	getTasks, err := t.taskRepo.Get(userID)
	if err != nil {
		zlog.Error(err)
		return nil, errs.HttpError(err)
	}

	tasks := []Task{}
	for _, getTask := range getTasks {
		tasks = append(tasks, Task{
			ID:        getTask.ID,
			UserID:    getTask.UserID,
			Detail:    getTask.Detail,
			Done:      getTask.Done,
			CreatedAt: getTask.CreatedAt,
			UpdatedAt: getTask.UpdatedAt,
		})
	}

	return tasks, nil
}

func (t taskService) Create(task NewTask) (*Task, error) {
	errValidate := validateForm(task)
	if errValidate != nil {
		return nil, errValidate
	}

	newTask := repository.NewTask{
		UserID: task.UserID,
		Detail: task.Detail,
		Done:   task.Done,
	}

	createdTask, err := t.taskRepo.Create(newTask)
	if err != nil {
		zlog.Error(err)
		return nil, errs.HttpError(err)
	}

	taskResponse := Task{
		ID:        createdTask.ID,
		UserID:    createdTask.UserID,
		Detail:    createdTask.Detail,
		Done:      createdTask.Done,
		CreatedAt: createdTask.CreatedAt,
		UpdatedAt: createdTask.UpdatedAt,
	}

	return &taskResponse, nil
}

func (t taskService) Show(id string) (*Task, error) {
	getTask, err := t.taskRepo.Show(id)
	if err != nil {
		zlog.Error(err)
		return nil, errs.HttpError(err)
	}

	taskResponse := Task{
		ID:        getTask.ID,
		UserID:    getTask.UserID,
		Detail:    getTask.Detail,
		Done:      getTask.Done,
		CreatedAt: getTask.CreatedAt,
		UpdatedAt: getTask.UpdatedAt,
	}

	return &taskResponse, nil
}

func (t taskService) ToggleDone(id string) (*Task, error) {
	updatedTask, err := t.taskRepo.ToggleDone(id)
	if err != nil {
		zlog.Error(err)
		return nil, errs.HttpError(err)
	}

	taskResponse := Task{
		ID:        updatedTask.ID,
		UserID:    updatedTask.UserID,
		Detail:    updatedTask.Detail,
		Done:      updatedTask.Done,
		CreatedAt: updatedTask.CreatedAt,
		UpdatedAt: updatedTask.UpdatedAt,
	}

	return &taskResponse, nil
}

func (t taskService) Update(id string, task EditTask) (*Task, error) {
	errValidate := validateForm(task)
	if errValidate != nil {
		return nil, errValidate
	}

	editTask := repository.EditTask{
		Detail: task.Detail,
		Done:   task.Done,
	}

	updatedTask, err := t.taskRepo.Update(id, editTask)

	if err != nil {
		zlog.Error(err)
		return nil, errs.HttpError(err)
	}

	taskResponse := Task{
		ID:        updatedTask.ID,
		UserID:    updatedTask.UserID,
		Detail:    updatedTask.Detail,
		Done:      updatedTask.Done,
		CreatedAt: updatedTask.CreatedAt,
		UpdatedAt: updatedTask.UpdatedAt,
	}

	return &taskResponse, nil
}

func (t taskService) Delete(id string) error {
	err := t.taskRepo.Delete(id)
	if err != nil {
		zlog.Error(err)
		return errs.HttpError(err)
	}

	return nil
}
