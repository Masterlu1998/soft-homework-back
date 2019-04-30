package models

import (
	"fmt"
)

type SecondHouse struct {
	HouseId         int     `orm:"pk;auto"`
	HouseName       string  
	HouseTotalPrice int     
	HouseUnitPrice  int     
	HouseArea       int     
	HouseLat        float64 
	HouseLng        float64 
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
