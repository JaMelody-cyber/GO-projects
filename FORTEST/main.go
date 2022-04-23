package main

import (
	"GoProjects/FORTEST/pkg2"
	"fmt"
)

func main() {
	fmt.Println(test01.Outer)
	test01.GetOuter()
	fmt.Println(test01.Outer)
	fmt.Println("it is a testing commit")
	fmt.Println("it is a testing commit--remote")
	fmt.Println("it is a testing commit--remote")
	strs := [...]string{"lisi", "zhangsan", "awang"}
	for _, x := range strs {
		fmt.Println(x)
	}
}
