package models

import "github.com/astaxie/beego/orm"
import "fmt"

type OptionsInfo struct {
	Id          int    `orm:"column(option_id);auto"`
	OptionName  string `orm:"column(option_name);size(191);null"`
	OptionValue string `orm:"column(option_value)"`
	Autoload    string `orm:"column(autoload);size(20)"`
}

func (t *OptionsInfo) TableName() string {
	return "options"
}

func (m *OptionsInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

var OptionMap = make(map[string]string)

func (t *OptionsInfo) AutolaodOption() {
	var info OptionsInfo
	//OptionMap:=make(map[string]string)
	list := make([]*OptionsInfo, 0)
	info.Query().Filter("Autoload", "yes").All(&list)
	for _, v := range list {
		fmt.Println(v.OptionName)
		OptionMap[v.OptionName] = v.OptionValue
	}
	//return OptionMap
}
