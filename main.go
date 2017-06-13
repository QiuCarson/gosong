package main

import (
	"phpsong/controllers"
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
	beego.AddFuncMap("GetPostImgByPostId", models.GetPostImgByPostId)
	beego.AddFuncMap("GetPostViews", models.GetPostViews)

	//models.GetPostViews(1)
	//beego.SetLevel(beego.LevelError)
	//models.GetPostImgByPostId(2775)
	//models.GetPostImgByPostId(1)
	//dmodels.GetPostImgByPostId(2751)
	models.PostIdByComment()
	beego.ErrorController(&controllers.ErrorController{})
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
