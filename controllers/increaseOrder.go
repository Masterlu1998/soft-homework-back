package controllers

import (
	"hzHouse/common"
	"hzHouse/models"

	"strconv"

	"github.com/astaxie/beego"
)

type IncreaseOrderController struct {
	beego.Controller
}

func (this *IncreaseOrderController) Post() {
	// 声明返回数据格式
	var resObj common.ResObj

	// 请求数据库
	resultCh := make(chan models.SearchResultInCh)
	go models.FindDealAmountByMonth(9, resultCh)
	go models.FindDealAmountByMonth(11, resultCh)
	count := 2
	groupSearchObj := make(map[string][]models.HouseDear)
	for val := range resultCh {
		count--
		groupSearchObj[strconv.Itoa(val.Month)] = val.Results
		if val.Err != nil {
			resObj.GetErrorObj(common.SearchFailed.Code, common.SearchFailed.Message, val.Err)
			this.Data["json"] = resObj
			this.ServeJSON()
			return
		}
		if count == 0 {
			break
		}
	}

	// 分析数据
	sepData := groupSearchObj["9"]
	novData := groupSearchObj["11"]
	type finalResult struct {
		HouseArea       int     `json:"house_area"`
		DealNumIncrease float64 `json:"deal_num_increase"`
		PriceIncrease   float64 `json:"price_increase"`
	}
	finalResults := make([]finalResult, 0, 10)
	for houseArea := 0; houseArea < 10; houseArea++ {
		dealNumIncrease := common.CalculateIncrement(sepData[houseArea].HouseDearNumber, novData[houseArea].HouseDearNumber)
		priceIncrease := common.CalculateIncrement(sepData[houseArea].HouseUnitPrice, novData[houseArea].HouseUnitPrice)
		item := finalResult{HouseArea: houseArea, DealNumIncrease: dealNumIncrease, PriceIncrease: priceIncrease}
		finalResults = append(finalResults, item)
	}

	// 返回响应
	obj := common.Obj{"increaseList": finalResults}
	resObj.GetSuccessObj(common.SearchSuccess.Code, common.SearchSuccess.Message, obj)
	this.Data["json"] = resObj
	this.ServeJSON()
}
