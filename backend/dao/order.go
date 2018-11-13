package dao

import(
  "github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
)

type OrderDaoInstance struct{}

var orderDaoObject *OrderDaoInstance = nil

func GetOrder() *OrderDaoInstance {
	if orderDaoObject == nil {
		orderDaoObject = new(OrderDaoInstance)
	}
	return orderDaoObject
}

func (ord *OrderDaoInstance) GetOrderList() (orders []*protocal.Order, err error){
  orders = make([]*protocal.Order, 0)

  err = daoFunctionLogWapper(&orders, nil, findRecorders);
  return
}

func (ord *OrderDaoInstance) InsertNewOrder(order *protocal.Order) (err error) {
	err = daoFunctionLogWapper(order, nil, saveRecorders)
	return
}

//Search for order by order id
func (ord *OrderDaoInstance) GetOrder(int orderID) (order *protocal.Order, err error){
  order = &protocal.Order{OrderID: orderID}
	err = daoFunctionLogWapper(order, nil, getRecorder)
	return
}
