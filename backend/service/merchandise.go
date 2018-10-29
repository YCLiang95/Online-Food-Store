package service

import (
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"github.com/YCLiang95/CS160Group1OFS/backend/dao"
	"errors"
)


var (
	saveErr  = errors.New("3000:save merchandises failed")
	updateErr  = errors.New("3001:update merchandises failed")
	getErr  = errors.New("3002:get merchandises failed")


)
func List() (mechandis []*protocal.Merchandise, err error) {
	if mechandis, err = dao.GetMechandisObject().List();err!=nil{
		mechandis=nil
		err = getErr
	}
	return
}

func Save(name string,weight float64,price float64,quantity int64, imageUrl string) (err error) {
	if err = dao.GetMechandisObject().InsertNewMechanDis(&protocal.Merchandise{
		Name:name,
		Weight:weight,
		Price:price,
		Quantity:quantity,
		ImageUrl:imageUrl,
	});err!=nil{
		err = saveErr
	}
	return
}


func UpdateMerchandise(mid int64,name string,weight float64,price float64,quantity int64, imageUrl string)(err error){
	if err = dao.GetMechandisObject().UpdateMerchandise(&protocal.Merchandise{
		Mid:mid,
		Name:name,
		Weight:weight,
		Price:price,
		Quantity:quantity,
		ImageUrl:imageUrl,
	});err!=nil{
		err = updateErr
	}
	return
}

func GerMerchandiseByPrimaryKey(mid int64) (merchandise *protocal.Merchandise,err error){
	merchandise,err= dao.GetMechandisObject().GetMerchandiseByPrimaryKey(mid)
	return
}