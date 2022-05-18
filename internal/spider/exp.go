package spider

import (
	"fmt"

	"encoding/json"
)

type WdataNovel struct {
	Rule     string `json:"rule"`
	RootRule string `json:"root_rule"`
}

type Wdata struct {
	Tag     string     `json:"tag"`
	Website string     `json:"website"`
	Novel   WdataNovel `json:"novel"`
}

func expInit() {

	t := &Wdata{}

	t.Tag = "ddxsku"
	t.Website = "http://www.ddxsku.com"
	t.Novel = WdataNovel{}

	t.Novel.Rule = "href=\"(.*?)\""
	t.Novel.RootRule = "http://www.ddxsku.com/xiaoshuo/(.*).html"

	if err := VailWdata(t); err != nil {
		return
	}

	rawBytes, err := json.MarshalIndent(t, "", "")
	fmt.Println(string(rawBytes), err)

}
