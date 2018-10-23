package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils"
	"fmt"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils/mysql-utils"
	"reflect"
	"errors"
)

type mysqlOptions struct {
	OrderBy    string
	LimitStart int
	Count      int
	TableName  string
	Filter     interface{}
}

func daoFunctionLogWapper(bean interface{}, options *mysqlOptions, daofunc func(sqlEnginer *xorm.Engine, data interface{}, options *mysqlOptions)error) ( err error) {
	if  err = daofunc(mysql_utils.GetInstance(), bean, options); err != nil {
		tableName := reflect.TypeOf(bean).String()
		LoggerFomat(bean, tableName, err)

	}
	return
}

func saveRecorders(sqlEnginer *xorm.Engine, data interface{}, options *mysqlOptions) (err error) {
	_,err = sqlEnginer.Insert(data)
	return
}

func getRecorder(sqlEnginer *xorm.Engine, bean interface{}, options *mysqlOptions) (err error) {
	_,err=sqlEnginer.Get(bean)
	return
}

func findRecorders(sqlEnginer *xorm.Engine, bean interface{}, options *mysqlOptions) (err error) {
	var (
		mysqlSession *xorm.Session
	)
	if options == nil {
		err = sqlEnginer.Find(bean)
		return
	}

	if options.TableName != "" {
		mysqlSession = sqlEnginer.Table(options.TableName)
	}
	if options.OrderBy != "" {
		if mysqlSession == nil {
			mysqlSession = sqlEnginer.OrderBy(options.OrderBy)
		} else {
			mysqlSession = mysqlSession.OrderBy(options.OrderBy)
		}
	}
	if options.LimitStart < 0 || options.Count <= 0 {
		err = errors.New("option params is incorrect")
		return
	}
	if options.LimitStart != 0 || options.Count != 0 {
		if mysqlSession == nil {
			mysqlSession = sqlEnginer.Limit(options.LimitStart, options.Count)
		} else {
			mysqlSession = mysqlSession.Limit(options.LimitStart, options.Count)
		}
	}
	if options.Filter != nil {
		err = mysqlSession.Find(bean, options.Filter)
	} else {
		err = mysqlSession.Find(bean)
	}
	return
}


func LoggerFomat(params interface{},table string, err error) {
	utils.Logger.Error(fmt.Sprintf(`
########MYSQL ERROR########

ERROR: %v

PARAMS:%v

Table: %v

###########################
      `, err,params,table))
}
