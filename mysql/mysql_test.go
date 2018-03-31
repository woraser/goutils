package mysql

import (
	"fmt"
	"runtime"
	"testing"
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

func TestDBase_InsertBatchData(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	myconf := MySQLConf{
		Host:       "127.0.0.1",
		User:       "phpmyadmin",
		Passwd:     "123456",
		DbName:     "db_wendao",
		Charset:    "utf8",
		Timeout:    5,
		Port:       3306,
		AutoCommit: true,
	}
	db := NewDBase(myconf)
	db.SetDebug(true)
	_, err := db.Conn()
	defer db.Close()
	fmt.Println(err)
	var data [][]interface{}
	fields := []string{"id1", "id2"}
	for i := 0; i < 100; i++ {
		var tmp []interface{}
		tmp = append(tmp, i)
		tmp = append(tmp, i+1)
		data = append(data, tmp)
	}
	ret, b, e := db.InsertBatchData("test", fields, data, true)
	fmt.Println(ret, b, e)
}
