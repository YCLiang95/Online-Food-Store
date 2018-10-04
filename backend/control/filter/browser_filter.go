package filter

import (
	"strings"
"strconv"
	"encoding/json"
	"net/http"
	"github.com/cs160/project/utils"
	"github.com/cs160/project/control/api/model/response"
)

type serverHandler func(write http.ResponseWriter, request *http.Request) (*response.ResponseModel,error)

func BrowserWapper(handler serverHandler) func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		utils.Logger.Notice("--Request URL--:", request.URL.String(), "--Request Address--:", request.RemoteAddr, "--Request Params--:", request.PostForm)
             	data,err := handler(w, request)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("content-type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
		if err != nil {
			GenerateErrResponse(w, err)
		} else {
			GenerateSuccessResponse(w, data)
		}
	}
}

func GenerateErrResponse(write http.ResponseWriter, err error) {
	ErrMessage := err.Error()
	ErrArr := strings.Split(ErrMessage, ":")
	status, _ := strconv.Atoi(ErrArr[0])
	model := &response.ResponseModel{
		Code:    status,
		Message: ErrArr[1],
		Data:    nil,
	}
	jsonByte, _ := json.Marshal(model)
	utils.Logger.Error("Request Response ERROR:", string(jsonByte))
	write.WriteHeader(200)
	write.Write(jsonByte)
}

func GenerateSuccessResponse(write http.ResponseWriter, model *response.ResponseModel) {
	resp, _ := json.Marshal(model)
	utils.Logger.Notice("Request Response SUCCESS:", string(resp))
	write.WriteHeader(200)
	write.Write(resp)
}
