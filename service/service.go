package service

import "github.com/klvfn/todos-clean-code/repository"

type Service struct {
	Todo TodoService
}

func NewService(dao *repository.Dao) *Service {
	todo := NewTodoService(dao)

	return &Service{
		Todo: todo,
	}
}
