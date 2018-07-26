package crontab

import (
	// "errors"
	"fmt"
	"github.com/astaxie/beego/cache"
	// // "github.com/astaxie/beego"
	// // "encoding/json"
	// // "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	// // "github.com/astaxie/beego/orm"
	// "github.com/midoks/novelsearch/app/libs"
	// "github.com/midoks/novelsearch/app/models"
	// // "regexp"
	// "strings"
	// "time"
	// "strconv"
)

func allBaiduTop() {
	// toxp, _ := libs.BaiduTop("http://top.baidu.com/buzz?b=353&c=10")

	bm, _ := cache.NewCache("file", `{"CachePath":"./.cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":0}`)
	// fmt.Println(top)

	// bm.Put("baidu_top", top, 1000*time.Second)

	c := bm.Get("baidu_top")
	fmt.Println(c)
}
