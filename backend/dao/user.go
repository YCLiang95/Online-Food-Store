package dao

import (
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils/mysql-utils"
//	"github.com/tianyun6655/Ofbank_Quantitative_Transaction/utils"
	"fmt"
)

type UserDaoInstance struct{}

var userDaoObject *UserDaoInstance = nil

func GetInstance() *UserDaoInstance {
	if userDaoObject == nil {
		userDaoObject = new(UserDaoInstance)
	}
	return userDaoObject
}

func (ui *UserDaoInstance) Register(email, password string) (err error) {
	user := &protocal.ProjectUser{Email: email, Password: password}
	if err = daoFunctionLogWapper(user, nil, saveRecorders); err != nil {
		return
	}
   return
}

func (ui *UserDaoInstance) GetUser(email string) (user *protocal.ProjectUser, err error) {
	user= &protocal.ProjectUser{Email:email}

	if err = daoFunctionLogWapper(user,nil,getRecorder); err != nil {
		user = nil
		return
	}
	if user.Uid==0{
		user=nil
	}
	return
}
