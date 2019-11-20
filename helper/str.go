/*
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     helper
 * @date        2018-01-25 19:19
 */
package helper

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"math"
	r "math/rand"
	"time"
)

func init() {
	//base62转码初始化
	for k, v := range base62CharToInt {
		base62IntToChar[v] = k
	}
}

var alphaNum = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

/**
    random generate str
	a :=RandomCreateBytes(16)
	fmt.Println(a)
*/
func RandomStr(n int, alphabets ...byte) string {
	if len(alphabets) == 0 {
		alphabets = alphaNum
	}
	var byteSlice = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(byteSlice); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range byteSlice {
		if randBy {
			byteSlice[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			byteSlice[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(byteSlice)
}

func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)
	if start < 0 {
		start += length
	} else if start > length {
		start = start % length
	}
	if end < 0 {
		end += length
	} else if end > length {
		end = end % length
	}
	if start > end || end < 0 || start < 0 {
		return ""
	}
	return string(rs[start:end])
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

var base62CharToInt = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var base62IntToChar = make(map[string]int)

func Base62Encode(num int64) string {
	baseStr := ""
	for {
		if num <= 0 {
			break
		}
		i := num % 62
		baseStr += base62CharToInt[i]
		num = (num - i) / 62
	}
	return baseStr
}

func Base62Decode(b62Str string) int64 {
	var rs int64 = 0
	for i := 0; i < len(b62Str); i++ {
		rs += int64(base62IntToChar[string(b62Str[i])]) * int64(math.Pow(62, float64(i)))
	}
	return rs
}

//半角字符转全角字符
func ToDBC(str string) string {
	ret := ""
	for _, rs := range str {
		if rs == 32 {
			ret += string(12288)
		} else if rs < 127 {
			ret += string(rs + 65248)
		} else {
			ret += string(rs)
		}
	}
	return ret
}

//全角字符转半角字符
func ToCBD(str string) string {
	ret := ""
	for _, rs := range str {
		if rs == 12288 {
			ret += string(rs - 12256)
			continue
		}
		if rs > 65280 && rs < 65375 {
			ret += string(rs - 65248)
		} else {
			ret += string(rs)
		}
	}
	return ret
}
