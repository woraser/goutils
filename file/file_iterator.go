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
	"github.com/liuyongshuai/goutils/helper"
	"os"
	"sync"
)

type fileIterator struct {
	ch         chan string   //缓存通道
	chSize     uint64        //通道的大小
	file       string        //文件的路径
	fp         *os.File      //文件的打开句柄
	rd         *bufio.Reader //按行读文件
	lock       *sync.Mutex   //同步用的
	isHaveInit bool          //是否已经初始化了
}

type IteratorLineFunc func(string)

func NewFileIterator() *fileIterator {
	return &fileIterator{
		chSize:     1000,
		lock:       new(sync.Mutex),
		isHaveInit: false,
	}
}

//设置通道大小
func (fi *fileIterator) SetChSize(n uint64) *fileIterator {
	fi.isHaveInit = false
	fi.chSize = n
	return fi
}

//设置文件
func (fi *fileIterator) SetFile(f string) *fileIterator {
	fi.isHaveInit = false
	fi.file = f
	return fi
}

//初始化操作
func (fi *fileIterator) Init() (*fileIterator, error) {
	fi.lock.Lock()
	defer fi.lock.Unlock()
	if fi.isHaveInit {
		return fi, nil
	}
	if fi.ch != nil {
		close(fi.ch)
	}
	if fi.fp != nil {
		fi.fp.Close()
	}
	fi.ch = make(chan string, fi.chSize)
	if !helper.FileExists(fi.file) {
		return nil, fmt.Errorf("file not exist,%s", fi.file)
	}
	fp, err := os.Open(fi.file)
	if err != nil {
		return nil, err
	}
	fi.isHaveInit = true
	fi.rd = bufio.NewReader(fp)
	fi.fp = fp
	go fi.readline()
	return fi, nil
}

//返回一个可遍历的通道
func (fi *fileIterator) IterLine(fn IteratorLineFunc) {
	if !fi.isHaveInit {
		fmt.Errorf("please execute Init() first\n")
		return
	}
	for line := range fi.ch {
		fn(line)
	}
}

//读每行数据
func (fi *fileIterator) readline() {
	if !fi.isHaveInit {
		fmt.Errorf("please execute Init() first\n")
		return
	}
	var buf bytes.Buffer
	for {
		line, isPrefix, err := fi.rd.ReadLine()
		if len(line) > 0 {
			buf.Write(line)
			if !isPrefix {
				fi.ch <- buf.String()
				buf.Reset()
			}
		}
		if err != nil {
			fmt.Printf("read file occur error\t%v\n", err)
			break
		}
	}
	close(fi.ch)
	fi.fp.Close()
}
