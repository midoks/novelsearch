package backends

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/midoks/webcron/app/lib"
	"github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/models"
	"strconv"
	"strings"
	"time"
)

type SysFuncController struct {
	BaseController
}

func (this *SysFuncController) Index() {

	result := models.FuncGetList()

	//对(栏目名)填充内容,利于后台观看
	for i := 0; i < len(result); i++ {
		for ci := 0; ci < len(result[i].List); ci++ {
			//println(result[i].List[ci].Name, len(result[i].List[ci].Name))
			fillcount := 20 - len(result[i].List[ci].Name)
			if fillcount > 0 && len(result[i].List[ci].Name) < 16 {
				tmp := strings.Repeat(" ", 20-len(result[i].List[ci].Name))
				result[i].List[ci].Name = fmt.Sprintf("%s%s", result[i].List[ci].Name, tmp)
			}
		}
	}

	this.Data["list"] = result
	this.display()
}

func (this *SysFuncController) Add() {

	row := new(models.SysFunc)
	id, err := this.GetInt("id")

	if err == nil {
		row, _ = models.FuncGetById(id)
	}
	this.Data["row"] = row

	if this.isPost() {

		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		row.Name = vars["name"]
		row.Pid, _ = strconv.Atoi(vars["pid"])
		row.Controller = vars["controller"]
		row.Action = vars["action"]
		row.Type, _ = strconv.Atoi(vars["type"])
		row.Icon = vars["icon"]
		row.Desc = vars["desc"]

		if vars["is_menu"] == "on" {
			row.IsMenu = 1
		} else {
			row.IsMenu = -1
		}

		row.Sort, _ = strconv.Atoi(vars["sort"])

		if id > 0 {

			row.UpdateTime = time.Now().Unix()
			err := row.Update()
			if err == nil {
				msg := fmt.Sprintf("更新功能的ID:%d|%s", id, row)
				this.uLog(msg)
				this.redirect(beego.URLFor("SysFuncController.Index"))
			}

		} else {

			row.Status = 0
			row.UpdateTime = time.Now().Unix()
			row.CreateTime = time.Now().Unix()
			id, err := orm.NewOrm().Insert(row)
			if err == nil {
				msg := fmt.Sprintf("添加功能的ID:%d|:%s", id, row)
				this.uLog(msg)
				this.redirect(beego.URLFor("SysFuncController.Index"))
			}
		}
	}

	listRow, _ := models.FuncGetListByPid(0)
	this.Data["listRow"] = listRow

	this.display()
}

func (this *SysFuncController) Lock() {

	id, err := this.GetInt("id")
	if err == nil {
		data, _ := models.FuncGetById(id)

		if data.Status > 0 {
			data.Status = -1
			this.uLog("功能锁定成功")
		} else {
			data.Status = 1
			this.uLog("功能解锁成功")
		}
		err = data.Update()

		if err == nil {
			this.retOk("功能修改成功")
		}
	}
	this.retFail("功能修改失败")
}

func (this *SysFuncController) Setmenu() {

	id, err := this.GetInt("id")
	if err == nil {
		data, _ := models.FuncGetById(id)

		if data.IsMenu > 0 {
			data.IsMenu = -1
			this.uLog("禁用菜单成功")
		} else {
			data.IsMenu = 1
			this.uLog("显示菜单成功")
		}
		err = data.Update()

		if err == nil {
			this.retOk("是否显示菜单修改成功")
		}
	}
	this.retFail("是否显示菜单修改失败")
}

func (this *SysFuncController) Del() {

	id, err := this.GetInt("id")
	if err == nil {
		num, err := models.FuncDelById(id)
		if err == nil {
			msg := fmt.Sprintf("删除功能ID:%d成功", num)
			this.uLog(msg)
			this.retOk(msg)
		}
	}
	this.retFail("非法参数")
}

func (this *SysFuncController) Sort() {

	id, err := this.GetInt("id")
	stype := this.GetString("type")

	if err == nil {
		data, _ := models.FuncGetById(id)
		if stype == "up" {
			data.Sort--
		} else if stype == "down" {
			data.Sort++
		}

		dataErr := data.Update()
		if dataErr == nil {
			msg := fmt.Sprintf("功能ID:%d,排序:%s,更新%d成功", data.Id, stype, data.Sort)
			this.uLog(msg)
			this.retOk(msg)
		}
	}
	this.retFail("非法参数")
}
