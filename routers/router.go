package routers

import (
	"hzHouse/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/get", &controllers.HouseController{})
	beego.Router("/getHouseList", &controllers.HouseInfoController{})
	beego.Router("/getMonthTrend", &controllers.MonthTrendController{})
	beego.Router("/getDealAmountPerMonth", &controllers.DealAmountController{})
	beego.Router("/getIncreaseOrder", &controllers.IncreaseOrderController{})
	beego.Router("/getPriceOrder", &controllers.PriceOrderController{})
}
