package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	obj := reflect.TypeOf(x)
	fmt.Printf("%v,  %v,  %v\n", obj, obj.Name(), obj.Kind())
}

type person struct {
	name string
	age  int
}

func main() {
	var x = new(person)
	x.name = "zhangsan"
	x.age = 13
	reflectType(x)
}
