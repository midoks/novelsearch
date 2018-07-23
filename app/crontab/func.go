package crontab

import (
	"errors"
	// "github.com/astaxie/beego"
	// "encoding/json"
	"github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/orm"
	// "github.com/midoks/novelsearch/app/libs"
	// "github.com/midoks/novelsearch/app/models"
	"regexp"
	// "strconv"
	"strings"
	// "time"
)

func getHttpData(url string) (string, error) {
	req := httplib.Get(url)

	str, err := req.String()
	if err != nil {
		return "", errors.New("资源获取错误!")
	}

	return str, nil
}

//匹配路径
func RegPathInfo(content, reg string) ([][]string, error) {
	match_exp := regexp.MustCompile(reg)
	list := match_exp.FindAllStringSubmatch(content, -1)

	if len(list) == 0 {
		return nil, errors.New("没有匹配到!")
	}
	return list, nil
}

//匹配当个信息
func RegNovelSigleInfo(content, reg string) (string, error) {

	match_exp := regexp.MustCompile(reg)
	name := match_exp.FindAllStringSubmatch(content, -1)
	// logs.Warn(name)
	if len(name) == 0 {
		return "", errors.New("没有匹配到!")
	}
	return strings.TrimSpace(name[0][1]), nil
}

//匹配当个信息
func RegNovelList(content, reg string) ([]map[string]interface{}, error) {

	match_exp := regexp.MustCompile(reg)
	name := match_exp.FindAllStringSubmatch(content, -1)

	list := make([]map[string]interface{}, len(name))

	if len(name) == 0 {
		return list, errors.New("没有匹配到!")
	}

	for k, v := range name {
		tmp := make(map[string]interface{})
		tmp["name"] = v[2]
		tmp["url"] = v[1]
		list[k] = tmp
		// fmt.Println(list[k], v[1], v[2])
	}
	return list, nil
}
