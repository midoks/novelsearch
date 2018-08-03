package crontab

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	"time"
)

func checkFileFunc() error {
	file_cache := beego.AppConfig.String("file_cache")

	file_cache_expire, err := beego.AppConfig.Int64("file_cache_expire")
	if err != nil {
		file_cache_expire = 86400
	}
	// fmt.Printf("过期时间(%d)s\n", file_cache_expire)

	files, _ := ioutil.ReadDir(file_cache)
	for _, f := range files {
		var fn = file_cache + "/" + f.Name()
		fileInfo, _ := os.Stat(fn)

		if fileInfo.IsDir() {
			checkFileChildFunc(fn, file_cache_expire)
		} else {
			modTime := fileInfo.ModTime()
			expire_time := time.Now().Unix() - modTime.Unix() - file_cache_expire
			if expire_time > 0 {
				// fmt.Println(modTime.Unix(), time.Now().Unix())
				del := os.Remove(fn)
				if del != nil {
					logs.Warn("删除失败(%s)[%d]!!!", fn, expire_time)
					continue
				}
				logs.Warn("文件过期(%s)[%d]!!!", fn, expire_time)
			}
		}
	}
	return nil
}

func checkFileChildFunc(path string, file_cache_expire int64) {
	files, _ := ioutil.ReadDir(path)
	// fmt.Println(path)
	if len(files) == 0 {
		delf := os.Remove(path)
		if delf != nil {
			logs.Warn("删除失败!!!%s", path)
		}
		logs.Warn("空文件夹(%s)删除!!!", path)
		return
	}

	for _, f := range files {
		var fn = path + "/" + f.Name()
		fileInfo, _ := os.Stat(fn)

		if fileInfo.IsDir() {
			checkFileChildFunc(fn, file_cache_expire)
		} else {
			modTime := fileInfo.ModTime()
			if time.Now().Unix()-modTime.Unix() > file_cache_expire {
				// fmt.Println(modTime.Unix(), time.Now().Unix())
				del := os.Remove(fn)
				if del != nil {
					logs.Warn("删除失败!!!")
				}
			}
		}
	}
}
