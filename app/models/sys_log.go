package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	SYS  = 1
	MAIL = 2
	ERR  = 3
)

type SysLog struct {
	Id      int
	Uid     int
	Type    int
	Msg     string
	AddTime int64
}

func (u *SysLog) TableName() string {
	return "sys_logs"
}

func (u *SysLog) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func LogGetList(page, pageSize int, filters ...interface{}) ([]*SysLog, int64) {
	offset := (page - 1) * pageSize

	list := make([]*SysLog, 0)

	query := orm.NewOrm().QueryTable("sys_logs")

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

func LogAdd(uid int, utype int, msg string) {
	o := orm.NewOrm()
	var log SysLog

	log.Uid = uid
	log.Type = utype
	log.Msg = msg
	log.AddTime = time.Now().Unix()

	id, err := o.Insert(&log)
	if err == nil {
		fmt.Println("SysLog.LogAdd:", id)
	}
}
