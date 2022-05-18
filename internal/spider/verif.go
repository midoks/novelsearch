package spider

import (
	"fmt"

	"github.com/midoks/novelsearch/internal/lazyregexp"
	"github.com/midoks/novelsearch/internal/tools"
)

func VailWdata(d *Wdata) error {

	data, err := tools.GetHttpData(d.Website)

	if err != nil {
		return err
	}

	novelHomeUrl := lazyregexp.New(d.Novel.Rule).FindAllStringSubmatch(data, -1)

	for _, v := range novelHomeUrl {

		if lazyregexp.New(d.Novel.RootRule).Match(tools.StringToBytes(v[1])) {
			fmt.Println(v[1])
		}
	}

	// fmt.Println(novelHomeUrl)

	return nil
}
