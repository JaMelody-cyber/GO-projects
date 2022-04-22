package main

/*
有两种支付方式，支付宝和微信；
每加一种支付方式，都要写他的checkWithXXX，这块代码其实是重复（逻辑重复）的。

*/

import (
	"fmt"
	"reflect"
)

type aliPay struct {
	amount int
}
type wePay struct {
	amount int
}

func checkWithAliPay(obj *aliPay) {
	obj.pay()
}

func checkWithWePay(obj *wePay) {
	obj.pay()
}

func (a *aliPay) pay() {
	fmt.Printf("user has paid %v dolloar with %v\n", a.amount, reflect.TypeOf(a))
}
func (w *wePay) pay() {
	fmt.Printf("user has paid %v dolloar with %v\n", w.amount, reflect.TypeOf(w))
}
func main() {
	order_1 := &aliPay{100}
	if true {
		checkWithAliPay(order_1)
	}
	order_2 := &wePay{100}
	if true {
		checkWithWePay(order_2)
	}
}
