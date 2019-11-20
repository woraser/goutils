package common

import (
	"fmt"
	"testing"
)

func BenchmarkGetUUId(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid, err := GetUUId()
		if err != nil {
			fmt.Println("err:", err)
			break
		}
		fmt.Println("uuid is:", uuid)
	}
}
