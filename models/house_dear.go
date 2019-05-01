package models

import (
	"fmt"
)

type HouseDear struct {
	HouseId         int    `orm:"pk;auto" json:"house_id"`
	HouseAreaName   string `json:"house_area_name"`
	HouseDearMonth  int    `json:"house_dear_month"`
	HouseUnitPrice  int    `json:"house_unit_price"`
	HouseDearNumber int    `json:"house_dear_number"`
	HouseArea       int    `json:"house_area"`
}

type SearchResultInCh struct {
	Results []HouseDear
	Month int
	Err error
}

func FindDealAmountByMonth(month int, resultCh chan SearchResultInCh) {
	qs := O.QueryTable(new(HouseDear))
	var results []HouseDear
	_, err := qs.Filter("house_dear_month", month).OrderBy("house_area").All(&results)
	if err != nil {
		fmt.Println(err)
		result := SearchResultInCh{ Err: err }
		resultCh <- result
		return
	}
	result := SearchResultInCh{ Results: results, Month: month }
	resultCh <- result
}
