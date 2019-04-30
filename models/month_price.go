package models

import (
	"fmt"
)

type MonthPrice struct {
	Month int `orm:"pk;auto"`
	UnitPrice int 
}

func FindMonthTrend(month int) []MonthPrice {
	var searchResult []MonthPrice
	qs := O.QueryTable(new(MonthPrice))
	fmt.Println(month)
	if month == 0 {
		qs.All(&searchResult)
	} else {
		qs.Filter("month", month).All(&searchResult)
	}
	return searchResult
}