package routers

import (
	"phpsong/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexHandle{}, "*:Index")
	beego.Router("/page/:page", &controllers.IndexHandle{}, "*:Index")
}
