package backends

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/crontab"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type AppItemController struct {
	BaseController
}

func (this *AppItemController) Index() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	searchType := this.GetString("search_type", "")
	searchWord := this.GetString("search_word", "")
	filters := make([]interface{}, 0)

	if searchType != "" {
		if strings.EqualFold(searchType, "msg") {
			searchType2 := fmt.Sprintf("%s__icontains", searchType)
			filters = append(filters, searchType2, searchWord)
		} else {
			filters = append(filters, searchType, searchWord)
		}
	}

	result, count := models.ItemGetList(page, this.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		row["Id"] = v.Id
		row["Name"] = v.Name
		row["PageIndex"] = v.PageIndex
		row["IsOfficial"] = v.IsOfficial
		row["Status"] = v.Status
		row["SpiderProgress"] = v.SpiderProgress
		row["UpdateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["CreateTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")

		list[k] = row
	}

	this.Data["search_type"] = searchType
	this.Data["search_word"] = searchWord
	this.Data["list"] = list
	this.Data["pageLink"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("SysUserController.Index"), true).ToString()
	this.display()
}

func (this *AppItemController) SearchAjax() {
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

func (this *AppItemController) Add() {

	data := new(models.AppItem)
	id, err := this.GetInt("id")

	if err == nil {
		data, _ = models.ItemGetById(id)
		//fmt.Println(id, data)
	}

	if this.isPost() {

		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		data.Name = vars["name"]
		data.IsOfficial = vars["is_official"]
		data.PageIndex = vars["page_index"]
		data.PageCharset = vars["page_charset"]
		data.PageIndexRule = vars["page_index_rule"]
		data.PathRule = vars["path_rule"]
		data.PathTpl = vars["path_tpl"]
		data.PathPageExp = vars["path_page_exp"]
		data.NameRule = vars["name_rule"]
		data.DescRule = vars["desc_rule"]
		data.AuthorRule = vars["author_rule"]
		data.CategoryRule = vars["category_rule"]
		data.StatusRule = vars["status_rule"]
		data.StatusEndMark = vars["status_end_mark"]
		data.ChapterPathRule = vars["chapter_path_rule"]
		data.ChapterPathExp = vars["chapter_path_exp"]
		data.ChapterListRule = vars["chapter_list_rule"]
		data.ContentExp = vars["content_exp"]
		data.ContentRule = vars["content_rule"]
		data.SosoExp = vars["soso_exp"]
		data.SosoPageArgs = vars["soso_page_args"]
		data.SosoRule = vars["soso_rule"]
		data.SpiderExp = vars["spider_exp"]
		data.SpiderRange = vars["spider_range"]
		data.SpiderRule = vars["spider_rule"]

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

func (this *AppItemController) Lock() {

	id, err := this.GetInt("id")
	if err == nil {
		data, _ := models.ItemGetById(id)

		if data.Status > 0 {
			data.Status = -1
			this.uLog("Item锁定成功")
		} else {
			data.Status = 1
			this.uLog("Item解锁成功")
		}
		err = data.Update()

		if err == nil {
			this.retOk("修改成功")
		}
	}
	this.retFail("修改失败")
}

func (this *AppItemController) Verify() {

	if this.isPost() {
		var url = ""
		var rule = ""
		var charset = ""
		var content = ""
		url = this.GetString("url", "")
		rule = this.GetString("rule", "")
		charset = this.GetString("charset", "")
		model := this.GetString("model", "")
		path_tpl := this.GetString("path_tpl", "")
		url = strings.TrimSpace(url)
		rule = strings.TrimSpace(rule)
		charset = strings.TrimSpace(charset)

		if strings.EqualFold(charset, "") {
			charset = "utf8"
		}

		if strings.EqualFold(url, "") {
			this.retFail("url不能为空")
			return
		}

		if !libs.IsUrlRe(url) {
			this.retFail("url不合法!")
			return
		}

		if strings.EqualFold(rule, "") {
			this.retFail("rule不能为空")
			return
		}

		req := httplib.Get(url)
		// req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		content, err := req.String()
		if err != nil {
			fmt.Println(err)
			this.retFail("获取url地址数据失败!")
			return
		}

		if strings.EqualFold(charset, "gbk") {
			content = libs.ConvertToString(content, "gbk", "utf8")
		}

		valid := regexp.MustCompile(rule)
		match := valid.FindAllStringSubmatch(content, -1)

		// fmt.Println(model)
		if strings.EqualFold(model, "1") {
			for i := 0; i < len(match); i++ {
				// fmt.Println(path_tpl, "{$ID}", match[i][1])
				match[i][1] = strings.Replace(path_tpl, "{$ID}", match[i][1], -1)
			}
		}
		// fmt.Println(url, content, match)
		this.retOk("验证成功", match)
	} else {
		this.retFail("非法请求")
	}
}

func (this *AppItemController) Del() {

	id, err := this.GetInt("id")
	if err == nil {
		num, err := models.ItemDelById(id)
		if err == nil {
			msg := fmt.Sprintf("删除item项目%s成功", num)
			this.uLog(msg)
			this.retOk(msg)
		}
	}
	this.retFail("非法参数")
}

func (this *AppItemController) AllSpider() {

	id, err := this.GetInt("id")
	if err == nil {
		r, err := models.ItemGetById(id)
		if err == nil {
			if r.Status != 1 {
				this.retFail("状态锁定中,无法操作!!")
				return
			}

			if r.SpiderExp != "" && r.SpiderRange != "" && r.SpiderRule != "" && r.PathTpl != "" {
				go crontab.CronWebRuleSpider(r, r.SpiderExp, r.SpiderRange, r.SpiderRule, r.PathTpl)
			} else {
				this.retFail("全站更新(条件不足)")
			}
			this.retOk("执行成功")
		}
	}
	this.retFail("非法参数")
}

func (this *AppItemController) OneSpider() {

	id, err := this.GetInt("id")
	url := this.GetString("url", "")
	name := this.GetString("name", "")
	if err == nil {
		r, err := models.ItemGetById(id)
		if err == nil {

			if r.Status != 1 {
				this.retFail("状态锁定中,无法操作!!")
				return
			}

			if !libs.IsUrlRe(url) {
				this.retFail("url地址不合法!")
			}

			if url != "" && name != "" {
				crontab.CronPathInfo(r, url, name)
			} else {
				this.retFail("全站更新(条件不足)")
			}
			this.retOk("执行成功")
		}
	}
}
