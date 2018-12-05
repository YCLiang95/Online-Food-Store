package main

import (
	"flag"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils"
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"github.com/YCLiang95/CS160Group1OFS/backend/control"
	//"github.com/YCLiang95/CS160Group1OFS/backend/utils/redis-utils"
)

var SERVER_PORT string

func main(){
	flag.StringVar(&SERVER_PORT, "port", "8085", "http server port")
	utils.CreateLogger("back-end.log")
	if err:=protocal.LoadConfig("./config.json",&protocal.G_Config);err!=nil{
		utils.Logger.Error("Failed to load config: ",err)
		return
	}
	flag.Parse()
	httpser := control.NewHttpServer(SERVER_PORT)
	httpser.StartServer()

}


func ForTest(){}