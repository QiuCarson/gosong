package models

import (
	"regexp"
	"strings"
	"time"

	"phpsong/phpserialize"

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

type Img struct {
	Meta_value string
}
type msg interface{}

func GetPostImgByPostId(post_id int, PostContent string) string {
	var imgs Img
	sql := "SELECT pt.meta_value FROM `so_posts` p Inner join `so_postmeta` pt  on p.ID=pt.post_id WHERE pt.meta_key='_so_attachment_metadata' and p.post_type='attachment' and p.post_parent=? limit 1"
	//sql = "select * from so_posts where id=?"
	err := orm.NewOrm().Raw(sql, post_id).QueryRow(&imgs)
	flag_bool := true
	if err == nil && false {

		var decodeRes interface{}
		var ok error
		decodeRes, ok = phpserialize.Decode(imgs.Meta_value)

		if ok == nil {

			decodeData, _ := decodeRes.(map[interface{}]interface{})

			file := decodeData["file"]
			post_thumbnail := decodeData["sizes"].(map[interface{}]interface{})
			post_thumbnail1 := post_thumbnail["post-thumbnail"].(map[interface{}]interface{})

			file_string := post_thumbnail1["file"].(string)
			if file_string != "" {
				thumbnail := string([]rune(file.(string))[:8]) + file_string

				flag_bool = false
				return thumbnail
			}
		}
	}
	if flag_bool {

		var digitsRegexp = regexp.MustCompile(`<img.*?(?: |\\t|\\r|\\n)?src=[\'"]?(.+?)[\'"]?(?:(?: |\\t|\\r|\\n)+.*?)?>`)
		img_s := digitsRegexp.FindStringSubmatch(PostContent)
		if len(img_s) == 2 {
			return img_s[1]
		}
	}

	return "/"

}

func (m *PostsInfo) GetCategoryIds(url string) (string, []string) {
	var info Postmeta
	var menus []Menu1
	var CategoryIds []string
	var CategoryName string

	menus = info.GetMenu(url)
	for _, v := range menus {
		if strings.Contains(v.Url, url) {
			CategoryIds = append(CategoryIds, v.Term_id)

		}

		if v.Url == "/"+url {
			CategoryName = v.Name
		}
		for _, k := range v.Sub_menu {
			if strings.Contains(k.Url, url) {
				CategoryIds = append(CategoryIds, k.Term_id)
			}
			if k.Url == "/"+url {
				CategoryName = k.Name
			}
		}
	}
	return CategoryName, CategoryIds
}

type Countmun struct {
	CountID string
}

func (m *PostsInfo) GetCategoryPosts(url string, offset, pagesize int64) (string, int64, []*PostsInfo) {
	CategoryName, CategoryIds := m.GetCategoryIds(url)
	CategoryId := strings.Join(CategoryIds, ",")

	var list []*PostsInfo
	sql := "SELECT p.* FROM so_posts p INNER JOIN so_term_relationships tr ON (p.ID = tr.object_id) WHERE 1=1  AND (tr.term_taxonomy_id IN (" + CategoryId + ")) AND p.post_type = 'post' AND (p.post_status = 'publish' OR p.post_status = 'private') GROUP BY p.ID ORDER BY p.post_date DESC LIMIT ?, ? "
	orm.NewOrm().Raw(sql, offset, pagesize).QueryRows(&list)
	var count []Countmun
	sql1 := "SELECT count(*) CountID FROM so_posts p INNER JOIN so_term_relationships tr ON (p.ID = tr.object_id) WHERE 1=1  AND (tr.term_taxonomy_id IN (" + CategoryId + ")) AND p.post_type = 'post' AND (p.post_status = 'publish' OR p.post_status = 'private') GROUP BY p.ID"
	orm.NewOrm().Raw(sql1).QueryRows(&count)

	return CategoryName, int64(len(count)), list
}

/*
func (m *PostsInfo) GetList(pagesize int64) []*PostsInfo {
	var info PostsInfo
	list := make([]*PostsInfo, 0)
	info.Query().OrderBy("-views").Limit(5, 0).All(&list, "ID", "post_date", "post_title", "post_content")
	return list
}


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
