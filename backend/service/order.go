package service

import (
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"time"
	"strconv"
	"github.com/YCLiang95/CS160Group1OFS/backend/dao"

	"errors"
	"strings"
	"github.com/valyala/fasthttp"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils"
	"fmt"
	"github.com/tidwall/gjson"
)

const (
	StartX1  =37.3534093
	StartY1  =-121.8540879
	StartXY ="37.3534093,-121.8540879"
	StartX2 = 37.5752331
	StartY2 = -122.3196909
	StartXY2 = "37.5752331, -122.3196909"
	Zipcode = 95116
	Zipcode2 = 94401

)




func PlaceOrder(uid int64,paymentType int, orderDetail []protocal.OrderReequestDetail, deliveryRequest protocal.DeliveryRequest)(err error){
	var(
		merchandises map[int64]*protocal.Merchandise
		totalValue float64
		orderDetails []*protocal.OrderDetail
		delivery *protocal.Delivery
		address string
		endx  float64
		endy float64
	)

	address=generateAddress(deliveryRequest.Address,deliveryRequest.State,deliveryRequest.City,deliveryRequest.Zipcode)
	if endx,endy,err = GetLoaction(address);err!=nil{
		return
	}
	order :=new(protocal.Order)
	order.Uid = uid
	order.Status =1
	order.PaymentMethod = int64(paymentType)
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

	useX:=StartX1
	useY:=StartY1
	useXY:=StartXY
	if deliveryRequest.Zipcode-Zipcode>deliveryRequest.Zipcode-Zipcode2{
		useX = StartX2
		useY =StartY2
		useXY = StartXY2
	}
	//Implement delivery detail
	delivery = &protocal.Delivery{
		Status:1,
		State:deliveryRequest.State,
		Address:deliveryRequest.Address,
		Zipcode:int64(deliveryRequest.Zipcode),
		DeliveryId:getOrderid(int(uid)),
		City:deliveryRequest.City,
		EndX: endx,
		EndY:endy,
		CurrentX:useX,
		CurrentY:useY,
		PreX:useX,
		PreY:useY,
		StartX: useX,
		StartY:useY,
	}
   if err=dao.GetOrderDao().InsertOrder(order,orderDetails,delivery);err!=nil{
       	return
   }

   go UpdateDelivery(delivery.DeliveryId,useXY,endx,endy)

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

func generateAddress(Address string,State string,City string, zipCode int)string{

	stateWithZip := State +" "+ strconv.Itoa(zipCode)

	arr:=make([]string,3)
	arr[0] = Address
	arr[1] = City
	arr[2] = stateWithZip

	return  strings.Join(arr,",")
}



func GetLoaction(address string) (endx, endy float64,err error) {
	var (
		reponseResult []byte
		responseCode  string
	)

	fastArgs := &fasthttp.Args{}
	fastArgs.Add("key", protocal.G_Config.GoogleKey)
	fastArgs.Add("query", address)
	if reponseResult, err = utils.SendRequest("https://maps.googleapis.com/maps/api/place/textsearch/json", utils.GET, fastArgs, nil); err != nil {
		fmt.Println(err)
		err = errors.New("1005:location error")
		return
	} else {
		responseCode = gjson.Get(string(reponseResult), "status").Str
		if responseCode != "OK" {
			err = errors.New("1005:location error")
			return
		}
		results := gjson.Get(string(reponseResult), "results.0").Raw
		endx = gjson.Get(results, "geometry.location.lat").Float()
		endy = gjson.Get(results, "geometry.location.lng").Float()
		return endx, endy, nil
	}
}

func GetDeliveryDetail(did int64)(delivery *protocal.Delivery,err error){
	if delivery,err = dao.GetOrderDao().GetDeliveryDetail(did);err!=nil{
		fmt.Println(err)
	}
	return
}

func UpdateDelivery(deliveryId string,startLocationXY string,EndLocationX, EndlocationY float64){

	type location struct {
		x float64
		y float64
	}
	var steps []location

	requestParams := &fasthttp.Args{}
	requestParams.Add("origin", startLocationXY)
	requestParams.Add("key", "AIzaSyDFC1JzMNhhZvnBlguIeSE3UBWKj6IBVKU")
	endLocationXString := strconv.FormatFloat(EndLocationX, 'f', -7, 64)
	endLocationYString := strconv.FormatFloat(EndlocationY, 'f', -7, 64)

	requestParams.Add("destination", endLocationXString+","+endLocationYString)
	if response, err := utils.SendRequest("https://maps.googleapis.com/maps/api/directions/json", utils.GET, requestParams, nil); err != nil {
		return
	}else {
		responseSteps := gjson.Get(string(response), "routes.0.legs.0.steps").Array()
		steps = make([]location, len(responseSteps))
		for i := 0; i < len(steps); i++ {
			newlocation := location{}
			if i == len(steps)-1 {
				newlocation = location{
					x: EndLocationX,
					y: EndlocationY,
				}
			} else {
				newlocation = location{
					x: gjson.Get(responseSteps[i].Raw, "end_location.lat").Float(),
					y: gjson.Get(responseSteps[i].Raw, "end_location.lng").Float(),
				}
			}
          steps[i]= newlocation
		}
	}
	for i:=0;i<len(steps);i++{
		time.Sleep(5*time.Second)
		if i==0{
			dao.GetOrderDao().UpdateDeliveryLocation(deliveryId,0,0,steps[i].x,steps[i].y,1)
		}else if i==len(steps)-1{
			dao.GetOrderDao().UpdateDeliveryLocation(deliveryId,steps[i-1].x,steps[i-1].y,steps[i].x,steps[i].y,2)
		}else{
			dao.GetOrderDao().UpdateDeliveryLocation(deliveryId,steps[i-1].x,steps[i-1].y,steps[i].x,steps[i].y,1)

		}

	}
}