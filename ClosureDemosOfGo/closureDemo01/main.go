package main

import "fmt"

func f1(a int) func() int {
	var n = 3
	fmt.Println(n)
	return func() int {
		a++
		return a
	}
}

func main() {
	a := f1(1)

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
