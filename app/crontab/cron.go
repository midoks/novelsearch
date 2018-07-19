package crontab

import (
	"fmt"
	// "github.com/astaxie/beego/toolbox"
)

func Init() {
	fmt.Println("crontab init")

	//首页爬取数据
	// PageIndexSpider() //test
	// tk1 := toolbox.NewTask("PageIndexSpider", "0/10 * * * * *", PageIndexSpider)
	// toolbox.AddTask("PageIndexSpider", tk1)

	//对每一个小说进行检查更新
	NovelIndexSpider()
	// tk2 := toolbox.NewTask("NovelIndexSpider", "0/10 * * * * *", NovelIndexSpider)
	// toolbox.AddTask("NovelIndexSpider", tk2)

}
