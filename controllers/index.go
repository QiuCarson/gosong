package controllers

import (
	"phpsong/models"
	"strconv"

	"fmt"

	"github.com/astaxie/beego/orm"
)

type IndexHandle struct {
	baseController
}

func (this *IndexHandle) Start() {
	this.TplName = "index.html"
}

//博客首页
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
	var infos models.PostsInfo
	infos.GetTop()
	/*top := info.GetTop()
	this.Data["top"] = top*/
	//友情链接

	pager = this.PageList(pagesize, page, count, false, "")
	this.Data["pager"] = pager
	//fmt.Println(list)
	this.TplName = "index.html"
}

//博客栏目页
func (this *IndexHandle) Category() {

	var (
		page         int64
		offset       int64
		count        int64
		pager        string
		CategoryName string
		info         models.PostsInfo
		pagesize     int64 = 10
		list         []*models.PostsInfo
	)
	categorystr := this.Ctx.Input.Param(":category")

	pagestr := this.Ctx.Input.Param(":page")

	page, _ = strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize

	CategoryName, count, list = info.GetCategoryPosts(categorystr, offset, pagesize)
	if len(list) < 1 {
		this.Abort("404")
		return
	}

	this.Data["list"] = list
	this.Data["categoryName"] = CategoryName

	pager = this.PageList(pagesize, page, count, false, categorystr)
	this.Data["pager"] = pager
	this.TplName = "list.html"
}

//博客文章页
func (this *IndexHandle) Article() {
	var (
		info    models.PostsInfo
		article models.PostsInfo
		meta    models.Postmeta
		id      int64
		num     int64
		err     error
	)
	idstr := this.Ctx.Input.Param(":id")
	id, err = strconv.ParseInt(idstr, 10, 64)

	if err != nil || id <= 0 {
		this.Abort("404")
		return
	}

	//读取数据
	err = info.Query().Filter("Id", id).One(&article)
	if err == orm.ErrNoRows {
		this.Abort("404")
		return
	}
	this.Data["article"] = article

	//更新查看次数
	num, err = meta.Query().Filter("PostId", id).Filter("MetaKey", "views").Update(orm.Params{"MetaValue": orm.ColValue(orm.ColAdd, 1)})
	fmt.Println(num)
	this.TplName = "article.html"
}

//博客TAG页
func (this *IndexHandle) Tags() {
	var (
		info models.TermsInfo
	)
	list := info.GetAllTags()
	this.Data["tags"] = list
	this.TplName = "tags.html"
}
