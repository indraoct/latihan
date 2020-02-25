package controllers

import (
	"latihan/resto/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DataController struct {
	beego.Controller
}

// @router /list [get]
func (this *DataController) List() {

	var arr_trx_detail []models.Trnjualdtl
	o := orm.NewOrm()

	o.QueryTable("Trnjualdtl").
		RelatedSel("Trnjualmst").
		RelatedSel("Mstitem").
		All(&arr_trx_detail)

	this.Data["json"] = arr_trx_detail
	this.ServeJSON()
}
