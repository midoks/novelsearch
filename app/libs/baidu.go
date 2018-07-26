package libs

import (
	"errors"
	// "fmt"
	"github.com/astaxie/beego/httplib"
	"regexp"
	"time"
)

var (
	BAIDU_KEY = "baidu_top_key"
)

//http://top.baidu.com/buzz?b=353
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
	toplist := make(map[string]string)
	toplist["全部"] = "http://top.baidu.com/buzz?b=7"
	toplist["玄幻奇幻"] = "http://top.baidu.com/buzz?b=353"
	toplist["都市言情"] = "http://top.baidu.com/buzz?b=355"
	toplist["武侠仙侠"] = "http://top.baidu.com/buzz?b=354"
	toplist["青春校园"] = "http://top.baidu.com/buzz?b=1508"
	toplist["穿越架空"] = "http://top.baidu.com/buzz?b=1509"
	toplist["科幻悬疑"] = "http://top.baidu.com/buzz?b=356"
	toplist["历史军事"] = "http://top.baidu.com/buzz?b=459"
	toplist["游戏竞技"] = "http://top.baidu.com/buzz?b=1512"
	toplist["耽美同人"] = "http://top.baidu.com/buzz?b=1510"
	toplist["文学经典"] = "http://top.baidu.com/buzz?b=1513"

	topData := make(map[string]interface{})
	for i, v := range toplist {
		top, err := BaiduTop(v)
		if err == nil {
			topData[i] = top
		}
	}

	SetCache(BAIDU_KEY, topData, 24*60*60*time.Second)
}

func GetAllBaiduTop() map[string]interface{} {
	c := GetCache(BAIDU_KEY)
	return c.(map[string]interface{})
}
