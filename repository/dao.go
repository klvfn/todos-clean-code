package repository

import "github.com/klvfn/todos-clean-code/repository/mysql"

type Dao struct {
	MysqlRepo *mysql.MysqlRepo
}

func NewDao(mysqlRepo *mysql.MysqlRepo) *Dao {
	return &Dao{
		MysqlRepo: mysqlRepo,
	}
}
