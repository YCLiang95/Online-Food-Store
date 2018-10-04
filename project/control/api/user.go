package api

import (
	"net/http"
	"github.com/cs160/project/control/api/model/response"
	"github.com/cs160/project/control/api/model/request"
	"github.com/cs160/project/service"
)

func Register(w http.ResponseWriter, r *http.Request) (*response.ResponseModel, error) {
	userRequest := request.UserRequest{}
	if err := request.GetStructFromRequest(r, &userRequest); err != nil {
		return nil, err
	}
	if err := service.Register(userRequest.Email, userRequest.Password); err != nil {
		return nil, err
	}
	return response.GenerateSuccessStruct("register user successfully", "54895786hdfkhas"), nil
}

func Login(w http.ResponseWriter, r *http.Request) (*response.ResponseModel, error) {
	userRequest := request.UserRequest{}
	if err := request.GetStructFromRequest(r, &userRequest); err != nil {
		return nil, err
	}
	user, err := service.Login(userRequest.Email, userRequest.Password)
	if err != nil {
		return nil, err
	}
	userResponse := &response.UserRegisterReponse{
		Email: user.Email,
		Uid:   user.Uid,
		Token: "99504859hjfkaflkd",
	}
	return response.GenerateSuccessStruct("login successfully", userResponse), nil
}
