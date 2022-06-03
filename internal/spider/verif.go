package spider

import (
	"errors"
	"fmt"
	"strings"

	"github.com/midoks/novelsearch/internal/lazyregexp"
	"github.com/midoks/novelsearch/internal/tools"
)

func VailWdata(d *Wdata) error {

	data, err := tools.GetHttpData(d.Website)

	if err != nil {
		return fmt.Errorf("index page error: [%s]%v", d.Website, err)
	}

	//home
	novelHomeUrl := lazyregexp.New(d.Novel.Rule).FindAllStringSubmatch(data, -1)

	alist := []string{}
	for _, v := range novelHomeUrl {
		if lazyregexp.New(d.Novel.RootRule).Match(tools.StringToBytes(v[1])) {
			alist = append(alist, v[1])
		}
	}

	if len(alist) < 1 {
		return errors.New("alist is empty!")
	}

	//vaild article list page
	oneArtPage := alist[0]

	artData, err := tools.GetHttpData(oneArtPage)
	if err != nil {
		return fmt.Errorf("home page error: %v", err)
	}

	llist := []string{}
	novelListUrl := lazyregexp.New(d.List.Rule).FindAllStringSubmatch(artData, -1)

	for _, v := range novelListUrl {
		if lazyregexp.New(d.List.RootRule).Match(tools.StringToBytes(v[1])) {
			llist = append(llist, v[1])
		}
	}

	if len(llist) < 1 {
		return errors.New("llist is empty!")
	}

	//chapter
	chapterPage := llist[0]
	chapterData, err := tools.GetHttpData(chapterPage)
	if err != nil {
		return fmt.Errorf("chapter page error: %v", err)
	}

	chapterListUrl := lazyregexp.New(d.Chapter.Rule).FindAllStringSubmatch(chapterData, -1)
	for _, v := range chapterListUrl {
		if !lazyregexp.New(d.Chapter.RootRule).MatchString(v[1]) {
			return errors.New("chapter has error match!")
		}
	}

	title := lazyregexp.New(d.Chapter.TitleRule).FindAllStringSubmatch(chapterData, -1)
	if len(title) < 1 {
		return errors.New("chapter title has error match!")
	}

	if d.Chapter.TitleTrimBool {
		titleTrim := strings.Trim(title[0][1], d.Chapter.TitleTrim)
		if len(titleTrim) < 1 {
			return errors.New("chapter title trim has error match!")
		}
	}

	author := lazyregexp.New(d.Chapter.AuthorRule).FindAllStringSubmatch(chapterData, -1)
	if len(author) < 1 {
		return errors.New("chapter author has error match!")
	}

	contentUrl := chapterListUrl[0][1]
	contentData, err := tools.GetHttpData(contentUrl)
	if err != nil {
		return fmt.Errorf("content page error: %v", err)
	}

	//content
	contentRealData := lazyregexp.New(d.Content.RootRule).FindAllStringSubmatch(contentData, -1)
	if len(contentRealData[0]) < 1 {
		return errors.New("content not match!")
	}

	return nil
}
