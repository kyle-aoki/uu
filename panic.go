package uu

import "fmt"

func MainRecover() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
