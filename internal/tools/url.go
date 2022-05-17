package tools

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func GetHttpData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("资源获取错误!")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return BytesToString(body), err
}
