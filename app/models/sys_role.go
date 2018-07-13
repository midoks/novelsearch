package models

import (
	"github.com/astaxie/beego/orm"
)

type SysRole struct {
	Id         int
	Name       string
	Desc       string
	List       string
	Status     int
	UpdateTime int64
	CreateTime int64
}

func (u *SysRole) TableName() string {
	return "sys_role"
}

func (u *SysRole) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func RoleGetAll() ([]*SysRole, int64) {

	list := make([]*SysRole, 0)

	query := orm.NewOrm().QueryTable("sys_role")
	total, _ := query.Count()
	query.OrderBy("-id").All(&list)

	return list, total
}

func RoleGetById(id int) (*SysRole, error) {

	u := new(SysRole)
	err := orm.NewOrm().QueryTable("sys_role").Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func RoleGetList(page, pageSize int, filters ...interface{}) ([]*SysRole, int64) {
	offset := (page - 1) * pageSize

	list := make([]*SysRole, 0)

	query := orm.NewOrm().QueryTable("sys_role")

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

func RoleDelById(id int) (int64,error) {
	return orm.NewOrm().Delete(&SysRole{Id: id})
}
