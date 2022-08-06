package mysql

import (
	"github.com/klvfn/todos-clean-code/entity"
	"gorm.io/gorm"
)

// TodoRepository contract for mysql todo repository
type TodoRepository interface {
	GetAll() ([]entity.Todo, error)
	GetByID(id int64) (entity.Todo, error)
	Create(todo entity.Todo) (int64, error)
	Update(id int64, todo entity.Todo) error
	Delete(id int64) error
}

type todoRepository struct {
	db *gorm.DB
}

// NewTodoRepository create new instance of mysql todo repository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r todoRepository) GetAll() ([]entity.Todo, error) {
	todos := []entity.Todo{}
	res := r.db.Find(&todos)
	if res.Error != nil {
		return todos, res.Error
	}
	return todos, nil
}

func (r todoRepository) GetByID(id int64) (entity.Todo, error) {
	todo := entity.Todo{}
	res := r.db.Where("id = ?", id).First(&todo)
	if res.Error != nil {
		return todo, res.Error
	}
	return todo, nil
}

func (r todoRepository) Create(todo entity.Todo) (int64, error) {
	var uid int64
	res := r.db.Create(&todo)
	if res.Error != nil {
		return uid, res.Error
	}
	uid = todo.ID
	return uid, nil
}

func (r todoRepository) Update(id int64, todo entity.Todo) error {
	res := r.db.Where("id = ?", id).Updates(todo)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r todoRepository) Delete(id int64) error {
	res := r.db.Delete(&entity.Todo{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
