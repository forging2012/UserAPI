//用户API相关的http Handle
//sam
//sam@6617.com
//2016-08-30
package controllers

import (
	"github.com/zituocn/UserAPI/models"
	//"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

//用户 Handle
type UserHandle struct {
	baseController
}

//Auth Key
const (
	key = "1BD6C43CA0BBF4B7ABA5E486D6A5AA2D"
)

// @Title jwt用户难测试
// @Description jwt用户验证测试，需要传入header - auth参数
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /staticblock/:key [get]
func (this *UserHandle) Auth() {
	var (
		auth string
		out  DataResponse
	)
	auth = strings.TrimSpace(this.Ctx.Request.Header.Get("auth"))
	if len(auth) == 0 {
		out = Reponse(1, "", "lost anth string")
	} else {
		token, _ := jwt.ParseWithClaims(auth, &models.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}
			return []byte(key), nil
		})
		if claims, ok := token.Claims.(*models.MyCustomClaims); ok && token.Valid {
			out = Reponse(0, claims, "")

		} else {
			out = Reponse(1, claims, "auth error")
		}
	}
	this.Data["json"] = out
	this.ServeJSON()
}

//用户登录
//POST
// /user/login/
func (this *UserHandle) Login() {
	var (
		username string
		password string
		tokenStr string
		err      error
		out      DataResponse
	)

	username = strings.TrimSpace(this.GetString("username"))
	password = strings.TrimSpace(this.GetString("password"))
	if username == "zituocn" && password == "123456" {
		expireToken := time.Now().Add(time.Hour * 24).Unix()
		claims := models.MyCustomClaims{
			username,
			password,
			jwt.StandardClaims{
				ExpiresAt: expireToken,
				Issuer:    "6617.com",
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err = token.SignedString([]byte(key))
		if err != nil {
			beego.Debug(err.Error())
		}
		out = Reponse(0, tokenStr, "")
	} else {
		out = Reponse(1, tokenStr, "username or password error")
	}

	this.Data["json"] = out
	this.ServeJSON()
}

//获取所有用户数据
//GET
// /user/
func (this *UserHandle) GetAll() {
	var (
		info models.User
	)
	list := info.GetAllUser()
	out := Reponse(0, list, "")
	this.Data["json"] = out
	this.ServeJSON()
}

//用户注册
//PUT
// /user/
func (this *UserHandle) Register() {
	this.Ctx.WriteString(this.Ctx.Request.Method + ":Register")
}

//更新某用户信息
//PUT
// /user/:id/

func (this *UserHandle) Update() {
	idstr := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString(this.Ctx.Request.Method + ":Update:" + idstr)
}

//获取某个用户的信息
//GET
// /user/:id/
func (this *UserHandle) GetOne() {
	idstr := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("GetUser:" + idstr)
}

//删除某个用户信息
//DELET
// /user/:id/
func (this *UserHandle) Delete() {
	idstr := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("Delete:" + idstr)
}
