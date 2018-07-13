package crontab

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"github.com/midoks/novelsearch/app/models"
)

type DatabaseCheck struct {
}

func (dc *DatabaseCheck) Check() error {
	b := models.MysqlPing()
	if b {
		return nil
	}
	return errors.New("can't connect database")
}

func Init() {
	fmt.Println("crontab init")

	toolbox.AddHealthCheck("database", &DatabaseCheck{})

	tk1 := toolbox.NewTask("getNovel", "0/30 * * * * *", func() error { fmt.Println("tk1"); return nil })
	toolbox.AddTask("获取小说", tk1)
}
