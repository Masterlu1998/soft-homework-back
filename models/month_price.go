package models

import (
	"fmt"
)

type MonthPrice struct {
	Month int `orm:"pk;auto" json:"month"`
	UnitPrice int `json:"unit_price"`
}

func FindMonthTrend(month int) ([]MonthPrice, error) {
	var searchResult []MonthPrice
	var err error

	qs := O.QueryTable(new(MonthPrice))
	fmt.Println(month)
	if month == 0 {
		_, err = qs.All(&searchResult)
	} else {
		_, err = qs.Filter("month", month).All(&searchResult)
	}
	return searchResult, err
}