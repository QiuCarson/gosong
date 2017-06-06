package models

import (
	"github.com/astaxie/beego/orm"
)

type Postmeta struct {
	Id        int    `orm:"column(meta_id);auto"`
	PostId    uint64 `orm:"column(post_id)"`
	MetaKey   string `orm:"column(meta_key);size(255);null"`
	MetaValue string `orm:"column(meta_value);null"`
}

func (t *Postmeta) TableName() string {
	return "postmeta"
}

func (m *Postmeta) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func GetPostViews(post_id int) *Postmeta {
	var (
		info Postmeta
		view *Postmeta
	)
	info.Query().Filter("PostId", post_id).Filter("MetaKey", "views").One(view, "Id", "PostId", "MetaKey", "MetaValue")
	return view
}

type Menu struct {
	Name       string
	Slug       string
	Post_title string
}

func (t *Postmeta) GetPostMenu() []*Menu {
	//var info Menu
	var list []*Menu
	sql := "select t.name,t.slug,p.post_title  from so_postmeta pm left join so_terms t on t.term_id=pm.meta_value left join so_posts p on pm.post_id=p.ID where pm.meta_key='_menu_item_object_id' order by menu_order asc"
	//sql = "select * from so_posts where id=?"
	orm.NewOrm().Raw(sql).QueryRows(&list)
	return list
}
