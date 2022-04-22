package tailfilesource

import (
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

var (
	TailObj *tail.Tail
)

func Init(filename string) (err error) {

	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	TailObj, err = tail.TailFile(filename, config)
	if err != nil {
		logrus.Error("tail file failed,error:", err)
		return
	}
	return
}
