package service

import (
	"errors"
	"github.com/cs160/project/dao"
	"github.com/cs160/project/dao/types"
)

var(
	RegisterError=errors.New("1000:failed to register user")
	LoginFailedErr=errors.New("1001:email/password is incorrect")
	UserSystemErr=errors.New("2000:network error,please try again")
)


func Register(email,password string)error{
	if err:=dao.GetUserDao().Register(email,password);err!=nil {
		return RegisterError
	}
	return nil
}

func Login(email,password string)( *types.ProjectUser,error) {
	user, err := dao.GetUserDao().GetUser(email)
	if err != nil{
		return nil,UserSystemErr
	}
	if user==nil||user.Password!=password{
		return nil,LoginFailedErr
	}
	return user,nil
}