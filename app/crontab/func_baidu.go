package crontab

import (
	"github.com/midoks/novelsearch/app/libs"
)

func BaiduTopAll() error {
	libs.CronSaveAllBaiduTop()
	return nil
}
