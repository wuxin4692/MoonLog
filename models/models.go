package models

import (
	"os"
	"path"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

const(
	_DB_NAME = "data/Moonlog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Server struct {
	Id                int64
	Hostname          string
	Inner             string
	Outer             string
	Cpu_core          int64
	Men               int64
	SGroup			  string
}
type Servergroup struct {
	Id                int64
	Groupname          string
	Env                string
}
type App struct {
	Id               int64
	Appname          string
	Http_port        int64
	Conf_add         string
	Log_add          string
	Host_id          int64
	Jenkins_name     int64
	Text             string
	P_id   int64
}

type User struct {
	Id               int64
	Username         string
	Passwd           string
	E_mail           string
	P_id             string
}


func RegisterDb() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Server),new(App),new(User),new(Servergroup))
	orm.RegisterDriver(_SQLITE3_DRIVER,orm.DRSqlite)
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)

}

func AddServer(hostname string,inner string,outer string,cpu_core string,mem string,ser_group string) error {
	o := orm.NewOrm()
	cpu,err := strconv.ParseInt(cpu_core,10,64)
	if err !=nil {
		return err
	}
	memory,err1 := strconv.ParseInt(mem,10,64)
	if err1 !=nil {
		return err1
	}
	server := &Server{
		Hostname:hostname,
		Inner:inner,
		Outer:outer,
		Cpu_core:cpu,
		Men:memory,
		SGroup:ser_group,
	}
	_,ERROR := o.Insert(server)
	return ERROR
}
func AddServergroup(groupname string,env string) error {
	o := orm.NewOrm()
	servergroup := &Servergroup{
		Groupname:groupname,
		Env:env,
	}
	_,ERROR := o.Insert(servergroup)
	return ERROR
}
func DelServer(id string) error {
	cid,err := strconv.ParseInt(id,10,64)
	if err !=nil {
		return err
	}

	o := orm.NewOrm()
	cate := &Server{Id:cid}
	_,err = o.Delete(cate)
	return err
}
func DelServergroup(id string) error {
	cid,err := strconv.ParseInt(id,10,64)
	if err !=nil {
		return err
	}

	o := orm.NewOrm()
	cate := &Servergroup{Id:cid}
	_,err = o.Delete(cate)
	return err
}

func Getallserver() ([]*Server,error){
	o := orm.NewOrm()
	server := make([]*Server,0)
	qs := o.QueryTable("Server")
	_,err := qs.All(&server)
	return server,err
}
func Getallservergroup() ([]*Servergroup,error){
	o := orm.NewOrm()
	servergroup := make([]*Servergroup,0)
	qs := o.QueryTable("Servergroup")
	_,err := qs.All(&servergroup)
	return servergroup,err
}