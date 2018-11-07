package api

import (
	"net/http"
	"github.com/YCLiang95/CS160Group1OFS/backend/control/api/model/response"
	"github.com/YCLiang95/CS160Group1OFS/backend/service"
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
	"github.com/YCLiang95/CS160Group1OFS/backend/control/api/model/request"
	"strconv"
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

	if err = request.GetStructFromRequest(r, &mechandisRequest); err != nil {
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

func UpdateMerchandis(w http.ResponseWriter, r *http.Request) (model *response.ResponseModel, err error) {
	var (
		merchandiseUpdateRequest = request.MerchandisUpdateRequest{}
	)

	if err = request.GetStructFromRequest(r, &merchandiseUpdateRequest); err != nil {
		return
	}
	if err = service.UpdateMerchandise(merchandiseUpdateRequest.Mid,
		merchandiseUpdateRequest.Name,
		merchandiseUpdateRequest.Weight,
		merchandiseUpdateRequest.Price,
		merchandiseUpdateRequest.Quantity,
		merchandiseUpdateRequest.ImageUrl); err != nil {
		return
	}
	model = &response.ResponseModel{
		Message: "update merchandise success",
		Code:    200,
		Data:    nil,
	}
	return
}

func GetMerchandise(w http.ResponseWriter, r *http.Request) (model *response.ResponseModel, err error) {
	var (
		merchandsie *protocal.Merchandise
		mid         int
	)
	if err = r.ParseForm(); err != nil {
		return
	}

	midString := r.PostForm.Get("mid")
	if mid, err = strconv.Atoi(midString); err != nil {
		return
	}

	if merchandsie, err = service.GerMerchandiseByPrimaryKey(int64(mid)); err != nil {
		return
	}
	model = &response.ResponseModel{
		Message: "get Merchandise success",
		Code:    200,
		Data:    merchandsie,
	}
	return
}
