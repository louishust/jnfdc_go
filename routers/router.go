package routers

import (
	"github.com/astaxie/beego"
	"jnfdc/controllers"
)

func init() {
	beego.Router("/wx", &controllers.MainController{})
}
