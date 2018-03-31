package file

import (
	"fmt"
	"testing"
)

func TestNewFileIterator(t *testing.T) {
	itor, err := NewFileIterator().SetFile("./snowflake.txt").Init()
	if err != nil {
		t.Errorf("%v", err)
	}

	//遍历每一行，写入库里
	itor.IterLine(func(line string) {
		fmt.Println(line)
	})
}
