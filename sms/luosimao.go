package sms

import (
	"fmt"
	"strings"
	"github.com/spf13/viper"

	"os"
	"io"
	"net/http"
	"microservice/utils"
)

func Send(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := utils.GetUrlParam("mobile", request)
	message := utils.GetUrlParam("message", request)

	respCode := sendLSM(mobile, message)

	fmt.Fprint(writer, utils.RespOK("短信发送状态", respCode))
}

func sendLSM(mobile string, message string) int{


	url := "https://sms-api.luosimao.com/v1/send.json"
	fmt.Println(url)
	client := &http.Client{}
	body := strings.NewReader("mobile=" + mobile +"&" + "message=" + message)
	fmt.Println(body)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded;charset=UTF-8")
	req.SetBasicAuth("api", "key-"+viper.GetString("sms.luosimao.api_key"))
	utils.CheckErr(err)
	resp, _ := client.Do(req)
	stdout := os.Stdout
	_, err = io.Copy(stdout, resp.Body)

	//返回的状态码
	status := resp.StatusCode
	return status
}
