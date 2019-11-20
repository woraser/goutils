package common

import (
	"fmt"
	"testing"
)

func TestGetCurrentDirectory(t *testing.T) {
	path, err := GetCurrentDirectory()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("current directory is:", path)
}

func TestGetUUId(t *testing.T) {
	uuid, err := GetUUId()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("uuid is:", uuid)
}
