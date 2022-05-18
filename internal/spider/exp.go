package spider

import (
	"fmt"

	"encoding/json"
)

type WdataNovel struct {
	Rule     string `json:"rule"`
	RootRule string `json:"root_rule"`
}

type WdataList struct {
	Rule     string `json:"rule"`
	RootRule string `json:"root_rule"`
}

type WdataChapter struct {
	Rule     string `json:"rule"`
	RootRule string `json:"root_rule"`
}

type Wdata struct {
	Tag     string       `json:"tag"`
	Website string       `json:"website"`
	Novel   WdataNovel   `json:"novel"`
	List    WdataList    `json:"list"`
	Chapter WdataChapter `json:"chapter"`
}

func expInit() {

	t := &Wdata{}

	t.Tag = "ddxsku"
	t.Website = "http://www.ddxsku.com"
	t.Novel = WdataNovel{}

	t.Novel.Rule = "href=\"(.*?)\""
	t.Novel.RootRule = "http://www.ddxsku.com/xiaoshuo/(.*).html"

	t.List = WdataList{}
	t.List.Rule = "href=\"(.*?)\""
	t.List.RootRule = "http://www.ddxsku.com/files/article/html/(.*?)/(.*?)/index.html"

	t.Chapter = WdataChapter{}
	t.Chapter.Rule = "<td class=\"L\"><a href=\"(.*?)\">(.*?)</a></td>"
	t.Chapter.RootRule = "http://www.ddxsku.com/files/article/html/(.*?)/(.*?)/index.html"

	if err := VailWdata(t); err != nil {
		return
	}

	rawBytes, err := json.MarshalIndent(t, "", "")
	fmt.Println(string(rawBytes), err)

}
