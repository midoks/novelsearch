package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/url"
)

func Init() {
	fmt.Println("db init")

	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(new(SysUser), new(SysFunc), new(SysRole), new(SysLog),
		new(AppItem), new(AppDebug), new(AppNovel))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

}

func MysqlPing() bool {
	o := orm.NewOrm()
	r = o.Raw("mysql ping")
	fmt.Println(o)

	return false
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
