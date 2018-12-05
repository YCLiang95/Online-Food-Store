package api

import (
	"net/http"

	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"github.com/YCLiang95/CS160Group1OFS/backend/service"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"io/ioutil"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) (model *protocal.ResponseModel, err error) {
	var (
		placeOrderRequest  protocal.PlaceOrderRequest
		body []byte
	)
	body, _ = ioutil.ReadAll(r.Body)
	if err = json.Unmarshal(body, &placeOrderRequest); err != nil {
		fmt.Println(err)
		err = errors.New("1001:system error")
		goto ERR
	}
	if err = service.PlaceOrder(placeOrderRequest.Uid, placeOrderRequest.PaymentType, placeOrderRequest.OrderDetail, placeOrderRequest.DeliveryDetail); err != nil {
		fmt.Println(err)
		err=errors.New("1001:system error")
		goto ERR

	}
	model = &protocal.ResponseModel{
		Data:    nil,
		Code:    200,
		Message: "place orders successfully",
	}

	return
ERR:
	return nil, err
}

func GetOrders(w http.ResponseWriter, r *http.Request) (model *protocal.ResponseModel, err error) {
	var (
		page       int
		count      int
		startIndex int
		orders     []*protocal.OrderResponse
		uid        int64
	)

	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	if page, err = strconv.Atoi(r.Form.Get("page")); err != nil {
		goto ERR
	}
	if count, err = strconv.Atoi(r.Form.Get("count")); err != nil {
		goto ERR
	}

	if uid, err = strconv.ParseInt(r.Form.Get("uid"), 10, 64); err != nil {
		fmt.Println(err)
		goto ERR
	}

	startIndex = (page - 1) * count

	if orders, err = service.GetOrders(uid, startIndex, count); err != nil {
		fmt.Println(err)
		goto ERR
	}

	model = &protocal.ResponseModel{
		Data:    orders,
		Code:    200,
		Message: "get orders successfully",
	}
	return

ERR:
	return nil, errors.New("1001:system error")

}


func HandlerDelivery(w http.ResponseWriter, r *http.Request) (model *protocal.ResponseModel, err error) {
	var(
		Did int
		delivery *protocal.Delivery

	)
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	if Did,err=strconv.Atoi(r.Form.Get("did"));err!=nil{
		fmt.Println(err)
		goto ERR
	}


	if delivery,err =service.GetDeliveryDetail(int64(Did));err!=nil{
		goto ERR
	}
	model = &protocal.ResponseModel{
		Data:    delivery,
		Code:    200,
		Message: "get delivery successfully",
	}
	return

ERR:
	return nil, errors.New("1001:system error")

}