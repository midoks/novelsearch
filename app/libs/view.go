package libs

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/models"
	"strconv"
	"strings"
	"time"
)

func Init() {
	fmt.Println("libs init")

	beego.AddFuncMap("hi", tplHello)
	beego.AddFuncMap("isStrInList", tplIsStrInList)
	beego.AddFuncMap("isIntInList", tplIsIntInList)
	beego.AddFuncMap("loadtimes", loadtimes)
	beego.AddFuncMap("adminPath", getAdminPath)
	beego.AddFuncMap("base64encode", tplBase64encode)

	beego.AddFuncMap("webName", func() string {
		return models.OptionGet(models.WEB_NAME, "小说搜索")
	})

	beego.AddFuncMap("webKeyWord", func() string {
		return models.OptionGet(models.WEB_KEYWORD, "关键字")
	})

	beego.AddFuncMap("webDesc", func() string {
		return models.OptionGet(models.WEB_DESC, "默认描述")
	})

	beego.AddFuncMap("webStat", func() string {
		return models.OptionGet(models.WEB_STAT, "")
	})

	beego.AddFuncMap("webNotice", func() string {
		return models.OptionGet(models.WEB_NOTICE, "")
	})
}

func tplBase64encode(in int) string {
	d := strconv.Itoa(in)
	return Base64encode(d)
}

func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func getAdminPath() string {
	admin_name := beego.AppConfig.String("admin_path")
	return admin_name
}

func tplHello(in string) (out string) {
	out = in + "world"
	return
}

func tplIsIntInList(check int, list string) (out bool) {
	out = false
	numList := strings.Split(list, ",")
	for i := 0; i < len(numList); i++ {
		if numList[i] == strconv.Itoa(check) {
			out = true
		}
	}
	return
}

func tplIsStrInList(check string, list string) (out bool) {
	out = false
	numList := strings.Split(list, ",")
	for i := 0; i < len(numList); i++ {
		if numList[i] == check {
			out = true
		}
	}
	return
}
