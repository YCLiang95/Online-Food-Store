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


func updateRecorder(sqlEngine *xorm.Engine, bean interface{}, options *mysqlOptions) (err error) {
	var (
		filterMap map[string]interface{}
		ok        bool
		session   *xorm.Session
		updateCount int64
	)
	if filterMap, ok = options.Filter.(map[string]interface{}); !ok {
		err = errors.New("illegal filter, please set to map")
		return
	}
	session = sqlEngine.NewSession()
	for key, value := range filterMap {
		session=session.Where(key, value)
	}

	if updateCount, err = session.Update(bean);updateCount==0&&err==nil{
	err = errors.New("nothing is changed")
		return
	}
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
	mysqlSession = sqlEnginer.NewSession()
	if options.TableName != "" {
		mysqlSession = sqlEnginer.Table(options.TableName)
	}
	if options.OrderBy != "" {
		mysqlSession = mysqlSession.OrderBy(options.OrderBy)
	}
	if options.LimitStart < 0 || options.Count <= 0 {
		err = errors.New("option params is incorrect")
		return
	}
	if options.LimitStart != 0 || options.Count != 0 {
		mysqlSession = mysqlSession.Limit(options.LimitStart, options.Count)

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
