package filter

import (
	"strings"
	"strconv"
	"encoding/json"
	"net/http"
	"github.com/YCLiang95/CS160Group1OFS/backend/utils"
	"time"
	"fmt"
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
)

type serverHandler func(write http.ResponseWriter, request *http.Request) (*protocal.ResponseModel, error)

func BrowserWapper(handler serverHandler) func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		data, err := handler(w, request)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("content-type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
		endTime := time.Now()
		usedTime := endTime.Sub(startTime).Seconds()

		if err != nil {
			GenerateErrResponse(w, err, request, usedTime)
		} else {
			GenerateSuccessResponse(w, data, request, usedTime)
		}
	}
}

func GenerateErrResponse(write http.ResponseWriter, err error, request *http.Request, timeUsed float64) {
	ErrMessage := err.Error()
	ErrArr := strings.Split(ErrMessage, ":")
	status, _ := strconv.Atoi(ErrArr[0])
	model := &protocal.ResponseModel{
		Code:    status,
		Message: ErrArr[1],
		Data:    nil,
	}
	jsonByte, _ := json.Marshal(model)
	utils.Logger.Error(fmt.Sprintf(`
--------Response ERROR--------

Request URL:%v,

Request Remote address: %v,

Reqeuest UsedTime %v s,

Request Params: %v,

Request ERROR %v

---------------------------------

`, request.RequestURI, request.RemoteAddr, timeUsed, request.PostForm, err))
	write.WriteHeader(200)
	write.Write(jsonByte)
}

func GenerateSuccessResponse(write http.ResponseWriter, model *protocal.ResponseModel, request *http.Request, timeUsed float64) {
	resp, _ := json.Marshal(model)
	utils.Logger.Notice(fmt.Sprintf(`
--------Response Success--------

Request URL:%v,

Request Remote address: %v,

Reqeuest UsedTime %v s,

Request Params: %v,

Response data: %v

---------------------------------

`, request.RequestURI, request.RemoteAddr, timeUsed, request.PostForm, string(resp)))
	write.WriteHeader(200)
	write.Write(resp)
}
