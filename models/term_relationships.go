package models

import "github.com/astaxie/beego/orm"

type TermRelationshipsInfo struct {
	Id              uint64 `orm:"column(object_id)"`
	TermTaxonomyIds uint64 `orm:"column(term_taxonomy_id)"`
	TermOrder       int    `orm:"column(term_order)"`
}

func (t *TermRelationshipsInfo) TableName() string {
	return "term_relationships"
}

func (m *TermRelationshipsInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
