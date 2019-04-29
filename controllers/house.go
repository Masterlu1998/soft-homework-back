package controllers

import (
	"hzHouse/common"
	"hzHouse/hbaseutil"
	"hzHouse/mysqlutil"

	"github.com/astaxie/beego"
)

type HouseController struct {
	beego.Controller
}

func (this *HouseController) Post() {
	scanResults, err := hbaseutil.ScanHousePrice()
	var obj map[string]interface{}
	var resObj common.ResObj
	if err != nil {
		resObj.GetErrorObj(-1, "查询失败", err)
		this.Data["json"] = &resObj
		this.ServeJSON()
		return
	}
	obj = map[string]interface{}{"resultList": scanResults}
	resObj.GetSuccessObj(0, "查询成功", obj)
	this.Data["json"] = &resObj
	this.ServeJSON()
}

type HouseInfoController struct {
	beego.Controller
}

func (this *HouseInfoController) Post() {
	secondHouses := mysqlutil.GetHouseList(0, 0)
	obj := map[string]interface{}{ "secondHouseList": secondHouses }
	var resObj common.ResObj
	resObj.GetSuccessObj(0, "查询成功", obj)
	this.Data["json"] = &resObj
	this.ServeJSON()
}
