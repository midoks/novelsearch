package backends

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/models"
	"strconv"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	filters := make([]interface{}, 0)

	filters = append(filters, "status", 1)
	itemList, count := models.ItemGetList(1, 10, filters...)

	list := make([]map[string]interface{}, len(itemList))
	for k, v := range itemList {

		row := make(map[string]interface{})
		idStr := strconv.Itoa(v.Id)

		row["Name"] = v.Name
		row["PageIndex"] = v.PageIndex
		row["Count"] = models.NovelCount(idStr)

		list[k] = row
	}

	this.Data["CountList"] = list
	fmt.Println(list)

	this.Data["NovelCount"] = models.NovelCount("")
	this.Data["NovelTodayCount"] = models.NovelTodayCount("")
	this.Data["ItemCount"] = count

	this.display()
}
