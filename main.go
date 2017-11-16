package main

import (
	_ "github.com/wuxin4692/MoonLog/routers"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/wuxin4692/MoonLog/models"
)
func init() {
	models.RegisterDb()
}

func main() {
	orm.Debug =true
	orm.RunSyncdb("default",false,true)
	beego.Run()
}

