package service

import (
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"github.com/YCLiang95/CS160Group1OFS/backend/dao"
)

func List() (mechandis []*protocal.Merchandise, err error) {
	mechandis, err = dao.GetMechandisObject().List()
	return
}

func Save(name string,weight float64,price float64,quantity int64, imageUrl string) (err error) {
	err = dao.GetMechandisObject().InsertNewMechanDis(&protocal.Merchandise{
		Name:name,
		Weight:weight,
		Price:price,
		Quantity:quantity,
		ImageUrl:imageUrl,
	})
	return
}
