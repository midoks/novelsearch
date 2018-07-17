package crontab

import (
	"fmt"
	// "github.com/astaxie/beego/toolbox"
)

func Init() {
	fmt.Println("crontab init")

	PageIndexSpider()

	//首页爬取数据
	//tk1 := toolbox.NewTask("PageIndexSpider", "0/10 * * * * *", PageIndexSpider)
	//toolbox.AddTask("PageIndexSpider", tk1)

	// tkList := toolbox.NewTask("getNovel", "0/10 * * * * *", spiderNovelList)
	// toolbox.AddTask("获取小说列表", tkList)

	//手册

}
