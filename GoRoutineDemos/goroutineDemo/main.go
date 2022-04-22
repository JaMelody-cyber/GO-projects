package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello() {
	fmt.Println("hello !")
	wg.Done()
}

func main() {
	wg.Add(1)
	go hello()
	fmt.Println("main")
	wg.Wait()
	chan_int := make(chan int, 1)
	for i := 1; i < 10; i++ {
		select {
		case x := <-chan_int:
			fmt.Println(x)
		case chan_int <- i:
			{

			}
		case <-time.After(time.Second):
		}
	}
	k, v := <-chan_int
	fmt.Println(k, v)
}
