package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wuxin4692/MoonLog/models"
)

type ServerController struct {
	beego.Controller
}

func (self *ServerController) Get() {
	op := self.Input().Get("op")

	switch op {
	case "add":
		self.Redirect("/server",301)
		return
	case "del":
		id:=self.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelServer(id)
		if err != nil {
			beego.Error(err)
		}
		self.Redirect("/server",301)
		return
	}
	self.Data["IsServer"] = true
	self.TplName = "server.html"

	var err error
	self.Data["Server"],err = models.Getallserver()
	if err != nil {
		beego.Error(err)
	}
}
func (self *ServerController) Post() {
	hostname:=self.Input().Get("hostname")
	inner:=self.Input().Get("inner")
	outer:=self.Input().Get("outer")
	cpu_core:=self.Input().Get("cpu_core")
	mem:=self.Input().Get("mem")
	server_group := self.Input().Get("server_group")


	var err error
	err = models.AddServer(hostname,inner,outer,cpu_core,mem,server_group)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/server",302)
}

func (self *ServerController) Add() {
	var err1 error
	self.Data["Sgroup"],err1 = models.Getallservergroup()
	if err1 != nil {
		beego.Error(err1)
	}
	self.TplName = "add_server.html"
}
