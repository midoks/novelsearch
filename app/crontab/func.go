package crontab

import (
	"errors"
	// "fmt"
	"github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	"github.com/midoks/novelsearch/app/libs"
	"path/filepath"
	"regexp"
	"strings"
)

func getHttpData(url string) (string, error) {
	req := httplib.Get(url)
	req.Header("Accept-Encoding", "gzip,deflate,sdch")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")

	str, err := req.String()
	if err != nil {
		return "", err
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
	// reg = fmt.Sprintf("`%s`", reg)
	// fmt.Println(reg)
	match_exp := regexp.MustCompile(reg)
	name := match_exp.FindAllStringSubmatch(content, -1)

	if len(name) == 0 {
		return "", errors.New("没有匹配到!")
	}
	// logs.Warn(reg, name[0])
	return strings.TrimSpace(name[0][1]), nil
}

func GetAbsoluteAddr(cur, result string) string {
	if libs.IsUrlRe(result) {
		return result
	}
	main := strings.Replace(filepath.Dir(cur), "http:/", "http://", -1)
	main = strings.Replace(main, "https:/", "https://", -1)
	return main + "/" + result
}

//匹配当个信息
func RegNovelListAutoPath(content, reg string, curPage string) ([]map[string]interface{}, error) {

	match_exp := regexp.MustCompile(reg)
	name := match_exp.FindAllStringSubmatch(content, -1)

	list := make([]map[string]interface{}, len(name))

	if len(name) == 0 {
		return list, errors.New("没有匹配到!")
	}

	for k, v := range name {
		tmp := make(map[string]interface{})

		tmp["name"] = v[2]
		ab_path := GetAbsoluteAddr(curPage, v[1])
		tmp["url"] = ab_path
		list[k] = tmp
		// fmt.Println(list[k], v[1], v[2])
	}
	return list, nil
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
