package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SoLinks struct {
	Id              int       `orm:"column(link_id);auto"`
	LinkUrl         string    `orm:"column(link_url);size(255)"`
	LinkName        string    `orm:"column(link_name);size(255)"`
	LinkImage       string    `orm:"column(link_image);size(255)"`
	LinkTarget      string    `orm:"column(link_target);size(25)"`
	LinkDescription string    `orm:"column(link_description);size(255)"`
	LinkVisible     string    `orm:"column(link_visible);size(20)"`
	LinkOwner       uint64    `orm:"column(link_owner)"`
	LinkRating      int       `orm:"column(link_rating)"`
	LinkUpdated     time.Time `orm:"column(link_updated);type(datetime)"`
	LinkRel         string    `orm:"column(link_rel);size(255)"`
	LinkNotes       string    `orm:"column(link_notes)"`
	LinkRss         string    `orm:"column(link_rss);size(255)"`
}

func (t *SoLinks) TableName() string {
	return "so_links"
}

func init() {
	orm.RegisterModel(new(SoLinks))
}

// AddSoLinks insert a new SoLinks into database and returns
// last inserted Id on success.
func AddSoLinks(m *SoLinks) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSoLinksById retrieves SoLinks by Id. Returns error if
// Id doesn't exist
func GetSoLinksById(id int) (v *SoLinks, err error) {
	o := orm.NewOrm()
	v = &SoLinks{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSoLinks retrieves all SoLinks matches certain condition. Returns empty list if
// no records exist
func GetAllSoLinks(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SoLinks))
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

	var l []SoLinks
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

// UpdateSoLinks updates SoLinks by Id and returns error if
// the record to be updated doesn't exist
func UpdateSoLinksById(m *SoLinks) (err error) {
	o := orm.NewOrm()
	v := SoLinks{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSoLinks deletes SoLinks by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSoLinks(id int) (err error) {
	o := orm.NewOrm()
	v := SoLinks{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SoLinks{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
