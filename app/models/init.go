package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	// "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"net/url"
	"time"
)

const (
	WEB_NAME        = "web_name"
	WEB_KEYWORD     = "web_keyword"
	WEB_DESC        = "web_desc"
	WEB_STAT        = "web_stat"
	WEB_NOTICE      = "web_notice"
	WEB_NOTICE_MAIL = "web_notice_mail"

	MAIL_HOST       = "mail_host"
	MAIL_PORT       = "mail_port"
	MAIL_USER       = "mail_user"
	MAIL_PWD        = "mail_pwd"
	MAIL_QUEUE_SIZE = "mail_queue_size"
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

	orm.RegisterModel(new(SysUser), new(SysFunc), new(SysRole), new(SysLog), new(SysOption),
		new(AppItem), new(AppDebug), new(AppNovel))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

}

func MysqlPing() bool {
	o := orm.NewOrm()
	r := o.Raw("show tables")
	fmt.Println(r)

	return false
}

func getCacheIntance() (cache.Cache, error) {
	cacheLink, err := cache.NewCache("memory", `{"interval":300}`)
	if err != nil {
		return nil, err
	}
	return cacheLink, nil
}

func getCache(key string) (bool, string) {
	cacheLink, err := getCacheIntance()
	if err == nil {
		val := cacheLink.Get(key)
		fmt.Println(val)
		if val != nil {
			return true, val.(string)
		}
	}
	return false, ""
}

func setCache(key string, val interface{}, timeout time.Duration) bool {
	cacheLink, err := getCacheIntance()
	if err == nil {
		err2 := cacheLink.Put(key, val, timeout)
		if err2 == nil {
			return true
		}
	}
	return false
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
