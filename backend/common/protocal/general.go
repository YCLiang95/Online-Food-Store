package protocal

import (
	"io/ioutil"
	"encoding/json"
)

var G_Config  *Config


type MySqlConfig struct {
	UserName string `json:"userName"`
	Address  string `json:"address"`
	Port     string `json:"port"`
	DbName   string `json:"dbName"`
	Password string `json:"password"`
	MaxIdel  int    `json:"maxIdel"`
	MaxOpen  int    `json:"maxOpen"`
}


type RedisConfig struct {
	MaxIdle int  `json:"maxIdle"`
	Timeout int  `json:"timeOut"`
	Host string   `json:"host"`
	DB int   `json:"db"`
	Password string `json:"password"`
	PoolSize int  `json:"pool_size"`
}
type Config struct {
	MysqlConfig MySqlConfig  `json:"mysql"`
    RedisConfig  RedisConfig `json:"redis"`
    GoogleKey string `json:"googleKey"`
}



func LoadConfig(filePath string, desc interface{}) (err error) {
	var (
		configByte []byte
	)
	if configByte, err = ioutil.ReadFile(filePath); err != nil {
		return
	}
	if err = json.Unmarshal(configByte, desc); err != nil {
		return
	}

	return
}

type ResponseModel struct {
	Message string  `json:"message"`
	Data  interface{} `json:"data"`
	Code int   `json:"code"`
}

type PageResponse struct {
	Total int     `json:"total"`
	Current int    `json:"current"`
	NextPage int  `json:"nextPage"`
	PrePage int    `json:"Prepage"`
	Data  interface{}   `json:"data"`
}

func GenerateSuccessStruct(message string,data interface{})*ResponseModel{
	return &ResponseModel{
		Data:data,
		Message:message,
		Code:200,
	}
}


