package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var Cfg = beego.AppConfig
var RunMode string

func init() {
	dbUser := Cfg.String("db_user")
	dbPass := Cfg.String("db_pass")
	dbHost := Cfg.String("db_host")
	dbPort := Cfg.String("db_port")
	dbName := Cfg.String("db_name")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)

	//beego.Info(dbLink)
	//orm.RegisterModel(new(Posts))
	orm.RegisterModelWithPrefix("so_", new(PostsInfo))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink)

	RunMode = Cfg.String("runmode")
	if RunMode == "dev" {
		orm.Debug = true
	}

	//orm.RegisterModelWithPrefix("so_", new(Options))
}
