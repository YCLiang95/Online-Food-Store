package dao

import (
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils/mysql-utils"
	"github.com/go-xorm/xorm"
	"fmt"
	"time"
)

type OrderDao struct{}

var orderDao *OrderDao = nil

func GetOrderDao() *OrderDao {
	if orderDao == nil {
		orderDao = new(OrderDao)
	}
	return orderDao
}

func (od *OrderDao) InsertOrder(order *protocal.Order, orderDetails []*protocal.OrderDetail, delivery *protocal.Delivery) (err error) {
	var (
		mysqlSession *xorm.Session
	)
	mysqlSession = mysql_utils.GetInstance().NewSession()
	if err = mysqlSession.Begin(); err != nil {
		LoggerFomat("DB Transction", "order", err)
		return
	}
	if _, err = mysqlSession.Insert(order); err != nil {
		LoggerFomat("DB Transction", "order", err)
		mysqlSession.Rollback()
		return
	}
	for i := 0; i < len(orderDetails); i++ {
		orderDetails[i].Oid = order.Oid
	}
	if _, err = mysqlSession.Insert(orderDetails); err != nil {
		LoggerFomat("DB Transction", "order", err)
		mysqlSession.Rollback()
		return
	}
	delivery.Oid = order.Oid
	if _, err = mysqlSession.Insert(delivery); err != nil {
		LoggerFomat("DB Transction", "order", err)
		mysqlSession.Rollback()
		return
	}
	mysqlSession.Commit()

	return
}

func (od *OrderDao) GetOrdersByUid(uid int64, startIndex, count int) (orders []*protocal.OrderResponse, err error) {
	var (
		result      []map[string]interface{}
		orderResult map[int64]*protocal.OrderResponse
	)
	if result, err = mysql_utils.GetInstance().QueryInterface(
		"SELECT * FROM"+
			"(SELECT `order`.oid,"+
			"`order_id`,"+
			"`order`.`create_time`,"+
			"`order`.status,"+
			"`total_merchandise_price`,"+
			"total_payment_price,"+
			"payment_method,"+
			"`delivery`.`delivery_id`,"+
			"`delivery`.`status` as `delivery_status`,"+
			"zipcode,"+
			"did,"+
			"address,"+
			"city,"+
			"state,"+
			"`delivery`.`create_time` as `delivery_create_time` "+
			"FROM `order` LEFT JOIN `delivery` on `order`.oid = `delivery`.oid WHERE `uid`=? Limit ?,? )c "+
			"LEFT JOIN `order_detail` ON c.oid = `order_detail`.oid", uid, startIndex, count); err != nil {
		LoggerFomat("DB Order get join", "order", err)
		return nil, err
	}
	fmt.Println(result)
	orderResult = make(map[int64]*protocal.OrderResponse)
	for _, value := range result {
		if _, ok := orderResult[value["oid"].(int64)]; !ok {
			orderResponse := new(protocal.OrderResponse)
			orderResponse.Order.Oid = value["oid"].(int64)
			orderResponse.Order.Uid = uid
			orderResponse.Order.UpdateTime = nil
			orderResponse.Order.PaymentMethod = value["payment_method"].(int64)
			orderResponse.Order.Status = value["status"].(int64)
			orderResponse.Order.CreateTime, _ = time.Parse("2006-01-02 15:04:05", string(value["create_time"].([]byte)))
			orderResponse.Order.TotalMerchandisePrice = float64(value["total_merchandise_price"].(float32))
			orderResponse.Order.TotalPaymentPrice = float64(value["total_payment_price"].(float32))
			orderResponse.Order.OrderId = string(value["order_id"].([]byte))
			orderResponse.Delivery.CreateTime, _ = time.Parse("2006-01-02 15:04:05", string(value["delivery_create_time"].([]byte)))
			orderResponse.Delivery.DeliveryId = string(value["delivery_id"].([]byte))
			orderResponse.Delivery.City = string(value["city"].([]byte))
			orderResponse.Delivery.State = string(value["state"].([]byte))
			orderResponse.Delivery.Zipcode = value["zipcode"].(int64)
			orderResponse.Delivery.Status = value["delivery_status"].(int64)
			orderResponse.Delivery.Address = string(value["address"].([]byte))
			orderResponse.Delivery.Did = value["did"].(int64)
			orderResult[value["oid"].(int64)] = orderResponse
		}
		orderDetial := new(protocal.OrderDetail)
		orderDetial.MerchandiseId = value["merchandise_id"].(int64)
		orderDetial.MerchandiseCount = value["merchandise_count"].(int64)
		orderDetial.MerchandisePrice = float64(value["merchandise_price"].(float32))
		orderDetial.MerchandiseName = string(value["merchandise_name"].([]byte))
		orderResult[value["oid"].(int64)].OrderDetail = append(orderResult[value["oid"].(int64)].OrderDetail, orderDetial)
	}

	orders = make([]*protocal.OrderResponse, 0)
	for _, value := range orderResult {
		orders = append(orders, value)
	}
	return
}

func (od *OrderDao) GetDeliveryDetail(did int64) (delivery *protocal.Delivery, err error) {
	delivery = &protocal.Delivery{
		Did: did,
	}
	_, err = mysql_utils.GetInstance().Get(delivery)
	return
}

func (od *OrderDao)UpdateDeliveryLocation(orderId string,preX,preY float64,currentX, currentY float64,status int64){

	delivery := &protocal.Delivery{
          PreX:preX,
          PreY:preY,
          CurrentY:currentY,
          CurrentX:currentX,
          Status:status,
	}
	_,err:=mysql_utils.GetInstance().
		Table("delivery").
		Where("delivery_id=?", orderId).
		Update(delivery)
		if err!=nil{
			fmt.Println(err)
		}
}

