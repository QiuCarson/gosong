package models

import (
	"time"

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

func (m *CommentsInfo) PostIdByComment() {

}
