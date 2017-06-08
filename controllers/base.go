package controllers

import (
	"phpsong/models"
	"strconv"

	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
}

func (this *baseController) Prepare() {
	var info models.Postmeta
	//var option models.OptionsInfo
	//this.Data["subnav"] = info.GetPostMenu()
	url := this.Ctx.Request.RequestURI
	this.Data["subnav"] = info.GetMenu(url)

	var options models.OptionsInfo
	options.AutolaodOption()
	this.Data["blogname"] = models.OptionMap["blogname"]
	this.Data["blogdescription"] = models.OptionMap["blogdescription"]

	//var m map[string]string
	//option.AutolaodOption()
	//this.Data["blogname"] = m["blogname"]
	//path := this.Ctx.Request.URL.String()
	//this.Data["currentUrl"] = this.Ctx.Request.RequestURI

}

//显示分页链接
func (this *baseController) PageList(pagesize, page, recordcount int64, first bool, path string) (pager string) {
	if recordcount == 0 {
		return ""
	}

	var pagecount int64
	pagecount = 0

	if recordcount%pagesize == 0 {
		pagecount = recordcount / pagesize
	} else {
		pagecount = (recordcount / pagesize) + 1
	}

	if pagecount < 2 {
		return "<span>共1页</span>"
	}
	if page > pagecount {
		return ""
	}

	pager = "<nav class='pagination' style='display: none;'><ul>"

	if page < pagecount {
		pager += "<li class='next-page'><a href='" + path + "/page/" + strconv.FormatInt(page+1, 10) + "'>下一页</a></li>"
	}
	pager += "</ul> </nav>"
	return pager

}
