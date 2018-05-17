/*
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     file
 * @date        2018-01-25 19:19
 */
package file

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/woraser/goutils/helper"
	"os"
)

type fileIterator struct {
	file string        //文件的路径
	fp   *os.File      //文件的打开句柄
	rd   *bufio.Reader //按行读文件
}

type IteratorLineFunc func(string)

func NewFileIterator() *fileIterator {
	return &fileIterator{}
}

//设置文件
func (fi *fileIterator) SetFile(f string) *fileIterator {
	fi.file = f
	return fi
}

//初始化操作
func (fi *fileIterator) Init() (*fileIterator, error) {
	if fi.fp != nil {
		fi.fp.Close()
	}
	if !helper.FileExists(fi.file) {
		return nil, fmt.Errorf("file not exist,%s", fi.file)
	}
	fp, err := os.Open(fi.file)
	if err != nil {
		return nil, err
	}
	fi.rd = bufio.NewReader(fp)
	fi.fp = fp
	return fi, nil
}

//返回一个可遍历的通道
func (fi *fileIterator) IterLine(fn IteratorLineFunc) {
	if fi.fp == nil {
		fmt.Errorf("please execute Init() first\n")
		return
	}
	var buf bytes.Buffer
	for {
		line, isPrefix, err := fi.rd.ReadLine()
		if len(line) > 0 {
			buf.Write(line)
			if !isPrefix {
				fn(buf.String())
				buf.Reset()
			}
		}
		if err != nil {
			fmt.Printf("read file occur error\t%v\n", err)
			break
		}
	}
	fi.fp.Close()
}
