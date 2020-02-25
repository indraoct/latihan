package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Mstitem))
}

type Mstitem struct {
	Mstitemoid int    `orm:"pk" json:"mstitemoid"`
	Itemdesc   string `json:"itemdesc"`
}
