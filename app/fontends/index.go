package fontends

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/crontab"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"strconv"
	"strings"
	"time"
)

type IndexController struct {
	CommonController
}

type BookLinkInfo struct {
	Name string
	Url  string
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

func (this *IndexController) Baidutop() {
	list, err := libs.GetAllBaiduTop()
	if err == nil {
		this.Data["list"] = list
	}
	this.display()
}

func (this *IndexController) Soso() {
	kw := this.GetString("wd", "")
	if strings.EqualFold(kw, "") {
		this.redirect("/")
	}

	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}
	pageSize := 10

	result, count := models.SosoNovelByKw(kw, page, pageSize)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})
		fName := models.ItemGetNameById(v.FromId)

		row["Id"] = v.Id
		row["Name"] = v.Name
		row["UniqueId"] = v.UniqueId
		row["Author"] = v.Author
		row["FromId"] = v.FromId
		row["FromName"] = fName
		row["LastChapter"] = v.LastChapter
		row["LastChapterUrl"] = v.LastChapterUrl
		row["Status"] = v.Status
		row["UpdateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d")
		list[k] = row
	}

	this.Data["pageLink"] = libs.NewPager(page, int(count), pageSize, "/s?wd="+kw, true).ToString()
	this.Data["list"] = list
	this.Data["kw"] = kw
	this.display()
}

func (this *IndexController) Content() {
	var err error
	id := this.Ctx.Input.Param(":id")
	chapter_id := this.Ctx.Input.Param(":chapter_id")

	id_real, err := libs.Base64decode(id)
	if err != nil {
		this.redirect("/")
		return
	}

	chapter_id_real, err := libs.Base64decode(chapter_id)
	if err != nil {
		this.redirect("/")
		return
	}

	chapter_id_int, err := strconv.Atoi(chapter_id_real)
	if err != nil {
		this.redirect("/")
		return
	}

	data, err := models.NovelGetByIdStr(id_real)

	if err != nil {
		this.redirect("/")
		return
	}

	ruleList, err := models.ItemGetById(data.FromId)
	if err != nil {
		this.redirect("/")
		return
	}

	var bli []BookLinkInfo
	err = json.Unmarshal([]byte(data.List), &bli)
	if err != nil {
		this.redirect("/")
		return
	}

	info := bli[chapter_id_int]
	url := info.Url

	urlData, err := libs.GetHttpData(url)
	if err != nil {

	}

	content, err := crontab.RegNovelSigleInfo(urlData, ruleList.ContentRule)
	if err != nil {

	}

	chapter_id_float, err := strconv.ParseFloat(chapter_id_real, 64)
	count_list := strconv.Itoa(len(bli))
	count_list_float, err := strconv.ParseFloat(count_list, 64)
	this.Data["Percent"] = fmt.Sprintf("%.2f", (chapter_id_float/count_list_float)*100)

	tmpPrev := chapter_id_int - 1
	if tmpPrev >= 0 {
		this.Data["Prev"] = tmpPrev
	}

	tmpNext := chapter_id_int + 1
	if tmpNext <= len(bli) {
		this.Data["Next"] = tmpNext
	}

	this.Data["Info"] = data
	this.Data["Title"] = info.Name
	this.Data["Content"] = content
	this.display()
}

func (this *IndexController) List() {
	var err error
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
		err := json.Unmarshal([]byte(data.List), &bli)
		if err == nil {
			row["List"] = bli
		}
		this.Data["info"] = row
	} else {
		this.redirect("/")
	}
	this.display()
}
