package spider

import (
	"fmt"

	"github.com/midoks/novelsearch/internal/lazyregexp"
	"github.com/midoks/novelsearch/internal/tools"
)

func VailWdata(d *Wdata) error {

	data, err := tools.GetHttpData(d.Website)

	novelHomeUrl := lazyregexp.New(d.Novel.Rule).FindAllStringSubmatch(data, -1)

	fmt.Println(data, err)
	fmt.Println(novelHomeUrl)

	return nil
}
