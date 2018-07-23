package models

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type SysOption struct {
	Id         int
	Name       string
	Value      string
	UpdateTime int64
	CreateTime int64
}

func getTnByOption() string {
	return "sys_option"
}

func (u *SysOption) TableName() string {
	return getTnByOption()
}

func (u *SysOption) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func OptionSet(name string, value string) bool {

	data := new(SysOption)
	err := orm.NewOrm().QueryTable(getTnByOption()).Filter("name", name).One(data)
	if err == nil {
		data.UpdateTime = time.Now().Unix()
		data.Value = value
		err = data.Update()
		if err == nil {
			return true
		}
		return true
	}

	data.Name = name
	data.Value = value
	data.UpdateTime = time.Now().Unix()
	data.CreateTime = time.Now().Unix()

	_, err = orm.NewOrm().Insert(data)
	if err == nil {
		return true
	}
	return false
}

func OptionGet(name string, def string) string {

	data := new(SysOption)
	err := orm.NewOrm().QueryTable(getTnByOption()).Filter("name", name).One(data)
	if err == nil {
		return data.Value
	}
	return def
}
