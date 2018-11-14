package api

import (
	"net/http"

	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"github.com/YCLiang95/CS160Group1OFS/backend/service"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) (model *protocal.ResponseModel, err error) {
	var(
		orderRequestDetail []protocal.OrderReequestDetail
		deliveryRequest   protocal.DeliveryRequest
	)
	orderRequest:=protocal.OrderRequest{}
	if  err = GetStructFromRequest(r,&orderRequest);err!=nil{
		return nil,err
	}
	if err=json.Unmarshal([]byte(orderRequest.OrderDetail),&orderRequestDetail);err!=nil{
		fmt.Println(err)
		goto ERR
	}
	if err=json.Unmarshal([]byte(orderRequest.DeliveryDetail),&deliveryRequest);err!=nil{
		fmt.Println(err)
		goto ERR

	}
	service.PlaceOrder(orderRequest.Uid,orderRequest.PaymentType,orderRequestDetail,deliveryRequest)
	return
ERR:

		return nil,errors.New("1001:system error")
}


func GetOrders(w http.ResponseWriter, r *http.Request) (model *protocal.ResponseModel, err error) {

	var(
		page int
		count int
		startIndex int
		orders []*protocal.OrderResponse
	)

	if err=r.ParseForm();err!=nil{
		goto ERR
	}

	if page,err=strconv.Atoi(r.Form.Get("page"));err!=nil{
		goto ERR
	}
	if count,err=strconv.Atoi(r.Form.Get("count"));err!=nil{
		goto ERR
	}


	startIndex = (page - 1) * count

   if orders,err=service.GetOrders(3,startIndex,count);err!=nil{
   	fmt.Println(err)
	   goto ERR
   }

   model = &protocal.ResponseModel{
   	Data:orders,
   	Code:200,
   	Message:"get orders successfully",
   }
   return

ERR:
	return nil,errors.New("1001:system error")

}