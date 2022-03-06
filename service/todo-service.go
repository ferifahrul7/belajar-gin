package service

import "belajar-gin/entity"

type TodoService interface {
	Save(entity.Todo) entity.Todo
	FindAll() []entity.Todo
}

type todoService struct {
	todos []entity.Todo
}

func New() TodoService {
	return &todoService{}
}

func (service *todoService) Save(todo entity.Todo) entity.Todo {
	service.todos = append(service.todos, todo)
	return todo
}
func (service *todoService) FindAll() []entity.Todo {
	return service.todos
}
