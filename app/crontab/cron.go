package crontab

import (
	"fmt"
	"github.com/midoks/novelsearch/app/libs"
	// "github.com/midoks/novelsearch/app/models"
	// "log"
	// "regexp"
	// "time"
)

func Init() {
	fmt.Println("crontab init")

	// spiderNovelList()
	// spiderNovelInfo()
	// spiderNovelChapter()
	// spiderNovelContent()

	libs.BaiduTop("http://top.baidu.com/buzz?b=353&c=10")

	// tk1 := toolbox.NewTask("getNovel", "0/30 * * * * *", func() error { fmt.Println("tk1"); return nil })
	// toolbox.AddTask("获取小说", tk1)

	//tkList := toolbox.NewTask("getNovel", "0/10 * * * * *", spiderNovelList)
	//toolbox.AddTask("获取小说列表", tkList)

	//手册

}
