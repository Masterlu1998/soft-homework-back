package models

type SecondHouse struct {
	HouseId         int     `orm:"pk;auto" json:"house_id"`
	HouseName       string  `json:"house_name"`
	HouseTotalPrice int     `json:"house_total_price"`
	HouseUnitPrice  int     `json:"house_unit_price"`
	HouseArea       int     `json:"house_area"`
	HouseLat        float64 `json:"house_lat"`
	HouseLng        float64 `json:"house_lng"`
}

func GetHouseList(index int, pageSize int) ([]SecondHouse, error) {
	var secondHouses []SecondHouse
	var err error
	qs := O.QueryTable(new(SecondHouse))
	if index != 0 {
		println(3333)
		_, err = qs.Limit(pageSize, (index - 1) * pageSize).All(&secondHouses)
	} else {
		_, err = O.Raw("SELECT * FROM second_house").QueryRows(&secondHouses)
	}
	return secondHouses, err
}

func GetHouseOrderList(order int) ([]SecondHouse, error) {
	qs := O.QueryTable(new(SecondHouse))
	var results []SecondHouse
	var err error
	if order == -1 {
		_, err = qs.OrderBy("-house_unit_price").All(&results)
	} else {
		_, err = qs.OrderBy("house_unit_price").All(&results)
	}
	if err != nil {
		return nil, err
	}
	return results, nil
}
