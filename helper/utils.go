/*
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     helper
 * @date        2018-01-25 19:19
 */
package helper

import (
	"unsafe"
	"strings"
	"os"
	"strconv"
	"bytes"
)

const N = int(unsafe.Sizeof(0))

func IsBigEndian() bool {
	x := 0x1234
	p := unsafe.Pointer(&x)
	p2 := (*[N]byte)(p)
	if p2[0] == 0 {
		return true
	}
	return false
}

//文件是否存在
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//是否全部为数字
func IsAllNumber(str string) bool {
	ret := strings.Trim(str, "0123456789")
	return len(ret) == 0
}

//字符串直接转为json
func StrToJSON(str string) string {
	var jsons bytes.Buffer
	for _, rn := range str {
		rint := int(rn)
		if rint < 128 {
			jsons.WriteRune(rn)
		} else {
			jsons.WriteString("\\u")
			jsons.WriteString(strconv.FormatInt(int64(rint), 16))
		}
	}
	return jsons.String()
}
