package models

type TermTaxonomy struct {
	Id          int    `orm:"column(term_taxonomy_id);auto"`
	TermId      uint64 `orm:"column(term_id)"`
	Taxonomy    string `orm:"column(taxonomy);size(32)"`
	Description string `orm:"column(description)"`
	Parent      uint64 `orm:"column(parent)"`
	Count       int64  `orm:"column(count)"`
}

func (t *TermTaxonomy) TableName() string {
	return "term_taxonomy"
}
