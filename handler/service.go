package handler

import (
	"context"
	"todo/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService
type ListTasksService interface {
	ListTasks(context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(context.Context, string) (*entity.Task, error)
}
