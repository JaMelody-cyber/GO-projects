package main

import "fmt"

func main() {
	a := new(int)
	transferFunc(a)
	fmt.Printf("%p\n", &a)
}
func transferFunc(a *int) {
	fmt.Printf("%p\n", &a)
}
