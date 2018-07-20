package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type AppDebug struct {
	Id      int
	Type    int
	Msg     string
	AddTime int64
}

func getTnByAppDebug() string {
	return TableName("debug")
}

func (u *AppDebug) TableName() string {
	return getTnByAppDebug()
}

func (u *AppDebug) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func DebugGetList(page, pageSize int, filters ...interface{}) ([]*AppDebug, int64) {
	offset := (page - 1) * pageSize

	list := make([]*AppDebug, 0)

	query := orm.NewOrm().QueryTable(getTnByAppDebug())

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

func DebugAdd(utype int, msg string) {
	var log AppDebug

	log.Type = utype
	log.Msg = msg
	log.AddTime = time.Now().Unix()

	id, err := orm.NewOrm().Insert(&log)
	if err == nil {
		fmt.Println("AppDebug.DebugAdd:", id)
	}
}
