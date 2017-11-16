package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wuxin4692/MoonLog/models"
)

type ServergroupController struct {
	beego.Controller
}

func (self *ServergroupController) Get() {
	op := self.Input().Get("op")

	switch op {
	case "add":
		self.Redirect("/servergroup",301)
		return
	case "del":
		id:=self.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err1 := models.DelServergroup(id)
		if err1 != nil {
			beego.Error(err1)
		}
		self.Redirect("/servergroup",301)
		return
	}
	self.Data["IsServergroup"] = true
	self.TplName = "server_group.html"

	var err error
	self.Data["Servergroup"],err = models.Getallservergroup()
	if err != nil {
		beego.Error(err)
	}
}
func (self *ServergroupController) Post() {
	groupname:=self.Input().Get("groupname")
	env:=self.Input().Get("env")

	var err error
	err = models.AddServergroup(groupname,env)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/servergroup",302)
}

func (self *ServergroupController) Add() {
	self.TplName = "add_servergroup.html"
}
