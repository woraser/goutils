package mysql

import (
	"testing"
	"fmt"
	"runtime"
)

func TestDBase_Conn(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	myconf := MySQLConf{
		Host:    "127.0.0.1",
		User:    "phpmyadmin",
		Passwd:  "123456",
		DbName:  "db_wendao",
		Charset: "utf8",
		Timeout: 5,
		Port:    3306,
	}
	db := NewDBase(myconf)
	db.SetDebug(true)
	_, err := db.Conn()
	defer db.Close()
	fmt.Println(err)
	sql := "select * from videoinfo limit 10"
	ret, err := db.FetchRows(sql)
	fmt.Println(ret, err)
}
