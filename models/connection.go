package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

var (
	O orm.Ormer
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:62795828lovE@tcp(116.62.156.102:3306)/house_price?charset=utf8", 30)
	orm.RegisterModel(new(SecondHouse), new(MonthPrice))
	O = orm.NewOrm()
	O.Using("default")
}
