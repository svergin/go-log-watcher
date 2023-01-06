package internal

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func WatchLog() {
	t, err := tail.TailFile("D:\\test.log", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		fmt.Println(line.Text)

	}

}
