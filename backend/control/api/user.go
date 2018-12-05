package api

import (
	"net/http"

	"github.com/YCLiang95/CS160Group1OFS/backend/service"

	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
)

func Register(w http.ResponseWriter, r *http.Request) (response *protocal.ResponseModel, err error) {

	userRequest := protocal.UserRequest{}
	if err = GetStructFromRequest(r, &userRequest); err != nil {
		return
	}
	if err = service.Register(userRequest.Email, userRequest.Password); err != nil {
		return
	}

	return protocal.GenerateSuccessStruct("register user successfully", nil), nil
}

func Login(w http.ResponseWriter, r *http.Request) (*protocal.ResponseModel, error) {
	var (
		token string
	)
	userRequest := protocal.UserRequest{}
	if err := GetStructFromRequest(r, &userRequest); err != nil {
		return nil, err
	}
	user, err := service.Login(userRequest.Email, userRequest.Password)
	if err != nil {
		return nil, err
	}

	userResponse := &protocal.UserRegisterReponse{
		Email: user.Email,
		Uid:   user.Uid,
		Token: token,
	}

	return protocal.GenerateSuccessStruct("login successfully", userResponse), nil
}
