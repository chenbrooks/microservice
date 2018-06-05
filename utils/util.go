package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetUrlParam(paramName string, r *http.Request) string {
	param := ""
	paramValue, found := r.Form[paramName]
	if found {
		if len(paramValue) > 0 {
			param = paramValue[0]
		}
	}
	return param
}

func GetRouterParamStr(paramName string, r *http.Request) string {
	vars := mux.Vars(r)
	paramValue := vars[paramName]
	fmt.Println("param: " + paramValue)
	return paramValue
}

func GetRouterParamInt(paramName string, r *http.Request) int {
	vars := mux.Vars(r)
	param, err := strconv.Atoi(vars[paramName])
	CheckErr(err)
	return param
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (resp *Resp) ToJsonStr() []byte {
	jsonStr, err := json.Marshal(resp)
	CheckErr(err)
	return jsonStr
}

func RespOK(msg string, data interface{}) string {
	resp := &Resp{1, msg, data}
	return string(resp.ToJsonStr())
}
