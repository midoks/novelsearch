package spider

import (
	"fmt"

	"encoding/json"
)

type WdataNovel struct {
	Rule     string `json:"rule"`
	RootRule string `json:"rule"`
}

type Wdata struct {
	Website string     `json:"website"`
	Novel   WdataNovel `json:"novel"`
}

func expInit() {

	t := &Wdata{}
	t.Website = "http://www.ddxsku.com"
	t.Novel.Rule = "href=\"(.*)\""
	t.Novel.RootRule = "http://www.ddxsku.com/xiaoshuo/(.*).html"

	if err := VailWdata(t); err != nil {
		return
	}

	rawBytes, err := json.MarshalIndent(t, "", "")
	fmt.Println(string(rawBytes), err)

}
