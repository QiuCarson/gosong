package routers

import (
	"phpsong/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexHandle{}, "*:Index")
	beego.Router("/page/:page", &controllers.IndexHandle{}, "*:Index")
	beego.Router("/:category", &controllers.IndexHandle{}, "*:Category")
	beego.Router("/:category/page/:page", &controllers.IndexHandle{}, "*:Category")
}
