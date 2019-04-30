package controllers

import (
	"hzHouse/common"
	"hzHouse/models"

	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type Obj map[string]interface{}

type MonthTrendController struct {
	beego.Controller
}

func (this *MonthTrendController) Post() {
	// 声明返回结构
	var resObj common.ResObj

	// 解析请求参数
	var reqObj struct {
		Month int
	}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &reqObj); err != nil {
		fmt.Println(err)
		resObj.GetErrorObj(-1, "解析请求失败", err)
		this.Data["json"] = resObj
		this.ServeJSON()
		return
	}
	month := reqObj.Month

	// 查询数据库
	results := models.FindMonthTrend(month)
	obj := Obj{ "monthThreadList": results }
	resObj.GetSuccessObj(0, "查询成功", obj)
	this.Data["json"] = resObj
	this.ServeJSON()
}