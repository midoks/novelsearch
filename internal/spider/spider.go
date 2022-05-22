package spider

import (
	// "fmt"
	"os"
	"path/filepath"
)

const RULE_DIR = "rule"

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Init(customConf string) {
	customDir := filepath.Dir(customConf)
	appDir := filepath.Dir(customDir) + "/" + RULE_DIR

	if !PathExists(appDir) {
		os.MkdirAll(appDir, os.ModePerm)
	}

	expInit(appDir)
}
