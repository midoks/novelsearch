package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type AppItem struct {
	Id              int
	Name            string
	PageIndex       string
	PathPageExp     string
	PathRule        string
	NameRule        string
	AuthorRule      string
	CategoryRule    string
	StatusRule      string
	ChapterPathRule string
	ChapterPathExp  string
	ChapterListRule string
	ContentExp      string
	ContentRule     string
	SosoExp         string
	SosoKwCharset   string
	SosoRule        string
	Status          int
	UpdateTime      int64
	CreateTime      int64
}

func (u *AppItem) TableName() string {
	return TableName("item")
}

func (u *AppItem) Update(fields ...string) error {
	u.UpdateTime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func ItemGetList(page, pageSize int, filters ...interface{}) ([]*AppItem, int64) {
	offset := (page - 1) * pageSize

	list := make([]*AppItem, 0)

	query := orm.NewOrm().QueryTable(TableName("item"))

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

func ItemGetById(id int) (*AppItem, error) {

	u := new(AppItem)
	err := orm.NewOrm().QueryTable(TableName("item")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func ItemGetByName(name string) (*AppItem, error) {

	u := new(AppItem)
	err := orm.NewOrm().QueryTable(TableName("item")).Filter("name", name).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func ItemDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&AppItem{Id: id})
}
