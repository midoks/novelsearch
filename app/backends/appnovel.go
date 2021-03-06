package backends

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/crontab"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"strconv"
	"time"
)

type AppNovelController struct {
	BaseController
}

func (this *AppNovelController) Index() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	searchType := this.GetString("search_type", "")
	searchWord := this.GetString("search_word", "")
	filters := make([]interface{}, 0)

	if searchType != "" {
		if libs.CheckStringIsExist(searchType, []string{"name", "author"}) {
			searchType2 := fmt.Sprintf("%s__icontains", searchType)
			filters = append(filters, searchType2, searchWord)
		} else {
			filters = append(filters, searchType, searchWord)
		}
	}

	result, count := models.NovelGetList(page, this.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		row["Id"] = v.Id
		row["Name"] = v.Name
		row["Desc"] = v.Desc
		row["Author"] = v.Author
		row["FromId"] = v.FromId
		row["List"] = v.List
		row["ChapterNum"] = v.ChapterNum
		row["LastChapter"] = v.LastChapter
		row["LastChapterUrl"] = v.LastChapterUrl
		row["Status"] = v.Status
		row["UpdateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["CreateTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}

	this.Data["search_type"] = searchType
	this.Data["search_word"] = searchWord
	this.Data["list"] = list
	this.Data["pageLink"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("AppNovelController.Index"), true).ToString()
	this.display()
}

func (this *AppNovelController) SearchAjax() {
	page := 1
	filters := make([]interface{}, 0)

	qstr := this.GetString("q")
	if qstr == "" {
		//this.retFail("搜索词不能为空")
	} else {
		q, err := strconv.Atoi(qstr)
		filters = append(filters, "status", 1)
		if err == nil {
			filters = append(filters, "id", q)
		} else {
			filters = append(filters, "name__icontains", qstr)
		}
	}

	result, _ := models.ItemGetList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {

		row := make(map[string]interface{})

		row["Id"] = v.Id
		row["Name"] = v.Name

		row["Status"] = v.Status
		row["UpdateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["CreateTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}
	this.retOk("ok", list)
}

func (this *AppNovelController) Add() {

	data := new(models.AppItem)
	id, err := this.GetInt("id")

	if err == nil {
		data, _ = models.ItemGetById(id)
	}

	if this.isPost() {

		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		data.Name = vars["name"]

		if id > 0 {

			data.UpdateTime = time.Now().Unix()
			err := data.Update()
			if err == nil {
				msg := fmt.Sprintf("更新Item的ID:%d|%s", id, data)
				this.uLog(msg)
				this.redirect(beego.URLFor("AppItemController.Index"))
			}
		} else {

			data.Status = 0
			data.UpdateTime = time.Now().Unix()
			data.CreateTime = time.Now().Unix()

			id, err := orm.NewOrm().Insert(data)
			if err == nil {
				msg := fmt.Sprintf("添加Item的ID:%d", id)
				this.uLog(msg)
				this.redirect(beego.URLFor("AppItemController.Index"))
			}
		}
	}

	this.Data["data"] = data
	this.Data["id"] = this.GetString("id")

	roleList, _ := models.RoleGetAll()
	this.Data["roleList"] = roleList

	this.display()
}

func (this *AppNovelController) Lock() {

	id, err := this.GetInt("id")
	if err == nil {
		data, _ := models.NovelGetById(id)

		if data.Status > 0 {
			data.Status = -1
		} else {
			data.Status = 1
		}
		err = data.Update()
		if err == nil {
			this.retOk("锁定成功")
		}
	}
	this.retFail("锁定失败")
}

func (this *AppNovelController) Del() {

	id, err := this.GetInt("id")
	if err == nil {
		num, err := models.NovelGetById(id)
		if err == nil {
			this.retOk(fmt.Sprintf("删除ID:%s成功", num))
		}
	}
	this.retFail("非法参数")
}

func (this *AppNovelController) Spider() {

	id, err := this.GetInt("id")
	if err == nil {
		novel, err := models.NovelGetById(id)
		r, err := models.ItemGetById(novel.FromId)
		if err == nil {
			if r.Status != 1 {
				this.retFail("状态锁定中,无法操作!!")
			}

			if novel.Name != "" && novel.Url != "" {
				go crontab.CronNovelUpdate(r, novel, novel.Url, novel.Name)
			} else {
				this.retFail("全站更新(条件不足)")
			}
			this.retOk("执行成功")
		}
	}
	this.retFail("非法参数")
}
