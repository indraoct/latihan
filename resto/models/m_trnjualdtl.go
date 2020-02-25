package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Trnjualdtl))
}

type Trnjualdtl struct {
	Trnjualdtloid int         `orm:"pk" json:"trnjualdtloid"`
	Trnjualmst    *Trnjualmst `orm:"rel(fk);column(trnjualmstoid)" json:"trnjualmst"`
	Mstitem       *Mstitem    `orm:"rel(fk);column(mstitemoid)" json:"mstitem"`
	Note          string      `json:"note"`
}
