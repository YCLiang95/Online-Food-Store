package dao

import (
	"github.com/cs160/project/dao/types"
	"github.com/cs160/project/utils/mysql-utils"
	"github.com/cs160/project/utils"
)


var userDaoObject *UserDaoInstance = nil


type UserDaoInstance struct {
	DaoInstance
}


func GetUserDao() *UserDaoInstance {
	if userDaoObject == nil {
		userDaoObject = new(UserDaoInstance)
		userDaoObject.tableName = "project_user"
		userDaoObject.databaseEnginer = mysql_utils.GetInstance()
	}
	return userDaoObject
}

func (ui *UserDaoInstance) Register(email, password string) error {
	user := &types.ProjectUser{Email: email, Password: password}
	_, err := ui.databaseEnginer.Insert(user)
	if err != nil {
		utils.Logger.Error("--UserDao--"," Failed to insert user informatiom: ",err)
		return err
	} else {
		return nil
	}
}

func (ui *UserDaoInstance) GetUser(email string) (*types.ProjectUser, error) {
	 user:=new(types.ProjectUser)
	user.Email = email
	bool, err := ui.databaseEnginer.Get(user)
	if err != nil {
		utils.Logger.Error("--UserDao--"," Failed to Get user informatiom: ",err)

		return nil, err
	}
	if !bool {
		return nil, nil
	}
	return user, nil
}