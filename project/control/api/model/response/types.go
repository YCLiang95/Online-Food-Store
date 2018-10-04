package response


type ResponseModel struct {
	Message string  `json:"message"`
	Data  interface{} `json:"data"`
	Code int   `json:"code"`
}

type PageResponse struct {
	Total int     `json:"total"`
	Current int    `json:"current"`
	NextPage int  `json:"nextPage"`
	PrePage int    `json:"Prepage"`
	Data  interface{}   `json:"data"`
}

func GenerateSuccessStruct(message string,data interface{})*ResponseModel{
	return &ResponseModel{
		Data:data,
		Message:message,
		Code:200,
	}
}


