package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct{
	beego.Controller
}

func (self *LoginController) Get() {
	isExit := self.Input().Get("exit") == "true"
	if isExit {
		self.Ctx.SetCookie("uname","",-1,"/")
		self.Ctx.SetCookie("passwd","",-1,"/")
		self.Redirect("/",301)
		return
	}
	self.TplName = "login.html"
}

func (self *LoginController) Post() {
	uname		:=	self.Input().Get("uname")
	passwd		:=	self.Input().Get("passwd")
	autologin 	:=	self.Input().Get("autologin") == "on"

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("passwd") == passwd {
			maxAge := 0
			if autologin{
				maxAge = 1 << 31 -1
			}
			self.Ctx.SetCookie("uname",uname,maxAge,"/")
			self.Ctx.SetCookie("passwd",passwd,maxAge,"/")
	}
	self.Redirect("/",301)
	return
}
//cookie 读取
func checkAccount(ctx *context.Context) bool {
	ck,err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value
	ck2,err := ctx.Request.Cookie("passwd")
	if err != nil {
		return false
	}
	passwd := ck2.Value
	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("passwd") == passwd
}