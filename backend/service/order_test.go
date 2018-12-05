package service

import (
	"testing"
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"fmt"
)

func TestGetLoaction(t *testing.T) {
	if err:=protocal.LoadConfig("../project/config.json",&protocal.G_Config);err!=nil{
       fmt.Println(err)
		return
	}

	fmt.Println(GetLoaction("107 N Amphlett Blvd, San Mateo, CA 94401"))
}
