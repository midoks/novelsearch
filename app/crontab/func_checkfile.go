package crontab

import (
	// "bytes"
	"fmt"
	// "github.com/midoks/novelsearch/app/libs"
	// "github.com/midoks/novelsearch/app/models"
	// "html/template"
	// "strings"
	"io/ioutil"
	"os"
)

func checkFileFunc() error {
	conn, err := os.Open(".cache")
	if err != nil {
		fmt.Println(err)
	}

	fd, err := ioutil.ReadAll(conn)
	fmt.Println(fd)
	return nil
}
