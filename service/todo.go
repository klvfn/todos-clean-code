package service

import (
	"github.com/klvfn/todos-clean-code/entity"
	"github.com/klvfn/todos-clean-code/repository"
)

// TodoService contract for todo service
type TodoService interface {
	GetAll() ([]entity.Todo, error)
	GetByID(id int64) (entity.Todo, error)
	Create(todo entity.Todo) (int64, error)
	Update(id int64, todo entity.Todo) error
	Delete(id int64) error
}

type todoService struct {
	dao *repository.Dao
}

// NewTodoService create an instance of todos service
func NewTodoService(dao *repository.Dao) TodoService {
	return &todoService{
		dao: dao,
	}
}

func (s todoService) GetAll() ([]entity.Todo, error) {
	todos, err := s.dao.MysqlRepo.Todo.GetAll()
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (s todoService) GetByID(id int64) (entity.Todo, error) {
	todo, err := s.dao.MysqlRepo.Todo.GetByID(id)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s todoService) Create(todo entity.Todo) (int64, error) {
	id, err := s.dao.MysqlRepo.Todo.Create(todo)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s todoService) Update(id int64, todo entity.Todo) error {
	err := s.dao.MysqlRepo.Todo.Update(id, todo)
	if err != nil {
		return err
	}
	return nil
}

func (s todoService) Delete(id int64) error {
	err := s.dao.MysqlRepo.Todo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
