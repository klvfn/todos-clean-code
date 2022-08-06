package mysql

import (
	"fmt"

	"github.com/klvfn/todos-clean-code/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ConnectMysql init mysql connection
func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.AppConfig.Mysql.User, config.AppConfig.Mysql.Pass, config.AppConfig.Mysql.Host, config.AppConfig.Mysql.Port, config.AppConfig.Mysql.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return db, err
	}
	return db, nil
}

type MysqlRepo struct {
	Todo TodoRepository
}

func NewMysqlRepo(db *gorm.DB) *MysqlRepo {
	todoRepo := NewTodoRepository(db)

	return &MysqlRepo{
		Todo: todoRepo,
	}
}
