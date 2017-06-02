package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Posts struct {
	ID                  int
	PostAuthor          int
	PostDate            time.Time `orm:"type(datetime)"`
	PostDateGmt         time.Time `orm:"type(datetime)"`
	PostContent         string    `orm:"type(longtext)"`
	PostTitle           string    `orm:"type(longtext)"`
	PostExcerpt         string    `orm:"type(longtext)"`
	PostStatus          string    `orm:"size(128)"`
	CommentStatus       string    `orm:"size(128)"`
	PingStatus          string    `orm:"size(128)"`
	PostPassword        string    `orm:"size(128)"`
	PostName            string    `orm:"size(128)"`
	ToPing              string    `orm:"type(longtext)"`
	Pinged              string    `orm:"type(longtext)"`
	PostModified        time.Time `orm:"type(datetime)"`
	PostModifiedGmt     time.Time `orm:"type(datetime)"`
	PostContentFiltered string    `orm:"type(longtext)"`
	PostParent          int
	Guid                string `orm:"size(128)"`
	MenuOrder           int
	PostType            string `orm:"size(128)"`
	PostMimeType        string `orm:"size(128)"`
	CommentCount        int
}

func init() {
	orm.RegisterModel(new(Posts))
}

func (m *Posts) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Posts) GetList() {
	var info Posts
	list := make([]*Posts, 0)
	info.Query().OrderBy("-views").Limit(5, 0).All(&list, "photo", "Id", "Title", "Name", "Addtime", "Hasepisode", "Episode")

}
