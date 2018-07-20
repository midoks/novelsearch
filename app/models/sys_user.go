package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type SysUser struct {
	Id         int
	Username   string
	Nick       string
	Sex        int
	Password   string
	Mail       string
	Tel        string
	Roleid     int
	Status     int
	UpdateTime int64
	CreateTime int64
}

func getTnByUser() string {
	return "sys_user"
}

func (u *SysUser) TableName() string {
	return getTnByUser()
}

func (u *SysUser) Update(fields ...string) error {
	u.UpdateTime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func UserGetList(page, pageSize int, filters ...interface{}) ([]*SysUser, int64) {
	offset := (page - 1) * pageSize

	list := make([]*SysUser, 0)

	query := orm.NewOrm().QueryTable(getTnByUser())

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			print(filters[k].(string), filters[k+1])
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

func UserGetById(id int) (*SysUser, error) {

	u := new(SysUser)
	err := orm.NewOrm().QueryTable(getTnByUser()).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserGetByName(username string) (*SysUser, error) {

	u := new(SysUser)
	err := orm.NewOrm().QueryTable(getTnByUser()).Filter("username", username).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&SysUser{Id: id})
}
