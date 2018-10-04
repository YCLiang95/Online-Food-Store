package dao

import "github.com/go-xorm/xorm"

type DaoInstance struct {
	tableName       string
	databaseEnginer *xorm.Engine
}

