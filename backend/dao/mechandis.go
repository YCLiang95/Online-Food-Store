package dao

import (
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"strconv"
	"strings"
)

type MechandisDao struct{}

var mechandisDao *MechandisDao = nil

func GetMechandisObject() *MechandisDao {
	if mechandisDao == nil {
		mechandisDao = new(MechandisDao)
	}
	return mechandisDao
}

func (md *MechandisDao) List() (mechandis []*protocal.Merchandise, err error) {

	mechandis = make([]*protocal.Merchandise, 0)

	err = daoFunctionLogWapper(&mechandis, nil, findRecorders);
	return
}

func (md *MechandisDao) InsertNewMechanDis(merchandis *protocal.Merchandise) (err error) {
	err = daoFunctionLogWapper(merchandis, nil, saveRecorders)
	return
}

func (md *MechandisDao) UpdateMerchandise(merchandise *protocal.Merchandise) (err error) {
	err = daoFunctionLogWapper(merchandise, &mysqlOptions{Filter: map[string]interface{}{"mid=?": merchandise.Mid}}, updateRecorder)
	return
}

func (nd *MechandisDao) GetMerchandiseByPrimaryKey(mid int64) (merchandis *protocal.Merchandise, err error) {
	merchandis = &protocal.Merchandise{Mid: mid}
	err = daoFunctionLogWapper(merchandis, nil, getRecorder)
	return
}


func (nd *MechandisDao) GetMerchandises(orderDetail []protocal.OrderReequestDetail) (merchandises map[int64]*protocal.Merchandise, err error) {

	condition:="`mid` IN ("

	midStrings:=make([]string,len(orderDetail))

	for i:=0;i<len(midStrings);i++{
		midStrings[i] = strconv.Itoa(int(orderDetail[i].Mid))
	}
	condition+=strings.Join(midStrings,",")

	condition+=")"

	merchandises = make(map[int64]*protocal.Merchandise)
	err = daoFunctionLogWapper(&merchandises, &mysqlOptions{Where:condition}, findRecorders)
	return
}


