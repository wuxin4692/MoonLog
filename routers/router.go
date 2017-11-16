package routers

import (
	"github.com/wuxin4692/MoonLog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/server",&controllers.ServerController{})
	beego.Router("/project",&controllers.ProjectController{})
	beego.Router("/servergroup",&controllers.ServergroupController{})
	beego.AutoRouter(&controllers.ServerController{})
	beego.AutoRouter(&controllers.ServergroupController{})
}
