package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type SysFunc struct {
	Id         int
	Name       string
	Pid        int
	Controller string
	Action     string
	Type       int
	IsMenu     int
	Icon       string
	Desc       string
	Sort       int
	Status     int
	UpdateTime int64
	CreateTime int64
}

type SysFuncNav struct {
	Info      SysFunc
	List      []SysFunc
	MenuOpen  bool
	ListCount int
}

func getTnByFunc() string {
	return "sys_func"
}

func (u *SysFunc) TableName() string {
	return getTnByFunc()
}

func (u *SysFunc) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func FuncGetNav(curController string, curAction string) (navNow []SysFuncNav, menuNameNow string, funcNameNow string, isAuthNow bool) {

	o := orm.NewOrm()
	var list []SysFunc

	res, _ := o.Raw("select * from sys_func where pid=? and status=? order by sort asc", 0, 1).QueryRows(&list)
	nav := make([]SysFuncNav, len(list))

	var curMenuName string = ""
	var curMenuFuncName string = ""
	var isAuth bool = false

	if res > 0 {
		for i := 0; i < len(list); i++ {
			var cList []SysFunc
			cres, _ := o.Raw("select * from sys_func where pid=? and status=? order by sort asc", list[i].Id, 1).QueryRows(&cList)
			if cres > 0 {
				nav[i].Info = list[i]
				nav[i].List = cList
				nav[i].ListCount = len(cList)
				nav[i].MenuOpen = false

				for ci := 0; ci < len(cList); ci++ {
					if strings.EqualFold(cList[ci].Controller, curController) && strings.EqualFold(cList[ci].Action, curAction) {
						nav[i].MenuOpen = true
						// fmt.Println("debug:", cList[ci].Controller, curController, cList[ci].Action, curAction)
						curMenuName = list[i].Name
						curMenuFuncName = cList[ci].Name

						isAuth = true
					}
				}
			}
		}
	}

	return nav, curMenuName, curMenuFuncName, isAuth
}

func FuncInGetNav(in string, curController string, curAction string) (navNow []SysFuncNav, menuNameNow string, funcNameNow string) {

	o := orm.NewOrm()
	var list []SysFunc

	sql := ""
	sql = fmt.Sprintf("select * from sys_func where pid=%d and `status`=%d order by sort asc", 0, 1)

	res, _ := o.Raw(sql).QueryRows(&list)
	nav := make([]SysFuncNav, len(list))
	var curMenuName string = ""
	var curMenuFuncName string = ""

	if res > 0 {
		for i := 0; i < len(list); i++ {
			var cList []SysFunc
			cres, _ := o.Raw("select * from sys_func where pid=? and `status`=? order by sort asc", list[i].Id, 1).QueryRows(&cList)
			if cres > 0 {
				nav[i].Info = list[i]
				nav[i].List = cList
				nav[i].ListCount = len(cList)
				nav[i].MenuOpen = false

				for ci := 0; ci < len(cList); ci++ {
					if strings.EqualFold(cList[ci].Controller, curController) && strings.EqualFold(cList[ci].Action, curAction) {
						nav[i].MenuOpen = true
						curMenuName = list[i].Name
						curMenuFuncName = cList[ci].Name
					}
				}
			}
		}
	}
	return nav, curMenuName, curMenuFuncName
}

func FuncGetList() []SysFuncNav {

	o := orm.NewOrm()
	var list []SysFunc

	o.Raw("select * from sys_func where pid=? order by sort asc", 0).QueryRows(&list)
	nav := make([]SysFuncNav, len(list))

	for i := 0; i < len(list); i++ {
		var cList []SysFunc
		o.Raw("select * from sys_func where pid=? order by sort asc", list[i].Id).QueryRows(&cList)

		nav[i].Info = list[i]
		nav[i].List = cList
		nav[i].ListCount = len(cList)
	}
	return nav
}

func FuncGetListByPid(pid int64) ([]SysFunc, error) {

	o := orm.NewOrm()
	var cList []SysFunc
	cres, cerr := o.Raw("select * from sys_func where pid=? and status=? order by sort asc", pid, 1).QueryRows(&cList)
	if cres > 0 {
		return cList, nil
	}
	return cList, cerr
}

func FuncGetById(id int) (*SysFunc, error) {

	sysfunc := new(SysFunc)
	sysfunc.Id = id

	err := orm.NewOrm().Read(sysfunc)
	if err != nil {
		return nil, err
	}
	return sysfunc, nil
}

func FuncDelById(id int) (int64, error) {
	return orm.NewOrm().Delete(&SysFunc{Id: id})
}
