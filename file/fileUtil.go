package file

import (
	"os"
	"io"
	"mime/multipart"
	"path/filepath"
	"bytes"
	"net/http"
	"fmt"
	"bufio"
	"io/ioutil"
	"errors"
	"go/ast"
)


/*读取文件*/
func ReadFileForByte(filePath string) ([]byte,error) {
	if CheckFileExist(filePath) == false {
		return nil,errors.New("file is not exits")
	}
	data, err := ioutil.ReadFile(filePath)
	return data,err
}

/*逐行 读取文件*/
func ReadFileBufferOnLine(filePath string) ([]string,error) {
	var content []string
	if CheckFileExist(filePath) == false {
		return content,errors.New("file is not exits")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return content,err
	}
	defer file.Close()
	bufferedReader := bufio.NewReader(file)
	count := 0
	for {
		line, readerError := bufferedReader.ReadString('\n')
		if (readerError != nil) && (readerError == io.EOF) {
			break
		}
		count++
		content = append(content, line)
	}
	return content,nil
}

/*按照字节写入文件*/
func WriteFileByByte(filePath string, content []byte) error {
	err := ioutil.WriteFile(filePath, content, 0666)
	return err
}

/*按照字符串写入文件*/
func WriteFileBufferByString(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	bufferedWrite := bufio.NewWriter(file)
	_,errs := bufferedWrite.WriteString(content)
	if errs != nil {
		return errs
	}
	bufferedWrite.Flush()
	return nil
}

// fileName:文件名字(带全路径)
// content: 写入的内容
func AppendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	return err
}

//文件拷贝
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

//获取文件大小
func GetFileSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	fileSize := fileInfo.Size() //获取size
	return fileSize
}

/*判断文件是否存在*/
func CheckFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}


/*上传文件*/
/*
uri:路径
params:额外参数
paramName:文件参数名称
path:文件完整路径
*/
func UploadFile(uri string, params map[string]string, paramName, path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return false
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return false
	}
	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	} else {
		if resp.StatusCode != 200 {
			return false
		}
	}
	return true
}