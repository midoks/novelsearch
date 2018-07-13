package libs

import (
	"fmt"
	"github.com/astaxie/beego"
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
