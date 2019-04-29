package routers

import (
	"hzHouse/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/get", &controllers.HouseController{})
}
