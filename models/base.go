package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var Cfg = beego.AppConfig
var RunMode string
var DbPrefix string

//var FileLogs *logs.BeeLogger

func init() {
	dbUser := Cfg.String("db_user")
	dbPass := Cfg.String("db_pass")
	dbHost := Cfg.String("db_host")
	dbPort := Cfg.String("db_port")
	dbName := Cfg.String("db_name")
	DbPrefix := Cfg.String("db_prefix")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)

	//beego.Info(dbLink)
	//orm.RegisterModel(new(Posts))
	orm.RegisterModelWithPrefix(DbPrefix, new(PostsInfo), new(TermsInfo), new(TermRelationshipsInfo), new(TermTaxonomy), new(Postmeta), new(OptionsInfo), new(CommentsInfo))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink)
	beego.SetLevel(beego.LevelError)
	RunMode = Cfg.String("runmode")
	if RunMode == "dev" || RunMode == "prod" {
		orm.Debug = true
	}
	//FileLogs = logs.NewLogger(1000)
	//FileLogs.SetLogger("file", `{"filename":”logs/test.log"}`)

	//orm.RegisterModelWithPrefix("so_", new(Options))
}
