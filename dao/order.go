package dao

import(
  "github.com/cs160/project/utils/mysql-utils"
  "github.com/cs160/project/utils"
)

type orderDaoInstance{
  DaoInstance
}

type Order struct{

}

func GetInvDao() *orderDaoInstance {
	if orderDaoObject == nil {
		orderDaoObject = new(InvDaoInstance)
		orderDaoObject.tableName = "order"
		orderDaoObject.databaseEnginer = mysql_utils.GetInstance()
	}
	return orderDaoObject
}

func (order *orderDaoInstance) getOrder(int userID){
  order, err := order.Where("userid = ?", userID).Find(&order)
  if err != nil{
		utils.Logger.Error("--OrderDao--"," Failed to Get Order Information: ", err)
		return nil, err
	}
  return order
}

func (order *orderDaoInstance) saveOrder(string username){
  db.Query("update order set where ")
}

func (order *orderDaoInstance) getAllOrders(string username){
  var orders []Order
  orders, err := order.Cols("order_id", "order_details").Find(&order)
  return orders
}
