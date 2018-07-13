package backends

import (
	"fmt"
	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "github.com/midoks/webcron/app/libs"
	"github.com/midoks/novelsearch/app/models"
	// "strconv"
	// "strings"
	// "time"
)

type ItemController struct {
	CommonController
}

func (this *ItemController) Index() {
	this.retOk("ok")
}

func (this *ItemController) Check() {
	if this.isPost() {
		this.retOk("ok")
	}
	this.retFail("非法请求")
}

//清除该项目所有计划任务
func (this *ItemController) ClearAll() {
	if this.isPost() {
		item_sign := this.GetString("item_sign")
		fmt.Println(item_sign)
		if item_sign == "" {
			this.retFail("项目标示(item_sign)不能为空")
		}

		list, count := models.ItemGetList(1, 10, "sign", item_sign)

		if count > 0 {
			this.retOk("ok", list[0])
		}
		this.retFail("该项目不存在")
	}
	this.retFail("非法请求")
}

func (this *ItemController) Add() {
	if this.isPost() {

		this.retOk("ok")
	}
	this.retFail("非法请求")
}
