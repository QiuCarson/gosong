package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SoComments struct {
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

func (t *SoComments) TableName() string {
	return "so_comments"
}

func init() {
	orm.RegisterModel(new(SoComments))
}

// AddSoComments insert a new SoComments into database and returns
// last inserted Id on success.
func AddSoComments(m *SoComments) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSoCommentsById retrieves SoComments by Id. Returns error if
// Id doesn't exist
func GetSoCommentsById(id int) (v *SoComments, err error) {
	o := orm.NewOrm()
	v = &SoComments{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSoComments retrieves all SoComments matches certain condition. Returns empty list if
// no records exist
func GetAllSoComments(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SoComments))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SoComments
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSoComments updates SoComments by Id and returns error if
// the record to be updated doesn't exist
func UpdateSoCommentsById(m *SoComments) (err error) {
	o := orm.NewOrm()
	v := SoComments{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSoComments deletes SoComments by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSoComments(id int) (err error) {
	o := orm.NewOrm()
	v := SoComments{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SoComments{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
