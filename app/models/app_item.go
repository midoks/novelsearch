package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type AppItem struct {
	Id              int
	Name            string
	IsOfficial      string
	PageCharset     string
	PageIndex       string
	PageIndexRule   string
	PathRule        string
	PathTpl         string
	PathPageExp     string
	NameRule        string
	DescRule        string
	AuthorRule      string
	CategoryRule    string
	StatusRule      string
	StatusEndMark   string
	ChapterPathRule string
	ChapterPathExp  string
	ChapterListRule string
	ContentExp      string
	ContentRule     string
	SosoExp         string
	SosoPageArgs    string
	SosoRule        string
	SpiderExp       string
	SpiderRange     string
	SpiderRule      string
	ErrMsg          string
	Status          int
	CronUpTime      int64
	UpdateTime      int64
	CreateTime      int64
}

func getTnByAppItem() string {
	return TableName("item")
}

func (u *AppItem) TableName() string {
	return getTnByAppItem()
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

	query := orm.NewOrm().QueryTable(getTnByAppItem())

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
	// orm.Debug = true
	u := new(AppItem)
	err := orm.NewOrm().QueryTable(getTnByAppItem()).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	// orm.Debug = false
	return u, nil
}

func ItemGetNameById(id int) string {
	// orm.Debug = true
	u := new(AppItem)
	err := orm.NewOrm().QueryTable(getTnByAppItem()).Filter("id", id).One(u)
	if err != nil {
		return "无数据"
	}
	// orm.Debug = false
	return u.Name
}

func ItemGetByName(name string) (*AppItem, error) {

	u := new(AppItem)
	err := orm.NewOrm().QueryTable(getTnByAppItem()).Filter("name", name).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func ItemCount() int64 {
	count, _ := orm.NewOrm().QueryTable(getTnByAppItem()).Filter("status", 1).Count()
	return count
}

func ItemDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&AppItem{Id: id})
}
