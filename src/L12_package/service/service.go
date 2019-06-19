package service

import (
	"fmt"
)

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func Add (param ...int) int {
	result := 0;
	for v,_ := range param {
		result += v
	}
	return result
}

