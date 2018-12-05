package protocal



type UserRequest struct {
	Email string  `json:"email"`
	Password string `json:"password"`
}



type ProjectUser struct {
	Uid   int64  `xorm:"autoincr"`
	Email string
	Password string
}

type UserRegisterReponse struct {
	Email string `json:"email"`
	Uid int64  `json:"uid"`
	Token string`json:"token"`

}
