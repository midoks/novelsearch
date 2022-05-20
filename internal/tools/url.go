package tools

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func GetHttpData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("resource http get error!")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return BytesToString(body), err
}
