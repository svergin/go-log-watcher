package logwatch

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func Start() {
	t, err := tail.TailFile("D:\\test.log", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		fmt.Println(line.Text)

	}

}
