package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
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
	DClick         int
	MClick         int
	YClick         int
	CClick         int
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

func SosoNovelByKw(kw string) []AppNovel {
	var list []AppNovel
	orm.NewOrm().QueryTable(getTnByAppNovel()).
		Filter("name__icontains", kw).
		OrderBy("-id").
		Limit(10).
		All(&list, "id", "name", "from_id", "unique_id")
	fmt.Println(list)
	return list
}

func NovelDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&AppItem{Id: id})
}
