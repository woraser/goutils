/**
 * @author      Liu Yongshuai
 * @package     file
 * @date        2018-04-04 12:23
 */
package file

import (
	"fmt"
	"os"
	"strings"
)

type FileTool struct {
	flist []string //读取目录下的文件列表结果
}

//初始化操作
func (ft *FileTool) Init() {
	ft.flist = make([]string, 0)
}

//判断文件或目录是否存在
func (ft *FileTool) IsExists(f string) bool {
	if _, err := os.Stat(f); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//读取目录下的所有文件
func (ft *FileTool) ReadDirFiles(dir string) []string {
	if !ft.IsExists(dir) {
		errmsg := fmt.Sprintf("dir %v not exists", dir)
		fmt.Println(errmsg)
		return ft.flist
	}
	dir = strings.TrimRight(dir, "/")
	//开始读取目录
	f, err := os.Open(dir)
	if err != nil {
		errmsg := fmt.Sprintf("open dir error %s %v", dir, err)
		fmt.Println(errmsg)
		return ft.flist
	}
	dlist, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		errmsg := fmt.Sprintf("read dir error %s %v", dir, err)
		fmt.Println(errmsg)
		return ft.flist
	}
	//开始递归读取目录，忽略掉隐藏目录及文件
	for _, v := range dlist {
		if strings.HasPrefix(v.Name(), ".") {
			continue
		}
		f := dir + "/" + v.Name()
		if v.IsDir() {
			ft.ReadDirFiles(f)
		} else {
			ft.flist = append(ft.flist, f)
		}
	}
	return ft.flist
}
