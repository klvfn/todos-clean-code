package redis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/klvfn/todos-clean-code/helper"
	"github.com/klvfn/todos-clean-code/pkg/todo/entity"
)

// TodoRepository contract for redis todo repository
type TodoRepository interface {
	GetAll(ctx context.Context) ([]entity.Todo, error)
	GetByID(ctx context.Context, id string) (entity.Todo, error)
	Create(ctx context.Context, todo entity.Todo) (string, error)
	Update(ctx context.Context, id string, todo entity.Todo) error
	Delete(ctx context.Context, id string) error
}

type todoRepository struct {
	RedisDB *redis.Client
}

// NewTodoRepository create new instance of redis todo repository
func NewTodoRepository(redisDB *redis.Client) TodoRepository {
	return &todoRepository{
		RedisDB: redisDB,
	}
}

func (tr todoRepository) GetAll(ctx context.Context) ([]entity.Todo, error) {
	var cursor uint64
	todos := []entity.Todo{}
	todoKeys := make([]string, 0)
	for {
		keys, cursor, err := tr.RedisDB.Scan(ctx, cursor, "*", 10).Result()
		if err != nil {
			return todos, err
		}
		todoKeys = append(todoKeys, keys...)
		if cursor == 0 {
			break
		}
	}

	items, err := tr.RedisDB.MGet(ctx, todoKeys...).Result()
	if err != nil {
		return todos, err
	}

	for _, i := range items {
		todo := entity.Todo{}
		err = json.Unmarshal([]byte(i.(string)), &todo)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (tr todoRepository) GetByID(ctx context.Context, id string) (entity.Todo, error) {
	todo := entity.Todo{}
	val, err := tr.RedisDB.Get(ctx, id).Result()
	if err != nil {
		return todo, err
	}
	err = json.Unmarshal([]byte(val), &todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (tr todoRepository) Create(ctx context.Context, todo entity.Todo) (string, error) {
	uid := helper.GenerateUniqueID()
	// Set todo ID
	todo.ID = uid
	todoByte, err := json.Marshal(todo)
	if err != nil {
		return uid, err
	}

	// Save to redis
	err = tr.RedisDB.Set(ctx, todo.ID, string(todoByte), 0).Err()
	if err != nil {
		return uid, err
	}
	return uid, nil
}

func (tr todoRepository) Update(ctx context.Context, id string, todo entity.Todo) error {
	todoByte, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	_, err = tr.RedisDB.GetSet(ctx, id, string(todoByte)).Result()
	if err != nil {
		return err
	}

	return nil
}

func (tr todoRepository) Delete(ctx context.Context, id string) error {
	_, err := tr.RedisDB.Del(ctx, id).Result()
	if err != nil {
		return err
	}
	return nil
}
