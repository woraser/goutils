package mysql

import (
	"testing"
	"github.com/kr/pretty"
	"fmt"
	"runtime"
	"github.com/liuyongshuai/goElemItem"
)

func TestFormatCond(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	delim := "and"
	cond := make(map[string]goElemItem.ItemElem)
	cond["id:in"] = goElemItem.MakeItemElem("444,666,888")
	cond["uid:in"] = goElemItem.MakeItemElem(444)
	cond["name:like"] = goElemItem.MakeItemElem("liuyongshuai")
	cond["title:rlike"] = goElemItem.MakeItemElem([]string{"sina", "baidu"}) //非法
	cond["tid:lt"] = goElemItem.MakeItemElem(400)
	cond["tags:find"] = goElemItem.MakeItemElem("google")
	cond["category:find"] = goElemItem.MakeItemElem([]interface{}{444, "didi"}) //非法
	cond["cnum"] = goElemItem.MakeItemElem("99999")
	cond["praiseNum"] = goElemItem.MakeItemElem([]interface{}{"aaaa", 999}) //非法
	sqlCond, param := FormatCond(cond, delim)
	/**
	sqlCond "`id` IN (?,?,?) AND `tid` < ? AND FIND_IN_SET(?,`tags`) AND `cnum` = ? AND `uid` IN (?) AND `name` LIKE ?"
	param {}{"444","666","888",int(400),"google","99999",int(444),"%liuyongshuai%",}
	 */
	fmt.Printf("sqlCond %# v\n", pretty.Formatter(sqlCond))
	fmt.Printf("param %# v\n", pretty.Formatter(convertArgs(param)))
}
