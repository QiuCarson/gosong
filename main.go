package main

import (
	"phpsong/models"
	_ "phpsong/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//beego.AddFuncMap("strip_tags", models.Strip_tags)
	//beego.AddFuncMap("subString", models.SubString)

	//beego.AddFuncMap("GetCategoryNameByCid", models.GetCategoryNameByCid)
	beego.AddFuncMap("GetCategoryNameByPostid", models.GetCategoryNameByPostid)
	//models.GetPostViews(1)
	//beego.SetLevel(beego.LevelError)
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
