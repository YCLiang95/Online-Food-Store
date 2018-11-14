package protocal



type UserRequest struct {
	Email string  `json:"email"`
	Password string `json:"password"`
}



type ProjectUser struct {
	Uid   int
	Email string
	Password string
}




type UserRegisterReponse struct {
	Email string `json:"email"`
	Uid int  `json:"uid"`
	Token string`json:"token"`
}
