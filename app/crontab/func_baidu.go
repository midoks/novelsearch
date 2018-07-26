package crontab

import (
	"github.com/midoks/novelsearch/app/libs"
)

func BaiduTopAll() error {
	_, err := libs.GetAllBaiduTop()
	if err != nil {
		libs.CronSaveAllBaiduTop()
	}
	return nil
}
