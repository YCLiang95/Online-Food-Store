package response

type UserRegisterReponse struct {
	Email string `json:"email"`
	Uid int  `json:"uid"`
	Token string`json:"token"`
}
