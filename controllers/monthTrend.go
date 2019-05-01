package controllers

import (
	"hzHouse/common"
	"hzHouse/models"

	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

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
		resObj.GetErrorObj(common.ParseJSONFailed.Code, common.ParseJSONFailed.Message, err)
		this.Data["json"] = resObj
		this.ServeJSON()
		return
	}
	month := reqObj.Month

	// 查询数据库
	results, err := models.FindMonthTrend(month)
	if err != nil {
		resObj.GetErrorObj(common.SearchFailed.Code, common.SearchFailed.Message, err)
		this.Data["json"] = resObj
		this.ServeJSON()
		return
	}

	// 返回响应
	obj := common.Obj{ "monthThreadList": results }
	resObj.GetSuccessObj(common.SearchSuccess.Code, common.SearchSuccess.Message, obj)
	this.Data["json"] = resObj
	this.ServeJSON()
}