package uu

import "fmt"

func Recover() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
