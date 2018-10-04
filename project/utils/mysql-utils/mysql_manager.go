package mysql_utils

import (
"github.com/cs160/project/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
//	"log"
//	"log/syslog"
)

var (
	mysqlEngin *xorm.Engine = nil
	config     *MysqlConfig = nil
)

func initMysqlEngin() *xorm.Engine {

	if config == nil {
		utils.Logger.Error("mysql configs is not instance")
	}
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.UserName, config.Password, config.Address, config.DbName)
	engine, err := xorm.NewEngine("mysql", mysqlInfo)
	if err != nil {
		utils.Logger.Error("Failed to set up mysql:", "--ERROR--", err)
		panic(err)
	}
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(config.MaxIdel)
	engine.SetMaxOpenConns(config.MaxOpen)
	utils.Logger.Notice("Mysql connection set up:", mysqlInfo)
	return engine
}

func SetConfig(in *MysqlConfig) {
	config = in
}

func GetInstance() *xorm.Engine {
	if mysqlEngin == nil {
		mysqlEngin = initMysqlEngin()
	}
	return mysqlEngin
}
