package main

import (
	"GoProjects/FORTEST/pkg2"
	"fmt"
)

var (
	i int
)

func main() {
	fmt.Println(pkg2.Outer)
	pkg2.GetOuter()
	fmt.Println(pkg2.Outer)
	fmt.Println("it is a testing commit")
	fmt.Println("it is a testing commit2")
	fmt.Println("it is a testing commit3")
	fmt.Println("it is a testing commit4")
}
