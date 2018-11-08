package main

import (
	"flag"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils/mysql-utils"
	"github.com/YCLiang95/CS160Group1OFS/backend/control/filter"
)

var SERVER_PORT string

func main(){
	flag.StringVar(&SERVER_PORT, "port", "8085", "http server port")
	utils.CreateLogger("Mobile_backend.log")
	mysql_utils.SetConfig(mysql_utils.LoadConfigFile("./mysql_json_local.conf"))
	flag.Parse()
	httpser := filter.NewHttpServer(SERVER_PORT)
	httpser.StartServer()

}
