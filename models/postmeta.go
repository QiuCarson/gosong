package models

import (
	"strings"

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

func GetPostViews(post_id int) string {
	var (
		//info Postmeta
		view *Postmeta
	)
	//fmt.Println(post_id)
	sql := "SELECT meta_value FROM so_postmeta  WHERE post_id=? AND meta_key='views'"
	err := orm.NewOrm().Raw(sql, post_id).QueryRow(&view)
	if err == nil {
		return view.MetaValue
	}
	//info.Query().Filter("PostId", post_id).Filter("MetaKey", "views").One(&view, "MetaValue")
	//info.Query().Filter("PostId", post_id).Filter("MetaKey", "views").One(&view)

	return "0"
	//return view
}

type Menu struct {
	Name        string
	Slug        string
	Post_title  string
	Post_parent string
	Term_id     string
	Url         string
}

func (t *Postmeta) GetPostMenu() []*Menu {
	//var info Menu
	var list []*Menu
	sql := "select t.slug,t.name,p.post_title,p.post_parent,t.term_id,p.menu_order from so_postmeta pm left join so_terms t on t.term_id=pm.meta_value left join so_posts p on pm.post_id=p.ID where pm.meta_key='_menu_item_object_id' group by t.slug order by p.menu_order asc"
	//sql = "select * from so_posts where id=?"
	orm.NewOrm().Raw(sql).QueryRows(&list)
	return list
}

func GetPostMenu() []*Menu {
	//var info Menu
	var list []*Menu
	sql := "select t.name,t.slug,p.post_title,p.post_parent,t.term_id  from so_postmeta pm left join so_terms t on t.term_id=pm.meta_value left join so_posts p on pm.post_id=p.ID where pm.meta_key='_menu_item_object_id' order by p.post_parent asc, menu_order asc"
	//sql = "select * from so_posts where id=?"
	orm.NewOrm().Raw(sql).QueryRows(&list)

	return list
}

type Menu1 struct {
	Name        string
	Slug        string
	Post_title  string
	Post_parent string
	Term_id     string
	Url         string
	Active      bool
	Sub_menu    []Menu
}

func (t *Postmeta) GetMenu(currentUrl string) []Menu1 {

	list := t.GetPostMenu()
	var menu1 []Menu1
	var menu3 []Menu1
	var menu2 Menu1
	var menu4 Menu
	for _, v := range list {
		if v.Name == "" {
			v.Name = v.Post_title
		}
		if v.Post_parent == "0" {
			menu2.Name = v.Name
			menu2.Post_parent = v.Post_parent
			menu2.Slug = v.Slug
			menu2.Term_id = v.Term_id
			menu2.Post_title = v.Post_title
			menu2.Url = "/" + v.Slug
			if currentUrl != "/" {
				menu2.Active = strings.Contains("/"+v.Slug, currentUrl)
			}

			menu1 = append(menu1, menu2)
		} else {
			menu2.Name = v.Name
			menu2.Post_parent = v.Post_parent
			menu2.Slug = v.Slug
			menu2.Term_id = v.Term_id
			menu2.Post_title = v.Post_title
			menu3 = append(menu3, menu2)
		}

	}
	for k, a := range menu1 {
		for _, b := range menu3 {
			if a.Term_id == b.Post_parent {
				menu4.Name = b.Name
				menu4.Post_parent = b.Post_parent
				menu4.Slug = b.Slug
				menu4.Term_id = b.Term_id
				menu4.Post_title = b.Post_title
				menu4.Url = a.Url + "/" + b.Slug
				menu1[k].Sub_menu = append(menu1[k].Sub_menu, menu4)

			}
		}
	}

	return menu1

}
