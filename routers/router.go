// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/zituocn/UserAPI/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/", &controllers.UserHandle{}, "get:GetAll"),
			beego.NSRouter("/", &controllers.UserHandle{}, "post:Register"),
			beego.NSRouter("/:id:int/", &controllers.UserHandle{}, "put:Update"),
			beego.NSRouter("/:id:int/", &controllers.UserHandle{}, "get:GetOne"),
			beego.NSRouter("/:id:int/", &controllers.UserHandle{}, "delete:Delete"),
			beego.NSRouter("/login/", &controllers.UserHandle{}, "post:Login"),
			beego.NSRouter("/auth/", &controllers.UserHandle{}, "get:Auth"),
		),
	)
	beego.AddNamespace(ns)
}
