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
	this.Data["subnav"] = info.GetPostMenu()

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
