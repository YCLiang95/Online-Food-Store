package api

import (
	"net/http"

	"github.com/YCLiang95/CS160Group1OFS/backend/service"

	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
)

func Register(w http.ResponseWriter, r *http.Request) (*protocal.ResponseModel, error) {
	userRequest := protocal.UserRequest{}
	if err := GetStructFromRequest(r, &userRequest); err != nil {
		return nil, err
	}
	if err := service.Register(userRequest.Email, userRequest.Password); err != nil {
		return nil, err
	}
	return protocal.GenerateSuccessStruct("register user successfully", "54895786hdfkhas"), nil
}

func Login(w http.ResponseWriter, r *http.Request) (*protocal.ResponseModel, error) {
	userRequest := protocal.UserRequest{}
	if err := GetStructFromRequest(r, &userRequest); err != nil {
		return nil, err
	}

	//err := r.ParseForm()
	//if err != nil {
	//	return nil, errors.New("5001:系统错误")
	//}
	//
	//email := r.PostForm.Get("email")
	//password := r.PostForm.Get("password")
	//
	//if email==""||password==""{
	//	return nil, errors.New("5001:系统错误")
	//}

	user, err := service.Login(userRequest.Email, userRequest.Password)
	if err != nil {
		return nil, err
	}
	userResponse := &protocal.UserRegisterReponse{
		Email: user.Email,
		Uid:   user.Uid,
		Token: "99504859hjfkaflkd",
	}
	return protocal.GenerateSuccessStruct("login successfully", userResponse), nil
}
