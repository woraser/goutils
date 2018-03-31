package slice

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestToSliceIface(t *testing.T) {
	str := "wendaosssss"
	ret := ToSliceIface(str)
	fmt.Println(ret)
}

func TestInSlice(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	var testSlice = []interface{}{444, "wendao", "555.6", "english", "科技"}
	//true false
	fmt.Println(InSlice(444, testSlice), InSlice("sssss", testSlice))
	test1 := []int{44, 666, 9, 23, 111}
	var ff float64 = 3.5555
	//false true true
	fmt.Println(InSlice(444, ToSliceIface(test1)),
		InSlice("wendao", ToSliceIface([]string{"wendao", "abc"})),
		InSlice(ff, ToSliceIface([]float64{3.7777, 7.13446456, 3.5555})))
}

func TestSliceMerge(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	var testSlice = []interface{}{444, "wendao", "555.6", "english", "科技"}
	test1 := []int{44, 666, 9, 23, 111}
	//[444 wendao 555.6 english 科技 44 666 9 23 111]
	fmt.Println(SliceMerge(testSlice, ToSliceIface(test1)))
}

func TestSliceFilter(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	var testSlice = []interface{}{444, "wendao", "555.6", "english", "科技"}
	ret := SliceFilter(testSlice, func(v interface{}) bool {
		val := reflect.ValueOf(v)
		return val.Kind() == reflect.String
	})
	//[wendao 555.6 english 科技]
	fmt.Println(ret)
}

func TestSliceReduce(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	var testSlice = []interface{}{444, "wendao", "555.6", "english", "科技"}
	ret := SliceReduce(testSlice, func(v interface{}) interface{} {
		val := reflect.ValueOf(v)
		return val.String()
	})
	//[<int Value> wendao 555.6 english 科技]
	fmt.Println(ret)
}
