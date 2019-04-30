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
	// 声明返回结构
	var resObj common.ResObj
	
	// 解析请求参数
	var reqObj struct {
		PageSize int `json:"pageSize"`
		PageIndex int `json:"pageIndex"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &reqObj)
	fmt.Println(333333)
	if err != nil {
		fmt.Println(err)
		resObj.GetErrorObj(-1, "解析请求失败", err)
		this.Data["json"] = &resObj
		this.ServeJSON()
		return
	}
	pageSize, pageIndex := reqObj.PageSize, reqObj.PageIndex

	// 调用数据库函数查询信息
	secondHouses := models.GetHouseList(pageIndex, pageSize)

	// 返回查询结果
	obj := map[string]interface{}{ "secondHouseList": secondHouses, "count": len(secondHouses) }
	resObj.GetSuccessObj(0, "查询成功", obj)
	this.Data["json"] = &resObj
	this.ServeJSON()
}
