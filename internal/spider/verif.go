package spider

import (
	"errors"
	// "fmt"

	"github.com/midoks/novelsearch/internal/lazyregexp"
	"github.com/midoks/novelsearch/internal/tools"
)

func VailWdata(d *Wdata) error {

	data, err := tools.GetHttpData(d.Website)

	if err != nil {
		return err
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
		return err
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
		return err
	}

	chapterListUrl := lazyregexp.New(d.Chapter.Rule).FindAllStringSubmatch(chapterData, -1)
	for _, v := range chapterListUrl {
		if !lazyregexp.New(d.Chapter.RootRule).Match(tools.StringToBytes(v[1])) {
			return errors.New("chapter has error match!")
		}
	}

	contentUrl := chapterListUrl[0][1]
	contentData, err := tools.GetHttpData(contentUrl)
	if err != nil {
		return err
	}

	//content
	contentRealData := lazyregexp.New(d.Content.RootRule).FindAllStringSubmatch(contentData, -1)
	if len(contentRealData[0]) < 1 {
		return errors.New("content not match!")
	}

	return nil
}
