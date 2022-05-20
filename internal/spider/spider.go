package spider

import (
	// "fmt"
	"os"
	"path/filepath"
)

const RULE_DIR = "rule"

func Init(customConf string) {
	customDir := filepath.Dir(customConf)
	appDir := filepath.Dir(customDir) + "/" + RULE_DIR
	os.MkdirAll(appDir, os.ModePerm)

	expInit(appDir)
}
