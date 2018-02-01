package file

import (
	"testing"
	"fmt"
	"strconv"
)

/**
CREATE TABLE `snowflake` (
  `gid` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`gid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='测试表';
 */
func TestNewFileIterator(t *testing.T) {
	return
	itor, err := NewFileIterator().SetFile("./snowflake.txt").Init()
	if err != nil {
		t.Errorf("%v", err)
	}
	//连接MySQL
	db, err := NewDBase(MakeMySQLConf().
		SetUser("liuyongshuai").
		SetDbName("db_wendao").
		SetHost("127.0.0.1").
		SetPort(3306).
		SetPasswd("123456")).Conn()
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	defer db.Close()

	//遍历每一行，写入库里
	itor.IterLine(func(line string) {
		id, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Errorf("%v", err)
			return
		}
		d := make(map[string]ItemElem)
		d["gid"] = MakeItemElem(id)
		_, ret, err := db.InsertData("snowflake", d, false)
		if !ret || err != nil {
			fmt.Errorf("%v", err)
			return
		}
	})
}
