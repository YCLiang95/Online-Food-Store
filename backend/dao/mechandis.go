package dao

import (
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
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
