package controllers

import (
	"hzHouse/common"
	"hzHouse/models"

	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
)

type PriceOrderController struct {
	beego.Controller
}

var (
	mysqlDescCode          = -1 // 倒序排序
	mysqlAscCode           = 1  // 正序排序
	invalidParamMessageStr = "invalid param message!"
)

func (this *PriceOrderController) Post() {
	// 声明返回结构
	var resObj common.ResObj

	// 解析请求
	var reqObj struct {
		ShowType int
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &reqObj)
	if err != nil {
		resObj.GetErrorObj(common.ParseJSONFailed.Code, common.ParseJSONFailed.Message, err)
		this.Data["json"] = resObj
		this.ServeJSON()
		return
	}
	showType := reqObj.ShowType

	var results interface{}
	switch showType {
	case 1:
		// 每平方米房价排名
		results, err = models.GetDealAmountListInNov()
		if err != nil {
			resObj.GetErrorObj(common.SearchFailed.Code, common.SearchFailed.Message, err)
			this.Data["json"] = resObj
			this.ServeJSON()
			return
		}
	case 2:
		// 房子总价倒序排名
		results, err = models.GetHouseOrderList(mysqlDescCode)
		if err != nil {
			resObj.GetErrorObj(common.SearchFailed.Code, common.SearchFailed.Message, err)
			this.Data["json"] = resObj
			this.ServeJSON()
			return
		}
	case 3:
		// 房子总价正序排名
		results, err = models.GetHouseOrderList(mysqlAscCode)
		if err != nil {
			resObj.GetErrorObj(common.SearchFailed.Code, common.SearchFailed.Message, err)
			this.Data["json"] = resObj
			this.ServeJSON()
			return
		}
	default:
		resObj.GetErrorObj(common.ReqParamInvalid.Code, common.ReqParamInvalid.Message, errors.New(invalidParamMessageStr))
		this.Data["json"] = resObj
		this.ServeJSON()
		return
	}

	// 返回响应
	obj := common.Obj{"house_oder_list": results}
	resObj.GetSuccessObj(common.SearchSuccess.Code, common.SearchSuccess.Message, obj)
	this.Data["json"] = resObj
	this.ServeJSON()
}
