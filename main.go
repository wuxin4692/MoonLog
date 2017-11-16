package main

import (
	_ "MoonLog/routers"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"MoonLog/models"
)
func init() {
	models.RegisterDb()
}

func main() {
	orm.Debug =true
	orm.RunSyncdb("default",false,true)
	beego.Run()
}

