package common

import (
	"path/filepath"
	"os"
	"strings"
	"crypto/rand"
	"encoding/hex"
	"io"
	"encoding/base64"
	"crypto/md5"
)

/*获取当前运行程序的当前目录*/
func GetCurrentDirectory() (string,error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "",err
	}
	return strings.Replace(dir, "\\", "/", -1),nil
}

/*生成唯一ID*/
func GetUUId() (string, error) {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b)), nil
}


/*生成32位md5字串*/
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
