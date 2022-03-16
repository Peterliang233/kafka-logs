package taillog

import (
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		return
	}
	return nil
}

// ReadChan 不断监听my.log文件里面的数据，一旦出现新的一条的数据，就立马扔到kafka里面
func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
