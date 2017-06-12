package routers

import (
	"phpsong/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexHandle{}, "*:Index")
	beego.Router("/:id([0-9]+).html", &controllers.IndexHandle{}, "*:Article")
	beego.Router("/tag/:tag(.*)", &controllers.IndexHandle{}, "*:TagList")
	beego.Router("/tag/:tag(.*)/page/:page([0-9]+)", &controllers.IndexHandle{}, "*:TagList")
	beego.Router("/tags", &controllers.IndexHandle{}, "*:Tags")
	beego.Router("/bookmark", &controllers.IndexHandle{}, "*:Bookmark")

	beego.Router("/page/:page([0-9]+)", &controllers.IndexHandle{}, "*:Index")

	beego.Router("/:category(.*)", &controllers.IndexHandle{}, "*:Category")
	beego.Router("/:category(.*)/page/:page([0-9]+)", &controllers.IndexHandle{}, "*:Category")

}
