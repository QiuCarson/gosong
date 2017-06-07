package models

import (
	"github.com/astaxie/beego/orm"
)

type TermsInfo struct {
	Id        int64  `orm:"column(term_id);auto"`
	Name      string `orm:"column(name);size(200)"`
	Slug      string `orm:"column(slug);size(200)"`
	TermGroup int64  `orm:"column(term_group)"`
}

func (t *TermsInfo) TableName() string {
	return "terms"
}

func (m *TermsInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
func GetCategoryNameByPostid(post_id int) string {
	//menu := Postmeta.GetMenu()
	info := GetCategoryTagAll(post_id)
	//var url string
	for _, v := range info {
		if info != nil && v.Taxonomy == "category" {
			/*for _, va := range menu {
				if va.Term_id == v.Term_id {
					url = va.Url
					break
				}
				for _, vb := range va.Sub_menu {
					if vb.Term_id == v.Term_id {
						url = vb.Url
						break
					}
				}
			}*/
			return v.Name
		}
	}
	return ""
}

type CategoryPost struct {
	Object_id uint64
	Name      string
	Taxonomy  string
	Term_id   string
}

func GetCategoryTagAll(post_id int) []*CategoryPost {
	CategoryPosts := make([]*CategoryPost, 0)
	sql := "SELECT  tr.object_id,t.name,tt.taxonomy,tt.term_id FROM so_terms AS t INNER JOIN so_term_taxonomy AS tt ON tt.term_id = t.term_id INNER JOIN so_term_relationships AS tr ON tr.term_taxonomy_id = tt.term_taxonomy_id  WHERE tt.taxonomy IN ('category', 'post_tag', 'post_format') AND tr.object_id =? ORDER BY t.name ASC"
	orm.NewOrm().Raw(sql, post_id).QueryRows(&CategoryPosts)
	//fmt.Println("CategoryPost", CategoryPosts)
	return CategoryPosts
}

/*
func GetPostsByCid(post_id int) []*TermRelationshipsInfo {
	var info TermRelationshipsInfo
	list := make([]*TermRelationshipsInfo, 0)
	//info.Query().Filter("Id", post_id).All(&list, "TermTaxonomyId")
	info.Query().Filter("Id", post_id).All(&list, "TermTaxonomyIds")
	return list
}
func GetCategory() []*TermsInfo {
	var info TermsInfo
	key := "category"
	list := make([]*TermsInfo, 0)
	err := GetCache(key, &list)
	if err != nil {
		info.Query().All(&list)
		SetCache(key, list)
	}
	return list
}

func GetCategoryByCid(cid int64) *TermsInfo {
	list := GetCategory()
	info := new(TermsInfo)
	for _, v := range list {
		if v.Id == cid {
			info = v
			return info
		}
	}
	return info
}

func GetCategoryNameByCid(cid int64) string {
	info := GetCategoryByCid(cid)
	if info != nil && info.Id > 0 {
		return info.Name
	}
	return ""
}
func GetCategoryNameByCids(infos []*TermRelationshipsInfo) string {

	for _, v := range infos {
		cid := int64(v.TermTaxonomyIds)
		info := GetCategoryByCid(cid)
		if info != nil && info.Id > 0 {
			return info.Name
		}

	}
	return ""
}
*/
