package fontends

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"strings"
	"time"
)

type IndexController struct {
	CommonController
}

func (this *IndexController) Index() {

	filters := make([]interface{}, 0)
	result, _ := models.NovelGetList(1, 8, filters...)

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		fName := models.ItemGetNameById(v.FromId)

		row["Id"] = v.Id
		row["Name"] = v.Name
		row["Desc"] = v.Desc
		row["Author"] = v.Author
		row["FromId"] = v.FromId
		row["FromName"] = fName
		row["UniqueId"] = v.UniqueId
		row["Status"] = v.Status
		row["UpdateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["CreateTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}
	this.Data["list"] = list

	this.display()
}

func (this *IndexController) Top() {
	this.display()
}

func (this *IndexController) Baidutop() {
	list, err := libs.GetAllBaiduTop()
	if err == nil {
		this.Data["list"] = list
	}
	this.display()
}

func (this *IndexController) Soso() {
	kw := this.GetString("wd")
	if strings.EqualFold(kw, "") {
		this.redirect("/")
	}
	fmt.Println(kw)
	list := models.SosoNovelByKw(kw)
	fmt.Println(list)
	this.display()
}

func (this *IndexController) Page() {
	this.display()
}

type BookLinkInfo struct {
	Name string
	Url  string
}

func (this *IndexController) Details() {
	unique_id := this.Ctx.Input.Param(":unique_id")
	data, err := models.NovelGetByUniqueId(unique_id)
	if err == nil {
		row := make(map[string]interface{})
		fName := models.ItemGetNameById(data.FromId)
		row["Id"] = data.Id
		row["Name"] = data.Name
		row["Desc"] = data.Desc
		row["Author"] = data.Author
		row["Category"] = data.Category
		row["FromId"] = data.FromId
		row["FromName"] = fName
		row["UniqueId"] = data.UniqueId
		row["Status"] = data.Status
		row["UpdateTime"] = beego.Date(time.Unix(data.UpdateTime, 0), "Y-m-d H:i:s")
		var bli []BookLinkInfo
		errJson := json.Unmarshal([]byte(data.List), &bli)
		// fmt.Println(bli)
		if errJson == nil {
			row["List"] = bli
		}
		// fmt.Println(data, err)
		this.Data["info"] = row
	} else {
		this.redirect("/")
	}

	this.display()
}
