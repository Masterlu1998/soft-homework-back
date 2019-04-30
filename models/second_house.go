package models

import (
	"fmt"
)

type SecondHouse struct {
	HouseId         int     `orm:"pk;auto" json:"house_id"`
	HouseName       string  `json:"house_name"`
	HouseTotalPrice int     `json:"house_total_price"`
	HouseUnitPrice  int     `json:"house_unit_price"`
	HouseArea       int     `json:"house_area"`
	HouseLat        float64 `json:"house_lat"`
	HouseLng        float64 `json:"house_lng"`
}

func GetHouseList(index int, pageSize int) []SecondHouse {
	var secondHouses []SecondHouse
	var count int64
	var err error
	qs := O.QueryTable(new(SecondHouse))
	if index != 0 {
		println(3333)
		count, err = qs.Offset((index - 1) * pageSize).Limit(pageSize).All(&secondHouses)
	} else {
		count, err = qs.All(&secondHouses)
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(count)
	return secondHouses
}
