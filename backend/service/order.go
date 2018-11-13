package service

import (
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"github.com/YCLiang95/CS160Group1OFS/backend/dao"
	"errors"
)

var (
	saveErr  = errors.New("7000:save order failed")
	getErr  = errors.New("7002:get order failed")
)

func ListOrder() (orders []*protocal.Order, err error) {
	if orders, err = dao.GetOrder().GetOrderList();err!=nil{
		orders=nil
		err = getErr
	}
	return
}

func SaveOrder (orderID int,status int,userID int,delivID int, totPay float, payID int, ordDet) (err error) {
	if err = dao.GetMechandisObject().InsertNewOrder(&protocal.Order{
		OrderID: orderID,
	  Status: status,
	  UserID: userID,
	  DeliveryID: delivID,
	  TotalPaymentPrice: totPay,
	  PaymentID: payID,
	  OrderDetails: ordDet,
	});err!=nil{
		err = saveErr
	}
	return
}
