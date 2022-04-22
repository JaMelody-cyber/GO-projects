package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	jobChan := make(chan int64, 1)
	resultChan := make(chan int64, 100)
	go randNum(jobChan)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go sumNum(jobChan, resultChan)
	}
	close(resultChan)
	for ret := range resultChan {
		fmt.Println(ret)
	}
	wg.Wait()
}

func randNum(jobChan chan<- int64) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		jobChan <- rand.Int63()
	}
	close(jobChan)
}

func sumNum(jobChan <-chan int64, resultChan chan<- int64) {
	defer wg.Done()
	for v := range jobChan {
		var s int64
		for v > 0 {
			s += v % 10
			v = v / 10
		}
		resultChan <- s
	}
}
