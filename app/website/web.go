package website

import (
	"fmt"
)

type WebInfo interface {
	Speak() string
}

func Init() {

	v := map[int]interface{}{}

	v[0] = Web23usSo{}

	fmt.Println(v)

	fmt.Println("website init")
}
