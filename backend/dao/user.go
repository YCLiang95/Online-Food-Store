package dao

import (
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
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
	 err = daoFunctionLogWapper(user, nil,saveRecorders)
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
