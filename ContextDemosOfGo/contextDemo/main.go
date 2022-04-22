package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancle := context.WithCancel(context.Background())
	//go subNode(ctx)
	//time.Sleep(3 * time.Second)
	//cancle()
	//time.Sleep(1 * time.Second)
	withValueDemo()
}

func subNode(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("subNode函数已被取消")
			return
		default:
			fmt.Println("subNode函数正在执行")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
func withValueDemo() {
	type withValueType string
	var key = withValueType("testKey")
	data := "This is test Value"
	ctx := context.WithValue(context.Background(), key, data)

	func(ctx context.Context) {
		fmt.Println("第一层", ctx.Value(key))
		func(ctx context.Context) {
			fmt.Println("第二层", ctx.Value(key))
			func(ctx context.Context) {
				fmt.Println("第三层", ctx.Value(key))
			}(ctx)
		}(ctx)
	}(ctx)
}
