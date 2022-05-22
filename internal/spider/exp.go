package spider

import (
	"encoding/json"
	"fmt"

	"github.com/midoks/novelsearch/internal/mgdb"
	"github.com/midoks/novelsearch/internal/tools"
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
	Rule          string `json:"rule"`
	RootRule      string `json:"root_rule"`
	TitleRule     string `json:"title_rule"`
	TitleTrimBool bool   `json:"title_trim_rule"`
	TitleTrim     string `json:"title_space_rule"`
	AuthorRule    string `json:"author_rule"`
}

type WdataContent struct {
	RootRule string `json:"root_rule"`
}

type Wdata struct {
	Tag     string       `json:"tag"`
	Website string       `json:"website"`
	Novel   WdataNovel   `json:"novel"`
	List    WdataList    `json:"list"`
	Chapter WdataChapter `json:"chapter"`
	Content WdataContent `json:"content"`
}

func expInit(ruleDir string) {

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
	t.Chapter.RootRule = "http://www.ddxsku.com/files/article/html/(.*?)/(.*?)/(.*?).html"
	t.Chapter.TitleRule = "<dd><h1>(.*?)</h1></dd>"
	t.Chapter.TitleTrimBool = true
	t.Chapter.TitleTrim = "最新章节"
	t.Chapter.AuthorRule = "<dd><h3>作者：(.*?)</h3><br>"

	t.Content = WdataContent{}
	t.Content.RootRule = `(?ims)<dd\s*id=\"contents\">(.*?)<\/dd>`

	//验证功能测试
	if err := VailWdata(t); err != nil {
		fmt.Println(err)
		return
	}

	rawBytes, err := json.MarshalIndent(t, "", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	f := ruleDir + "/" + t.Tag + ".json"
	tools.WriteFile(f, tools.BytesToString(rawBytes))

	mgdb.NovelSourceAdd(mgdb.NovelSource{Name: t.Tag, RuleJson: tools.BytesToString(rawBytes)})
	// fmt.Println(string(rawBytes))
}
