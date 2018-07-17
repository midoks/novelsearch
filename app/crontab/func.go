package crontab

import (
	// "errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/midoks/novelsearch/app/libs"
	"log"
	"regexp"
)

func spiderNovelList() error {
	req := httplib.Get("https://www.23us.so/")

	str, err := req.String()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(str)

	// var validID = regexp.MustCompile(`https://www.23us.so/xiaoshuo/3574.(html)`)
	var validID = regexp.MustCompile(`https://www.23us.so/xiaoshuo/(\d*).(html)`)
	var p = libs.RemoveDuplicatesAndEmpty(validID.FindAllString(str, -1))
	fmt.Println(p, len(p))

	return nil
}

func spiderNovelInfo() error {
	req := httplib.Get("https://www.23us.so/xiaoshuo/43.html")

	str, err := req.String()
	if err != nil {
		log.Fatal(err)
	}

	var validID = regexp.MustCompile(`<h1>(.*)全文阅读</h1>`)
	fmt.Println(validID.FindStringSubmatch(str))

	var validID2 = regexp.MustCompile(`<th>小说作者</th>\n<td>&nbsp;(.*)</td>`)
	fmt.Println(validID2.FindStringSubmatch(str))

	var validID3 = regexp.MustCompile(`<th>小说状态</th>\n<td>&nbsp;(.*)</td>`)
	fmt.Println(validID3.FindStringSubmatch(str))

	var validID4 = regexp.MustCompile(`<th>小说类别</th>\n<td>&nbsp;<a href=".*">(.*)</a></td>`)
	fmt.Println(validID4.FindStringSubmatch(str))

	var validID5 = regexp.MustCompile(`<a class="read" href="(.*)" title=".*">最新章节</a>`)
	fmt.Println(validID5.FindStringSubmatch(str))

	return nil
}

func spiderNovelChapter() error {
	req := httplib.Get("https://www.23us.so/files/article/html/0/43/index.html")

	str, err := req.String()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(str)

	var validID5 = regexp.MustCompile(`(?U)<td class="L"><a href="(.*)">(.*)</a></td>`)
	list := validID5.FindAllStringSubmatch(str, -1)
	fmt.Println(len(list))
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
		//fmt.Println(validID5.FindStringSubmatch(list[i]))
	}
	return nil
}

func spiderNovelContent() error {
	req := httplib.Get("https://www.23us.so/files/article/html/0/43/11417212.html")

	str, err := req.String()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(str)
	var validID = regexp.MustCompile(`(?iUs)<dd id="contents">(.*)</dd>`)
	fmt.Println(validID.FindStringSubmatch(str)[1])
	return nil
}
