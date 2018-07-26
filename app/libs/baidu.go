package libs

import (
	"errors"
	// "fmt"
	"github.com/astaxie/beego/httplib"
	"regexp"
)

//http://top.baidu.com/buzz?b=353&c=10
//单个榜单的数据
func BaiduTop(url string) ([]string, error) {
	req := httplib.Get(url)
	content, err := req.String()

	if err != nil {
		return nil, errors.New("请求url失败!")
	}

	content = ConvertToString(content, "gbk", "utf8")
	valid := regexp.MustCompile(`<a class="list-title" target="_blank" href=".*">(.*)</a>`)
	match := valid.FindAllStringSubmatch(content, -1)

	count := len(match)
	var list = []string{}
	if count > 0 {
		for i := 0; i < count; i++ {
			list = append(list, match[i][1])
		}
		return list, nil
	}

	return nil, errors.New("没有匹配到值!")
}

func CronSaveAllBaiduTop() {

}
