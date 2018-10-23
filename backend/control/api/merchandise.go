package api

import (
	"net/http"
	"github.com/YCLiang95/CS160Group1OFS/backend/control/api/model/response"
	"github.com/YCLiang95/CS160Group1OFS/backend/service"
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"github.com/YCLiang95/CS160Group1OFS/backend/control/api/model/request"
)

func List(w http.ResponseWriter, r *http.Request) (model *response.ResponseModel, err error) {
	var (
		mechandis []*protocal.Merchandise
	)
	if mechandis, err = service.List(); err != nil {
		model = nil
		return
	}
	model = &response.ResponseModel{
		Message: "get merchandis list success",
		Code:    200,
		Data:    mechandis,
	}
	return
}


func SaveMerchandis(w http.ResponseWriter, r *http.Request) (model *response.ResponseModel, err error) {
	var (
	   mechandisRequest = request.MerchandisRequest{}

	)

	if err=request.GetStructFromRequest(r,&mechandisRequest);err!=nil{
		return
	}

	if err = service.Save(mechandisRequest.Name,
		mechandisRequest.Weight,
		mechandisRequest.Price,
		mechandisRequest.Quantity,
		mechandisRequest.ImageUrl); err != nil {
		model = nil
		return
	}
	model = &response.ResponseModel{
		Message: "save merchandis  success",
		Code:    200,
		Data:    nil,
	}
	return
}