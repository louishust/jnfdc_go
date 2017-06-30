package routers

import (
	"jnfdc/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/jnfdc",
			beego.NSInclude(
				&controllers.JnfdcController{},
			),
		),
	)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Content-type", "Origin", "X-whatever", "X-CaseSensitive", "token", "sign", "ts"},
	}))

	beego.AddNamespace(ns)
}
