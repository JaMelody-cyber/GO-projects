package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filename := "E:\\GoProjects\\LogProjectTestDemo\\tailfDemo\\my.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tails, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("tail file failed,error:", err)
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("file is reopening")
			time.Sleep(time.Second * 100)
		}
		fmt.Println("msg:", msg.Text)
	}
}
