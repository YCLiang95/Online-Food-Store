package mysql_utils

import (
"github.com/YCLiang95/CS160Group1OFS/backend/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
)

var (
	mysqlEngin *xorm.Engine = nil
)



func initMysqlEngin() *xorm.Engine {

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		protocal.G_Config.MysqlConfig.UserName,
		protocal.G_Config.MysqlConfig.Password,
		protocal.G_Config.MysqlConfig.Address,
		protocal.G_Config.MysqlConfig.DbName)

	engine, err := xorm.NewEngine("mysql", mysqlInfo)
	if err != nil {
		utils.Logger.Error("Failed to set up mysql:", "--ERROR--", err)
		panic(err)
	}
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(protocal.G_Config.MysqlConfig.MaxIdel)
	engine.SetMaxOpenConns(protocal.G_Config.MysqlConfig.MaxOpen)
	utils.Logger.Notice("Mysql connection set up:", mysqlInfo)
	return engine
}


func GetInstance() *xorm.Engine {
	if mysqlEngin == nil {
		mysqlEngin = initMysqlEngin()
	}
	return mysqlEngin
}
