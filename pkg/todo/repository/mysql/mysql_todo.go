package mysql

import (
	"github.com/klvfn/todos-clean-code/pkg/todo/entity"
	"gorm.io/gorm"
)

// TodoRepository contract for redis todo repository
type TodoRepository interface {
	GetAll() ([]entity.Todo, error)
	GetByID(id int64) (entity.Todo, error)
	Create(todo entity.Todo) (int64, error)
	Update(id int64, todo entity.Todo) error
	Delete(id int64) error
}

type todoRepository struct {
	MysqlDB *gorm.DB
}

// NewTodoRepository create new instance of redis todo repository
func NewTodoRepository(mysqlDB *gorm.DB) TodoRepository {
	return &todoRepository{
		MysqlDB: mysqlDB,
	}
}

func (tr todoRepository) GetAll() ([]entity.Todo, error) {
	todos := []entity.Todo{}
	res := tr.MysqlDB.Find(&todos)
	if res.Error != nil {
		return todos, res.Error
	}
	return todos, nil
}

func (tr todoRepository) GetByID(id int64) (entity.Todo, error) {
	todo := entity.Todo{}
	res := tr.MysqlDB.Where("id = ?", id).First(&todo)
	if res.Error != nil {
		return todo, res.Error
	}
	return todo, nil
}

func (tr todoRepository) Create(todo entity.Todo) (int64, error) {
	var uid int64
	res := tr.MysqlDB.Create(&todo)
	if res.Error != nil {
		return uid, res.Error
	}
	uid = todo.ID
	return uid, nil
}

func (tr todoRepository) Update(id int64, todo entity.Todo) error {
	res := tr.MysqlDB.Where("id = ?", id).Updates(todo)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (tr todoRepository) Delete(id int64) error {
	res := tr.MysqlDB.Delete(&entity.Todo{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
