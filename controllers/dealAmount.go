package controllers

import (
	"hzHouse/common"
	"hzHouse/models"

	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type DealAmountController struct {
	beego.Controller
}

func (this *DealAmountController) Post() {
	// 声明返回结构
	var resObj common.ResObj

	// 解析请求参数
	var reqObj struct {
		StartMonth int
		EndMonth   int
	}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &reqObj); err != nil {
		resObj.GetErrorObj(common.ParseJSONFailed.Code, common.ParseJSONFailed.Message, err)
		this.Data["json"] = resObj
		this.ServeJSON()
		return
	}

	// 数据库查询
	resultCh := make(chan models.SearchResultInCh)
	for i := reqObj.StartMonth; i <= reqObj.EndMonth; i++ {
		go models.FindDealAmountByMonth(i, resultCh)
	}

	// 返回结果
	obj := common.Obj{}
	count := reqObj.EndMonth - reqObj.StartMonth + 1
	for val := range resultCh {
		count--
		obj[strconv.Itoa(val.Month)] = val.Results
		if count == 0 {
			break
		}

		// 判断是否有错误
		if val.Err != nil {
			resObj.GetErrorObj(common.SearchFailed.Code, common.SearchFailed.Message, val.Err)
			this.Data["json"] = resObj
			this.ServeJSON()
			return
		}
	}
	resObj.GetSuccessObj(common.SearchSuccess.Code, common.SearchSuccess.Message, obj)
	this.Data["json"] = resObj
	this.ServeJSON()
}
