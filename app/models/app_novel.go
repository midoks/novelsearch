package models

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type AppNovel struct {
	Id             int
	FromId         int
	Url            string
	UniqueId       string
	Name           string
	Category       string
	Desc           string
	Author         string
	List           string
	ChapterNum     int
	LastChapter    string
	LastChapterUrl string
	BookStatus     int
	Status         int
	UpdateTime     int64
	CreateTime     int64
}

func getTnByAppNovel() string {
	return TableName("novel")
}

func (u *AppNovel) TableName() string {
	return getTnByAppNovel()
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

	query := orm.NewOrm().QueryTable(getTnByAppNovel())

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
	err := orm.NewOrm().QueryTable(getTnByAppNovel()).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func NovelGetByIdStr(id string) (*AppNovel, error) {
	u := new(AppNovel)
	err := orm.NewOrm().QueryTable(getTnByAppNovel()).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func NovelGetByName(name string) (*AppNovel, error) {

	u := new(AppNovel)
	err := orm.NewOrm().QueryTable(getTnByAppNovel()).Filter("name", name).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CronNovelGetByNameAndFromId(name string, fromid string) (*AppNovel, error) {
	u := new(AppNovel)
	err := orm.NewOrm().QueryTable(getTnByAppNovel()).Filter("name", name).Filter("fromId", fromid).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func NovelGetByUniqueId(unique_id string) (*AppNovel, error) {

	u := new(AppNovel)
	err := orm.NewOrm().QueryTable(getTnByAppNovel()).Filter("unique_id", unique_id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CronNovelGetByStatus(status string, interval int64, num int) []*AppNovel {
	list := make([]*AppNovel, 0)

	now := time.Now().Unix()
	intervalBefore := now - interval
	orm.NewOrm().QueryTable(getTnByAppNovel()).
		Filter("status", status).
		Filter("update_time__lt", intervalBefore).
		OrderBy("-id").
		Limit(num).
		All(&list)
	return list
}

func SosoNovelByKw(kw string, page, pageSize int) ([]AppNovel, int64) {
	offset := (page - 1) * pageSize

	var list []AppNovel
	// orm.Debug = true

	cond := orm.NewCondition()
	cond1 := cond.And("name__istartswith", kw).
		Or("author__istartswith", kw)

	query := orm.NewOrm().QueryTable(getTnByAppNovel()).
		SetCond(cond1).
		OrderBy("-id")

	total, _ := query.Count()

	query.Limit(pageSize, offset).All(&list, "id",
		"name",
		"from_id",
		"unique_id",
		"author",
		"last_chapter",
		"last_chapter_url",
		"update_time")
	// orm.Debug = false
	return list, total
}

func NovelCount(from_id string) int64 {
	// orm.Debug = true
	var count int64
	query := orm.NewOrm().QueryTable(getTnByAppNovel())
	if strings.EqualFold(from_id, "") {
		count, _ = query.Count()
	} else {
		count, _ = query.Filter("from_id", from_id).Count()
	}
	// orm.Debug = false
	return count
}

func NovelTodayCount(from_id string) int64 {
	// orm.Debug = true

	t := time.Now()
	the_time := time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, time.Local)
	unix_time := the_time.Unix()
	// fmt.Println(t.Year(), t.Month(), t.Day(), unix_time)

	query := orm.NewOrm().QueryTable(getTnByAppNovel()).Filter("create_time__gt", unix_time)
	var count int64
	if strings.EqualFold(from_id, "") {
		count, _ = query.Count()
	} else {
		count, _ = query.Filter("from_id", from_id).Count()
	}
	// orm.Debug = false
	return count
}

func NovelDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&AppItem{Id: id})
}
