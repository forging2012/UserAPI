//base handle
//一些共用的方法
package controllers

import (
	"github.com/astaxie/beego"
	//"reflect"
	"time"
)

type baseController struct {
	beego.Controller
}

type DataResponse struct {
	Err        int         `json:"err"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	ServerTime string      `json:"servertime"`
}

func Reponse(errCode int, data interface{}, msg string) DataResponse {
	resp := DataResponse{
		Err:        errCode,
		Msg:        msg,
		Data:       data,
		ServerTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	return resp
}
