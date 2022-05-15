package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"

	"github.com/midoks/novelsearch/internal/conf"
	"github.com/midoks/novelsearch/internal/log"
	"github.com/midoks/novelsearch/internal/mgdb"
	"github.com/midoks/novelsearch/internal/tools"
)

func autoMakeCustomConf(customConf string) error {

	if tools.IsExist(customConf) {
		return nil
	}

	// auto make custom conf file
	cfg := ini.Empty()
	if tools.IsFile(customConf) {
		if err := cfg.Append(customConf); err != nil {
			return err
		}
	}

	cfg.Section("").Key("app_name").SetValue("novelsearch")
	cfg.Section("").Key("run_mode").SetValue("prod")

	cfg.Section("web").Key("http_port").SetValue("11011")
	cfg.Section("session").Key("provider").SetValue("memory")

	cfg.Section("mongodb").Key("addr").SetValue("127.0.0.1:27017")
	cfg.Section("mongodb").Key("db").SetValue("novelsearch")

	cfg.Section("security").Key("install_lock").SetValue("true")
	secretKey := tools.RandString(15)
	cfg.Section("security").Key("secret_key").SetValue(secretKey)

	os.MkdirAll(filepath.Dir(customConf), os.ModePerm)
	if err := cfg.SaveTo(customConf); err != nil {
		return err
	}

	return nil
}

func Init(customConf string) error {
	var err error

	if customConf == "" {
		customConf = filepath.Join(conf.CustomDir(), "conf", "app.conf")
	} else {
		customConf, err = filepath.Abs(customConf)
		if err != nil {
			return fmt.Errorf("custom conf file get absolute path: %v", err)
		}
	}

	err = autoMakeCustomConf(customConf)
	if err != nil {
		return err
	}

	conf.Init(customConf)
	log.Init()
	mgdb.Init()

	return nil
}

func init() {
	Init("")
}
