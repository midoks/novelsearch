package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type AppNovel struct {
	Id         int
	Name       string
	Desc       string
	Status     int
	UpdateTime int64
	CreateTime int64
}

func (u *AppNovel) TableName() string {
	return TableName("novel")
}

func (u *AppNovel) Update(fields ...string) error {
	u.UpdateTime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func NovelGetList(page, pageSize int, filters ...interface{}) ([]*AppNovel, int64) {
	offset := (page - 1) * pageSize

	list := make([]*AppNovel, 0)

	query := orm.NewOrm().QueryTable(TableName("novel"))

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			// print(filters[k].(string), filters[k+1])
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

func NovelGetById(id int) (*AppNovel, error) {

	u := new(AppNovel)
	err := orm.NewOrm().QueryTable(TableName("novel")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func NovelGetByName(name string) (*AppNovel, error) {

	u := new(AppNovel)
	err := orm.NewOrm().QueryTable(TableName("novel")).Filter("name", name).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func NovelDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&AppItem{Id: id})
}
