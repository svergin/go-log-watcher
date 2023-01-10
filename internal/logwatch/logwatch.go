package logwatch

import (
	"github.com/hpcloud/tail"
)

func Start(filename string) (*tail.Tail, error) {
	return tail.TailFile(filename, tail.Config{Follow: true})
	// if err != nil {
	// 	panic(err)
	// }
	// for line := range t.Lines {
	// 	fmt.Println(line.Text)

	// }

}
