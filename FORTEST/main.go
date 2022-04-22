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
}
