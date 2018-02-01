# mysql
针对golang查询MySQL的二次封装
```
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
id := 10000
d := make(map[string]ItemElem)
d["gid"] = MakeItemElem(id)
_, ret, err := db.InsertData("table_name", d, false)
if !ret || err != nil {
    fmt.Errorf("%v", err)
    return
}
```
