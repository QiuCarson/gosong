package controllers

import (
	"phpsong/models"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type IndexHandle struct {
	baseController
}

func (this *IndexHandle) Start() {
	this.TplName = "index.html"
}

func (this *IndexHandle) Index() {
	var (
		page     int64
		offset   int64
		pager    string
		info     models.PostsInfo
		pagesize int64 = 10
		list     []*models.PostsInfo
	)
	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize

	cond := orm.NewCondition()
	cond1 := cond.And("post_status", "publish").Or("post_status", "publish")
	cond2 := cond.AndCond(cond1).AndCond(cond.And("post_type", "post"))
	query := info.Query().SetCond(cond2)
	//query := info.Query().Filter("post_type", "post").Filter("post_status", "publish")
	count, _ := query.Count()
	if count > 0 {
		num, _ := query.OrderBy("-post_date").Limit(pagesize, offset).All(&list)
		if num < 1 {
			this.Abort("404")
			return
		}
	}
	this.Data["list"] = list

	//推荐文章

	pager = this.PageList(pagesize, page, count, false, "")
	this.Data["pager"] = pager
	//fmt.Println(list)
	this.TplName = "index.html"
}
