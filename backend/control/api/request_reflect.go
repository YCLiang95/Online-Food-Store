package api

import (
	"net/http"
	"reflect"
	"strconv"
	"errors"
)
var (
	MissingParamErr=errors.New("800:Missing Params")
	IllegalParamErr=errors.New("801:Parmas is illegal")
	SystemErr=errors.New("802:System error")

)

func GetStructFromRequest(r *http.Request, des interface{}) error {
	value := reflect.ValueOf(des)
	t := reflect.TypeOf(des)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		value = value.Elem()
		//key:=reformKeyName(t.Name())
		//fmt.Println(key)
		//err:=SetValue(value,r.FormValue(key))
		//if err!=nil{
		//	return err
		//}

	}
	if value.Kind()!=reflect.Struct{
         return SystemErr
	}

	fNum := t.NumField()
	for i := 0; i < fNum; i++ {
		key := reformKeyName(t.Field(i).Name)
		err := SetValue(value.Field(i), r.FormValue(key))
		if err != nil {
			return err
		}
	}


	return nil
}

func SetValue(v reflect.Value,value string)error{
	if value==""{
		return MissingParamErr
	}
	switch v.Interface().(type) {
	 case string:
		 v.SetString(value)
 	case int64,int:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return IllegalParamErr
		}
		 v.SetInt(intValue)
	case float64:
		floatValue,err:=strconv.ParseFloat(value,64)
		if err!=nil{
			return IllegalParamErr
		}
		v.SetFloat(floatValue)
	}
	return nil
}

func reformKeyName(name string) string {
	sr := []rune(name)
	sr[0] = sr[0] + 32
	return string(sr)
}
