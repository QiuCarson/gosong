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
