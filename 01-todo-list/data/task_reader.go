package data

import "github.com/benfleuty/goprojects/todoapp/model"

type TaskReader interface {
	GetAll() []model.Task
	WriteTask(*string) (model.Task, error)
}
