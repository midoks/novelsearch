package cmd

import (
	"encoding/json"
	// "errors"
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli"

	"github.com/midoks/novelsearch/internal/lazyregexp"
	"github.com/midoks/novelsearch/internal/mgdb"
	"github.com/midoks/novelsearch/internal/spider"
	"github.com/midoks/novelsearch/internal/tools"
)

var Robot = cli.Command{
	Name:        "robot",
	Usage:       "This Robot services",
	Description: `Start robot services`,
	Action:      runRobotService,
	Flags: []cli.Flag{
		stringFlag("mode,m", "index", "choose mode to run (index,soso,chapter)"),
	},
}

func runRobotService(c *cli.Context) error {

	mode := c.String("mode")

	// 首页抓取
	if mode == "index" {
		nlist, err := mgdb.NovelSourceOriginFind(1)
		if err != nil {
			return fmt.Errorf("find error: %v", err)
		}

		var w spider.Wdata
		for _, novel := range nlist {

			err := json.Unmarshal([]byte(novel.RuleJson), &w)
			if err != nil {
				return fmt.Errorf("json.Unmarshal error: %v", err)
			}

			homeData, err := tools.GetHttpData(w.Website)
			if err != nil {
				return fmt.Errorf("home page error: %v", err)
			}

			llist := []string{}
			novelListUrl := lazyregexp.New(w.List.Rule).FindAllStringSubmatch(homeData, -1)

			for _, v := range novelListUrl {
				if lazyregexp.New(w.List.RootRule).Match(tools.StringToBytes(v[1])) {
					llist = append(llist, v[1])
				}
			}

			for _, vlist := range llist {

				chapterPageData, err := tools.GetHttpData(vlist)
				if err != nil {
					return fmt.Errorf("chapter page error: %v", err)
				}

				realTitle := ""
				title := lazyregexp.New(w.Chapter.TitleRule).FindAllStringSubmatch(chapterPageData, -1)
				if w.Chapter.TitleTrimBool {
					titleTrim := strings.Trim(title[0][1], w.Chapter.TitleTrim)
					realTitle = titleTrim
				}

				realAuthor := ""
				authorRule := lazyregexp.New(w.Chapter.AuthorRule).FindAllStringSubmatch(chapterPageData, -1)
				realAuthor = authorRule[0][1]

				chapterList := lazyregexp.New(w.Chapter.Rule).FindAllStringSubmatch(chapterPageData, -1)

				realChapter := [][]string{}
				for _, cl := range chapterList {
					tmp := []string{}

					tmp = append(tmp, cl[1])
					tmp = append(tmp, cl[2])

					realChapter = append(realChapter, tmp)
				}

				dataLen := len(realChapter)

				newsest := realChapter[dataLen-1]
				rDataChapter, _ := json.Marshal(realChapter)

				n := mgdb.Novel{
					Name:          realTitle,
					Author:        realAuthor,
					NewChapter:    newsest[0],
					NewChapterUrl: newsest[1],
					Chapter:       string(rDataChapter),
					Source:        w.Tag,
					Url:           vlist,
				}
				mgdb.NovelAdd(n)

				fmt.Println(realTitle)

				time.Sleep(time.Second * 1)
			}

		}

	}

	// 搜索抓取

	// 章节更新

	// 未完结更新
	return nil
}
