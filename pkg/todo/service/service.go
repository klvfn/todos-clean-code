package service

import (
	"context"
	"time"

	"github.com/klvfn/todos-clean-code/pkg/todo/entity"
	"github.com/klvfn/todos-clean-code/pkg/todo/repository/redis"
)

// TodoService contract for todo service
type TodoService interface {
	GetAll(ctx context.Context) ([]entity.Todo, error)
	GetByID(ctx context.Context, id string) (entity.Todo, error)
	Create(ctx context.Context, todo entity.Todo) (string, error)
	Update(ctx context.Context, id string, todo entity.Todo) error
	Delete(ctx context.Context, id string) error
}

type todoService struct {
	RedisTodoRepo redis.TodoRepository
	CtxTimeout    time.Duration
}

// NewTodoService create an instance of todos service
func NewTodoService(redisTodoRepo redis.TodoRepository, ctxTimeout time.Duration) TodoService {
	return &todoService{
		RedisTodoRepo: redisTodoRepo,
		CtxTimeout:    ctxTimeout,
	}
}

func (s todoService) GetAll(ctx context.Context) ([]entity.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()
	todos, err := s.RedisTodoRepo.GetAll(ctx)
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (s todoService) GetByID(ctx context.Context, id string) (entity.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()
	todo, err := s.RedisTodoRepo.GetByID(ctx, id)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s todoService) Create(ctx context.Context, todo entity.Todo) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()
	id, err := s.RedisTodoRepo.Create(ctx, todo)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s todoService) Update(ctx context.Context, id string, todo entity.Todo) error {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()
	err := s.RedisTodoRepo.Update(ctx, id, todo)
	if err != nil {
		return err
	}
	return nil
}

func (s todoService) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()
	err := s.RedisTodoRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
