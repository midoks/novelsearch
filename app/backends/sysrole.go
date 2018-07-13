package backends

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/models"
	"strings"
	"time"
)

type SysRoleController struct {
	BaseController
}

func (this *SysRoleController) Index() {

	list, _ := models.RoleGetAll()
	this.Data["list"] = list

	this.display()
}

func (this *SysRoleController) Add() {

	data := new(models.SysRole)
	id, err := this.GetInt("id")
	if err == nil {
		data, _ = models.RoleGetById(id)
	}

	this.Data["data"] = data

	if this.isPost() {

		vars := make(map[string]string)
		box := make([]string, 0, 10000)
		this.Ctx.Input.Bind(&vars, "vars")
		this.Ctx.Input.Bind(&box, "box")

		data.Name = vars["name"]
		data.Desc = vars["desc"]
		data.List = strings.Join(box, ",")

		if id > 0 {

			data.UpdateTime = time.Now().Unix()
			err := data.Update()
			if err == nil {
				msg := fmt.Sprintf("更新角色的ID:%d|%s", id, data)
				this.uLog(msg)
				this.redirect(beego.URLFor("SysRoleController.Index"))
			}

		} else {

			data.Status = 0
			data.UpdateTime = time.Now().Unix()
			data.CreateTime = time.Now().Unix()

			id, err := orm.NewOrm().Insert(data)
			if err == nil {
				msg := fmt.Sprintf("添加角色的ID:%d|:%s", id, data)
				this.uLog(msg)
				this.redirect(beego.URLFor("SysRoleController.Index"))
			}
		}
	}

	funcList := models.FuncGetList()
	this.Data["funcList"] = funcList
	this.display()
}

func (this *SysRoleController) Lock() {

	id, err := this.GetInt("id")
	if err == nil {
		data, _ := models.RoleGetById(id)

		if data.Status > 0 {
			data.Status = -1
			msg := fmt.Sprintf("角色ID:%d,锁定成功", data.Id)
			this.uLog(msg)
		} else {
			data.Status = 1
			msg := fmt.Sprintf("角色ID:%d,解锁成功", data.Id)
			this.uLog(msg)
		}

		err = data.Update()
		if err == nil {
			this.retOk("角色修改成功")
		}
	}
	this.retFail("非法请求")
}

func (this *SysRoleController) Del() {

	id, err := this.GetInt("id")
	if err == nil {
		num, err := models.RoleDelById(id)
		if err == nil {
			msg := fmt.Sprintf("删除角色ID:%d成功", num)
			this.uLog(msg)
			this.retOk(msg)
		}
	}
	this.retFail("非法参数")
}
