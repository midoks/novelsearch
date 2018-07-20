package crontab

import (
	"fmt"
	"github.com/astaxie/beego/toolbox"
)

//	 minute hour day month week   command
//顺序：分    时   日   月    周
func Init() {
	fmt.Println("crontab init")

	tk0 := toolbox.NewTask("test", "0 * * * * *", func() error { fmt.Println("begin--cron"); return nil })
	toolbox.AddTask("test", tk0)
	//首页爬取数据
	// PageIndexSpider() //test
	// tk1 := toolbox.NewTask("PageIndexSpider", "0/10 * * * * *", PageIndexSpider)
	// toolbox.AddTask("PageIndexSpider", tk1)

	//对每一个小说进行检查更新
	NovelIndexSpider()
	tk2 := toolbox.NewTask("NovelIndexSpider", "0/30 * * * * *", NovelIndexSpider)
	toolbox.AddTask("NovelIndexSpider", tk2)

}
