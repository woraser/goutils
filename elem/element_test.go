package elem

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/kr/pretty"
	"runtime"
	"testing"
)

func TestNewItemElem(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	m := make(map[string]string)
	m["wendao"] = "aaaa"
	m["abc"] = "defffff"
	aa := []byte("wendaoaaaaaa")
	bytesBuffer := bytes.NewBuffer(aa)
	var btmp interface{}
	binary.Read(bytesBuffer, binary.BigEndian, &btmp)
	tmp := []interface{}{
		456.9875345, "wendao", "3.440000004", 235657789809,
		[]int{888, 999}, -999999, -123434654567, m,
		`{"wendao":55555,"44":{"b":9999,"rrrr":"23423424234"},"45":"aaaa"}`,
		btmp,
	}
	for _, d := range tmp {
		elem := MakeItemElem(d)
		fmt.Printf("OriginData:\t%v\n", elem.RawData())
		toBool, err := elem.ToBool()
		fmt.Printf("--toBool:\n\t%# v\n\t%# v\n", pretty.Formatter(toBool), err)
		toInt, err := elem.ToInt()
		fmt.Printf("--toInt:\n\t%# v\n\t%# v\n", pretty.Formatter(toInt), err)
		toInt8, err := elem.ToInt8()
		fmt.Printf("--toInt8:\n\t%# v\n\t%# v\n", pretty.Formatter(toInt8), err)
		toInt16, err := elem.ToInt16()
		fmt.Printf("--toInt16:\n\t%# v\n\t%# v\n", pretty.Formatter(toInt16), err)
		toInt32, err := elem.ToInt32()
		fmt.Printf("--toInt32:\n\t%# v\n\t%# v\n", pretty.Formatter(toInt32), err)
		toInt64, err := elem.ToInt64()
		fmt.Printf("--toInt64:\n\t%# v\n\t%# v\n", pretty.Formatter(toInt64), err)
		toUint, err := elem.ToUint()
		fmt.Printf("--toUint:\n\t%# v\n\t%# v\n", pretty.Formatter(toUint), err)
		toUint8, err := elem.ToUint8()
		fmt.Printf("--toUint8:\n\t%# v\n\t%# v\n", pretty.Formatter(toUint8), err)
		toUint16, err := elem.ToUint16()
		fmt.Printf("--toUint16:\n\t%# v\n\t%# v\n", pretty.Formatter(toUint16), err)
		toUint32, err := elem.ToUint32()
		fmt.Printf("--toUint32:\n\t%# v\n\t%# v\n", pretty.Formatter(toUint32), err)
		toUint64, err := elem.ToUint64()
		fmt.Printf("--toUint64:\n\t%# v\n\t%# v\n", pretty.Formatter(toUint64), err)
		toFloat32, err := elem.ToFloat32()
		fmt.Printf("--toFloat32:\n\t%# v\n\t%# v\n", pretty.Formatter(toFloat32), err)
		toFloat64, err := elem.ToFloat64()
		fmt.Printf("--toFloat64:\n\t%# v\n\t%# v\n", pretty.Formatter(toFloat64), err)
		toStr := elem.ToString()
		fmt.Printf("--toStr:%s\n", toStr)
		toSlice, err := elem.ToSlice()
		fmt.Printf("--toSlice%# v\n\t%# v\n", pretty.Formatter(convSlice(toSlice)), err)
		toMap, err := elem.ToMap()
		fmt.Printf("--toMap%# v\n\t%# v\n", pretty.Formatter(convMap(toMap)), err)
	}

}

func convSlice(param []ItemElem) []interface{} {
	args := make([]interface{}, len(param))
	for i := range param {
		args[i] = param[i].RawData()
	}
	return args
}
func convMap(d map[ItemElem]ItemElem) map[interface{}]interface{} {
	ret := make(map[interface{}]interface{})
	for k, v := range d {
		ret[k.RawData()] = v.RawData()
	}
	return ret
}
