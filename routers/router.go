package routers

import (
	"github.com/shwinpiocess/gmps/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
