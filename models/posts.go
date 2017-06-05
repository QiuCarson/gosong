package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type PostsInfo struct {
	Id                  int       `orm:"column(ID);auto"`
	PostAuthor          uint64    `orm:"column(post_author)"`
	PostDate            time.Time `orm:"column(post_date);type(datetime)"`
	PostDateGmt         time.Time `orm:"column(post_date_gmt);type(datetime)"`
	PostContent         string    `orm:"column(post_content)"`
	PostTitle           string    `orm:"column(post_title)"`
	PostExcerpt         string    `orm:"column(post_excerpt)"`
	PostStatus          string    `orm:"column(post_status);size(20)"`
	CommentStatus       string    `orm:"column(comment_status);size(20)"`
	PingStatus          string    `orm:"column(ping_status);size(20)"`
	PostPassword        string    `orm:"column(post_password);size(20)"`
	PostName            string    `orm:"column(post_name);size(200)"`
	ToPing              string    `orm:"column(to_ping)"`
	Pinged              string    `orm:"column(pinged)"`
	PostModified        time.Time `orm:"column(post_modified);type(datetime)"`
	PostModifiedGmt     time.Time `orm:"column(post_modified_gmt);type(datetime)"`
	PostContentFiltered string    `orm:"column(post_content_filtered)"`
	PostParent          uint64    `orm:"column(post_parent)"`
	Guid                string    `orm:"column(guid);size(255)"`
	MenuOrder           int       `orm:"column(menu_order)"`
	PostType            string    `orm:"column(post_type);size(20)"`
	PostMimeType        string    `orm:"column(post_mime_type);size(100)"`
	CommentCount        int64     `orm:"column(comment_count)"`
}

func init() {
	//orm.RegisterModel(new(PostsInfo))
}
func (m *PostsInfo) TableName() string {
	return "posts"
}
func (m *PostsInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *PostsInfo) GetList(pagesize int64) []*PostsInfo {
	var info PostsInfo
	list := make([]*PostsInfo, 0)
	info.Query().OrderBy("-views").Limit(5, 0).All(&list, "ID", "post_date", "post_title", "post_content")
	return list
}

/*
//html代码过滤
func Strip_tags(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
*/
