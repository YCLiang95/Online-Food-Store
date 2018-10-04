package mysql_utils

import (
	"OFBankBrowserServer/utils"
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	utils.CreateLogger("test")
	SetConfig(LoadConfigFile("../../mysql_json.conf"))
	mysql := GetInstance()
	fmt.Println(mysql)
}
