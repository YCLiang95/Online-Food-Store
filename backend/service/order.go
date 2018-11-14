package service

import (
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"time"
	"strconv"
	"github.com/YCLiang95/CS160Group1OFS/backend/dao"

	"errors"
)

func PlaceOrder(uid int64,paymentType int, orderDetail []protocal.OrderReequestDetail, deliveryRequest protocal.DeliveryRequest)(err error){
	var(
		merchandises map[int64]*protocal.Merchandise
		totalValue float64
		orderDetails []*protocal.OrderDetail
		delivery *protocal.Delivery
	)
	order :=new(protocal.Order)
	order.Uid = uid
	order.Status =1
	order.PaymentMethod = paymentType
	order.OrderId=getOrderid(int(uid))
	//Get all the merchandise which is need by order
	if merchandises,err=dao.GetMechandisObject().GetMerchandises(orderDetail);err!=nil{
		return
	}
    //Calculate the total price
	orderDetails = make([]*protocal.OrderDetail,0)
	for _,value:=range orderDetail{
		totalValue+=float64(value.Count)*merchandises[value.Mid].Price
		orderDetails = append(orderDetails,&protocal.OrderDetail{
           MerchandiseName:merchandises[value.Mid].Name,
           MerchandiseCount: int64(value.Count),
           MerchandisePrice:merchandises[value.Mid].Price,
           MerchandiseId:value.Mid,
		})
	}
	order.TotalMerchandisePrice = totalValue
	order.TotalPaymentPrice = totalValue

	//Implement delivery detail
	delivery = &protocal.Delivery{
		Status:1,
		State:deliveryRequest.State,
		Address:deliveryRequest.Address,
		Zipcode:int64(deliveryRequest.Zipcode),
		DeliveryId:getOrderid(int(uid)),
		City:deliveryRequest.City,
	}
   if err=dao.GetOrderDao().InsertOrder(order,orderDetails,delivery);err!=nil{
       	return
   }

    return
}

func GetOrders(uid int64,startIndex,count int)([]*protocal.OrderResponse,error){
	if result,err:=dao.GetOrderDao().GetOrdersByUid(uid,startIndex,count);err!=nil{
		return nil,errors.New("1002:System error")
	}else{
		return result,nil
	}

}


func getOrderid(uid int)(id string){

	id=time.Now().Format("20060102150405")
	id+=strconv.Itoa(int(uid))
	uninxM:=strconv.Itoa(int(time.Now().UnixNano()/1e6))
	id+=uninxM[len(uninxM)-2:]

	return
}