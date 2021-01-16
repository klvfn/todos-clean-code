package service

import (
	"context"
	"time"

	"github.com/klvfn/todos-clean-code/pkg/todo/entity"
	"github.com/klvfn/todos-clean-code/pkg/todo/repository/mysql"
)

// TodoService contract for todo service
type TodoService interface {
	GetAll(ctx context.Context) ([]entity.Todo, error)
	GetByID(ctx context.Context, id int64) (entity.Todo, error)
	Create(ctx context.Context, todo entity.Todo) (int64, error)
	Update(ctx context.Context, id int64, todo entity.Todo) error
	Delete(ctx context.Context, id int64) error
}

type todoService struct {
	MysqlTodoRepo mysql.TodoRepository
	CtxTimeout    time.Duration
}

// NewTodoService create an instance of todos service
func NewTodoService(mysqlTodoRepo mysql.TodoRepository, ctxTimeout time.Duration) TodoService {
	return &todoService{
		MysqlTodoRepo: mysqlTodoRepo,
		CtxTimeout:    ctxTimeout,
	}
}

func (s todoService) GetAll(ctx context.Context) ([]entity.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()
	todos, err := s.MysqlTodoRepo.GetAll()
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (s todoService) GetByID(ctx context.Context, id int64) (entity.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()
	todo, err := s.MysqlTodoRepo.GetByID(id)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s todoService) Create(ctx context.Context, todo entity.Todo) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()
	id, err := s.MysqlTodoRepo.Create(todo)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s todoService) Update(ctx context.Context, id int64, todo entity.Todo) error {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()
	err := s.MysqlTodoRepo.Update(id, todo)
	if err != nil {
		return err
	}
	return nil
}

func (s todoService) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()
	err := s.MysqlTodoRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
