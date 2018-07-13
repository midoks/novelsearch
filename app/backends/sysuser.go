package backends

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"strconv"
	"strings"
	"time"
)

type SysUserController struct {
	BaseController
}

func (this *SysUserController) Index() {
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

	result, count := models.UserGetList(page, this.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {

		row := make(map[string]interface{})

		roleData, roleErr := models.RoleGetById(v.Roleid)

		row["Id"] = v.Id
		row["Username"] = v.Username
		row["Nick"] = v.Nick
		row["Sex"] = v.Sex
		row["Mail"] = v.Mail
		row["Tel"] = v.Tel
		row["Roleid"] = v.Roleid
		if roleErr != nil || roleData.Name == "" {
			row["Rolename"] = "无权限"
		} else {
			row["Rolename"] = roleData.Name
		}

		row["Status"] = v.Status
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

func (this *SysUserController) Repwd() {

	if this.isPost() {
		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")
		tmpUser := this.user

		if vars["password"] != "" {
			tmpUser.Password = libs.Md5([]byte(vars["password"]))
		}

		tmpUser.Nick = vars["nick"]
		tmpUser.Mail = vars["mail"]
		tmpUser.Tel = vars["tel"]
		// tmpUser.Roleid, _ = strconv.Atoi(vars["roleid"])
		tmpUser.Sex, _ = strconv.Atoi(vars["sex"])
		err := tmpUser.Update()
		if err == nil {
			msg := fmt.Sprintf("修改信息:%s", tmpUser)
			this.uLog(msg)
			this.redirect(beego.URLFor("SysUserController.Index"))
		}
	}
	this.display()
}

func (this *SysUserController) Add() {

	data := &models.SysUser{}
	id, err := this.GetInt("id")

	if err == nil {
		data, _ = models.UserGetById(id)
	}

	if this.isPost() {

		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		if vars["password"] != "" {
			data.Password = libs.Md5([]byte(vars["password"]))
		}

		data.Nick = vars["nick"]
		data.Mail = vars["mail"]
		data.Tel = vars["tel"]
		data.Username = vars["username"]
		data.Roleid, _ = strconv.Atoi(vars["roleid"])
		data.Sex, _ = strconv.Atoi(vars["sex"])

		if id > 0 {

			data.UpdateTime = time.Now().Unix()
			err := data.Update()
			if err == nil {
				msg := fmt.Sprintf("更新用户的ID:%d|%s", id, data)
				this.uLog(msg)
				this.redirect(beego.URLFor("SysUserController.Index"))
			}
		} else {

			data.Status = 0
			data.UpdateTime = time.Now().Unix()
			data.CreateTime = time.Now().Unix()

			id, err := orm.NewOrm().Insert(data)

			fmt.Println(err)
			if err == nil {
				msg := fmt.Sprintf("添加用户的ID:%d", id)
				this.uLog(msg)
				this.redirect(beego.URLFor("SysUserController.Index"))
			}
		}
	}

	this.Data["data"] = data
	this.Data["id"] = this.GetString("id")

	roleList, _ := models.RoleGetAll()
	this.Data["roleList"] = roleList

	this.display()
}

func (this *SysUserController) Lock() {

	id, err := this.GetInt("id")
	if err == nil {
		data, _ := models.UserGetById(id)

		if data.Status > 0 {
			data.Status = -1
			this.uLog("锁定成功")
		} else {
			data.Status = 1
			this.uLog("解锁成功")
		}
		err = data.Update()

		if err == nil {
			this.retOk("修改成功")
		}
	}
	this.retFail("修改失败")
}

func (this *SysUserController) Del() {

	id, err := this.GetInt("id")
	if err == nil {
		num, err := models.UserDelById(id)
		if err == nil {
			msg := fmt.Sprintf("删除用户%s成功", num)
			this.uLog(msg)
			this.retOk(msg)
		}
	}
	this.retFail("非法参数")
}
