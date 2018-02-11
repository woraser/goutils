/*
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     helper
 * @date        2018-01-25 19:19
 */
package helper

import (
	"crypto/rand"
	r "math/rand"
	"time"
	"crypto/md5"
	"fmt"
)

var alphaNum = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

/**
生成随机字符串的字节切片。n为串的长度，alphabets为随机串的指定范围
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

//截取字符串
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

//md5转换
func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
