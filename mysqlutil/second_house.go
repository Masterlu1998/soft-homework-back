package mysqlutil

import (
	"fmt"
)

type SecondHouse struct {
	HouseId         int     `orm:"pk;auto" json:"houseId"`
	HouseName       string  
	HouseTotalPrice int     
	HouseUnitPrice  int     
	HouseArea       int     
	HouseLat        float64 
	HouseLng        float64 
}

func init() {
}

func GetHouseList(index int, pageSize int) []SecondHouse {
	var secondHouses []SecondHouse
	qs := O.QueryTable("second_house")
	if index != 0 {
		//
		qs.Offset((index - 1) * pageSize).Limit(pageSize)
	}
	count, err := qs.All(&secondHouses)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(count)
	return secondHouses
}
