package mysql

import (
	"testing"
	"github.com/kr/pretty"
	"fmt"
	"runtime"
	"github.com/liuyongshuai/goutils/elem"
)

func TestFormatCond(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	delim := "and"
	cond := make(map[string]elem.ItemElem)
	cond["id:in"] = elem.MakeItemElem("444,666,888")
	cond["uid:in"] = elem.MakeItemElem(444)
	cond["name:like"] = elem.MakeItemElem("liuyongshuai")
	cond["title:rlike"] = elem.MakeItemElem([]string{"sina", "baidu"}) //非法
	cond["tid:lt"] = elem.MakeItemElem(400)
	cond["tags:find"] = elem.MakeItemElem("google")
	cond["category:find"] = elem.MakeItemElem([]interface{}{444, "didi"}) //非法
	cond["cnum"] = elem.MakeItemElem("99999")
	cond["praiseNum"] = elem.MakeItemElem([]interface{}{"aaaa", 999}) //非法
	sqlCond, param := FormatCond(cond, delim)
	/**
	sqlCond "`id` IN (?,?,?) AND `tid` < ? AND FIND_IN_SET(?,`tags`) AND `cnum` = ? AND `uid` IN (?) AND `name` LIKE ?"
	param {}{"444","666","888",int(400),"google","99999",int(444),"%liuyongshuai%",}
	 */
	fmt.Printf("sqlCond %# v\n", pretty.Formatter(sqlCond))
	fmt.Printf("param %# v\n", pretty.Formatter(convertArgs(param)))
}
