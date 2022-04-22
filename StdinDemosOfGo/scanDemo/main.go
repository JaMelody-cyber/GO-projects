package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(input)
}
