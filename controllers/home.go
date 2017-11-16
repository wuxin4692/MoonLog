package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (self *MainController) Get() {
	self.Data["IsHome"] = true
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	self.TplName = "home.html"
}
