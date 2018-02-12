package helper

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSubstr(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	str := "党的领导是中国特色社会主义最本质的特征松树番茄,谁喜欢吃西红柿"
	fmt.Println(Substr(str, 2, 6))  //领导是中
	fmt.Println(Substr(str, 2, 60)) //领导是中国特色社会主义最本质的特征松树番茄,谁喜欢吃西
	fmt.Println(Substr(str, 20, 6)) //""
}

func TestRandomStr(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	fmt.Println(RandomStr(5))  //hHnZV
	fmt.Println(RandomStr(10)) //3X4gPDCu2y
}

func TestBase62(t *testing.T) {
	i := 349879
	b62 := Base62Encode(i)
	fmt.Println(b62)
	fmt.Println(Base62Decode(b62))
}