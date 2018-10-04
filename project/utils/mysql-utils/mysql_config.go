package mysql_utils

import (
	"github.com/cs160/project/utils"
	"encoding/json"
	"io/ioutil"
)

type MysqlConfig struct {
	UserName string `json:"userName"`
	Address  string `json:"address"`
	Port     string `json:"port"`
	DbName   string `json:"dbName"`
	Password string `json:"password"`
	MaxIdel  int    `json:"maxIdel"`
	MaxOpen  int    `json:"maxOpen"`
}

func LoadConfigFile(url string) *MysqlConfig {
	configData, e := ioutil.ReadFile(url)

	if e != nil {
		utils.Logger.Error(url+" file read fail:", e)
	}
	var config MysqlConfig
	err := json.Unmarshal(configData, &config)
	if err != nil {
		utils.Logger.Error("load json mysql config fail:", err)
	}
	return &config
}
