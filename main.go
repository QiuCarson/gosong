package main

import (
	_ "phpsong/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	beego.SetStaticPath("/static", "static")
	beego.Run()
}
