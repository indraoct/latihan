package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Trnjualmst))
}

type Trnjualmst struct {
	Trnjualmstoid int    `orm:"pk" json:"trnjualmstoid"`
	Trnjualmstno  string `json:"trnjualmstno"`
}
