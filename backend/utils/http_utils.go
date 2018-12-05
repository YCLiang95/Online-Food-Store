package utils

import (
	"github.com/valyala/fasthttp"
	"errors"
	"time"
	"strconv"
	"fmt"
)

const (
	POST = iota
	GET
)

func SendRequest(url string, method int, params interface{}, header map[string]string) ([]byte,error) {
	var (
		req         *fasthttp.Request
		fastArgs    *fasthttp.Args
		transResult bool
		resp   *fasthttp.Response
		responseCode int
	)

	req = &fasthttp.Request{}

	switch method {
	case GET:
		req.Header.SetMethod("GET")
		if fastArgs, transResult = params.(*fasthttp.Args); !transResult {
			return nil,errors.New("params error, please use fast args")
		}
		url = url + "?" + fastArgs.String()
		fmt.Println(url)
		break
	case POST:
		req.Header.SetMethod("POST")
		switch typeResult := params.(type) {
		case string:
			req.Header.SetContentType("application/json")
			req.SetBodyString(typeResult)
			break
		case *fasthttp.Args:
			typeResult.WriteTo(req.BodyWriter())
			break
		default:
			return nil,errors.New("params type error")
		}
		break
	default:
		return nil,errors.New("methods error")

	}

	if header != nil {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}
	req.SetRequestURI(url)
    resp = &fasthttp.Response{}
	if err:=fasthttp.DoTimeout(req,resp,time.Second*5);err!=nil{
		return nil,err
	}
	responseCode = resp.StatusCode()

	if responseCode!=200{
		return nil, errors.New("The response code is in correct: "+strconv.Itoa(responseCode))
	}

	return resp.Body(),nil
}
