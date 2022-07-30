package service

import (
	"context"
	"todo/entity"
	"todo/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister
type TaskAdder interface {
	AddTask(context.Context, store.Execer, *entity.Task) error
}

type TaskLister interface {
	ListTasks(context.Context, store.Queryer) (entity.Tasks, error)
}
