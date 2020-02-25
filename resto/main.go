package main

import (
	_ "latihan/resto/routers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// Set database
	setDatabase()
	beego.BConfig.WebConfig.AutoRender = false
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func setDatabase() {
	// set default database
	configDb := beego.AppConfig.String("mysql_user") + ":" + beego.AppConfig.String("mysql_pass") + "@tcp(" +
		beego.AppConfig.String("mysql_url") + ")/" + beego.AppConfig.String("mysql_db") + "?charset=utf8&parseTime=true&loc=Asia%2fJakarta"

	orm.RegisterDriver("mysql", orm.DRMySQL)

	// MaxIdleConns = 0, SetMaxOpenConns = 2400
	err := orm.RegisterDataBase("default", "mysql", configDb, 0, 2400)

	// Connection Setting
	db, err := orm.GetDB("default")
	if err == nil {
		db.SetConnMaxLifetime(time.Second * 10)
	} else {
		beego.Error(err)
	}

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
}
