package controllers

import (
	"hzHouse/common"
	"hzHouse/hbaseutil"
	"hzHouse/models"
	
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
)

type HouseController struct {
	beego.Controller
}

func (this *HouseController) Post() {
	scanResults, err := hbaseutil.ScanHousePrice()
	var obj common.Obj
	var resObj common.ResObj
	if err != nil {
		resObj.GetErrorObj(common.SearchFailed.Code, common.SearchFailed.Message, err)
		this.Data["json"] = &resObj
		this.ServeJSON()
		return
	}
	obj = common.Obj{"resultList": scanResults}
	resObj.GetSuccessObj(common.SearchSuccess.Code, common.SearchSuccess.Message, obj)
	this.Data["json"] = &resObj
	this.ServeJSON()
}

type HouseInfoController struct {
	beego.Controller
}

func (this *HouseInfoController) Post() {
	// 声明返回结构
	var resObj common.ResObj
	
	// 解析请求参数
	var reqObj struct {
		PageSize int `json:"pageSize"`
		PageIndex int `json:"pageIndex"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &reqObj)
	if err != nil {
		fmt.Println(err)
		resObj.GetErrorObj(common.ParseJSONFailed.Code, common.ParseJSONFailed.Message, err)
		this.Data["json"] = &resObj
		this.ServeJSON()
		return
	}
	pageSize, pageIndex := reqObj.PageSize, reqObj.PageIndex

	// 调用数据库函数查询信息
	secondHouses, err := models.GetHouseList(pageIndex, pageSize)
	if err != nil {
		resObj.GetErrorObj(common.SearchFailed.Code, common.SearchFailed.Message, err)
		this.Data["json"] = resObj
		this.ServeJSON()
		return
	}

	// 返回查询结果
	obj := common.Obj{ "secondHouseList": secondHouses, "count": len(secondHouses) }
	resObj.GetSuccessObj(common.SearchSuccess.Code, common.SearchSuccess.Message, obj)
	this.Data["json"] = &resObj
	this.ServeJSON()
}
