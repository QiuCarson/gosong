package models

import (
	"strconv"
	"strings"
	"time"

	"fmt"

	"github.com/astaxie/beego/orm"
)

type CommentsInfo struct {
	Id                 int       `orm:"column(comment_ID);auto"`
	CommentPostID      uint64    `orm:"column(comment_post_ID)"`
	CommentAuthor      string    `orm:"column(comment_author)"`
	CommentAuthorEmail string    `orm:"column(comment_author_email);size(100)"`
	CommentAuthorUrl   string    `orm:"column(comment_author_url);size(200)"`
	CommentAuthorIP    string    `orm:"column(comment_author_IP);size(100)"`
	CommentDate        time.Time `orm:"column(comment_date);type(datetime)"`
	CommentDateGmt     time.Time `orm:"column(comment_date_gmt);type(datetime)"`
	CommentContent     string    `orm:"column(comment_content)"`
	CommentKarma       int       `orm:"column(comment_karma)"`
	CommentApproved    string    `orm:"column(comment_approved);size(20)"`
	CommentAgent       string    `orm:"column(comment_agent);size(255)"`
	CommentType        string    `orm:"column(comment_type);size(20)"`
	CommentParent      uint64    `orm:"column(comment_parent)"`
	UserId             uint64    `orm:"column(user_id)"`
	CommentMailNotify  int8      `orm:"column(comment_mail_notify)"`
}

func (t *CommentsInfo) TableName() string {
	return "comments"
}

func (m *CommentsInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

type CommentCount struct {
	Countmun int64
}
type CommentId struct {
	CommentID int
}

func PostIdByComment() {
	var (
		post_id      = 2775
		commentCount CommentCount
		offset       int64
		pagesize     int64
		comment      []*CommentsInfo
		comment_ids  []string
		//comment_ids1   []string
		//comment_parent []*CommentsInfo
		//comment_parent_id_str string
	)
	pagesize = 10

	sqlCount := "SELECT  COUNT(*) countmun FROM so_comments  WHERE ( ( comment_approved = '1' ) OR ( user_id = 1 AND comment_approved = '0' ) ) AND comment_post_ID = ? AND comment_parent = 0 "
	orm.NewOrm().Raw(sqlCount, post_id).QueryRow(&commentCount)

	offset = commentCount.Countmun / pagesize * pagesize
	sqlComment := "SELECT * FROM so_comments  WHERE ( ( comment_approved = '1' ) OR ( user_id = 1 AND comment_approved = '0' ) ) AND comment_post_ID = ? AND comment_parent = 0  ORDER BY comment_date_gmt ASC, comment_ID ASC LIMIT ?,? "
	orm.NewOrm().Raw(sqlComment, post_id, offset, pagesize).QueryRows(&comment)
	for _, v := range comment {
		comment_ids = append(comment_ids, strconv.Itoa(v.Id))
	}
	comment_ids_str := strings.Join(comment_ids, ",")
	comment_parent_id_str := commentfor(comment_ids_str, post_id)

	sqlParentComment := "SELECT * FROM so_comments  WHERE comment_id in(" + comment_parent_id_str + ")"
	orm.NewOrm().Raw(sqlParentComment).QueryRows(&comment)
	fmt.Println(comment)

}

func commentfor(comment_ids_str string, post_id int) string {
	var (
		comment_parent   []*CommentsInfo
		comment_ids      []string
		comment_ids_strs string
	)
	fmt.Println(comment_ids_str)
	sqlCommentId := "SELECT comment_ID FROM so_comments  WHERE ( ( comment_approved = '1' ) OR ( user_id = 1 AND comment_approved = '0' ) ) AND comment_post_ID = ? AND comment_parent IN (" + comment_ids_str + ")  ORDER BY comment_date_gmt ASC, comment_ID ASC"
	num, _ := orm.NewOrm().Raw(sqlCommentId, post_id).QueryRows(&comment_parent)
	if num < 1 {
		return ""
	}
	for _, v := range comment_parent {
		//fmt.Println(v.Id)
		comment_ids = append(comment_ids, strconv.Itoa(v.Id))

	}
	//fmt.Println(comment_ids)
	comment_ids_str = strings.Join(comment_ids, ",")
	//fmt.Println("comment_ids_str", comment_ids_str)
	comment_ids_str_temp := commentfor(comment_ids_str, post_id)
	//fmt.Println("comment_ids_str_temp", comment_ids_str_temp)
	if len(comment_ids) > 1 {
		comment_ids_strs = comment_ids_str + "," + comment_ids_str_temp
	} else {
		comment_ids_strs = comment_ids_str
	}

	//fmt.Println("comment_ids_strs", comment_ids_strs)
	return comment_ids_strs

}
