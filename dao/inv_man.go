package dao

import (
	"github.com/cs160/project/utils/mysql-utils"
	"github.com/cs160/project/utils"
)

var invDaoObject *invDaoInstance = nil

type invDaoInstance struct{
	DaoInstance
}
//For inventory management, only the name and quantity are of importance
type Merch struct{
	string merchName
	int quantity
}

func GetInvDao() *InvDaoInstance {
	if invDaoObject == nil {
		invDaoObject = new(InvDaoInstance)
		invDaoObject.tableName = "merchandise"
		invDaoObject.databaseEnginer = mysql_utils.GetInstance()
	}
	return invDaoObject
}

//Get all items
func (inv *invDaoInstance) GetAllMerch(){
	var merch []Merch
	merch, err := inv.Cols("name", "quantity").Find(&inv)
	if err != nil{
		utils.Logger.Error("--InvDao--"," Failed to Get inventory information: ", err)
		return nil, err
	}
	return merch
}

//Get all rows where item quantity is low
func (inv *InvDaoInstance) GetLowQuantity(){
	var lowQ []Merch
	lowQ, err := inv.Where(quantity < 10).Cols("name", "quantity").Find(&inv)
	if err != nil{
		utils.Logger.Error("--InvDao--"," Failed to Get inventory information: ", err)
		return nil, err
	}
	return lowQ
}

//Update inventory quantity with input value
func (inv *InvDaoInstance) Update(string name, int newQ){
	newQ, err := inv.Where("name = ?", name).Cols(quantity).Update(&inv)
	if err != nil{
		utils.Logger.Error("--InvDao--", "Failed to update information: ", err)
		return nil, error
	}
	return nil
}
