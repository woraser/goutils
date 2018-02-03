/*
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     mysql
 * @date        2018-01-25 19:19
 */
package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	gItem "github.com/liuyongshuai/goutils/elem"
	"strconv"
)

//MySQL连接的类
type DBase struct {
	Db      *sql.DB
	IsDebug bool
	Conf    MySQLConf
}

func NewDBase(conf MySQLConf) *DBase {
	return &DBase{
		Db:      nil,
		IsDebug: false,
		Conf:    conf,
	}
}

//提取查询SQL的结果
func reFormatRowsData(rows *sql.Rows) (ret []map[string]gItem.ItemElem, err error) {
	if rows == nil {
		return
	}
	defer rows.Close()
	//所有的列名
	cols, err := rows.Columns()
	if err != nil {
		return ret, err
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return ret, err
	}
	//是否有错误
	if err = rows.Err(); err != nil {
		return ret, err
	}
	//存储每一行值的slice，每次传地址给scan方法
	vals := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(vals))
	for i := range vals {
		scanArgs[i] = &vals[i]
	}
	//遍历每一行，提取数据
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return ret, fmt.Errorf("get data failed")
		}
		tmp := make(map[string]gItem.ItemElem)
		//根据不同的类型来将[]byte转换一下
		for i, colByteVal := range vals {
			colType := colTypes[i]
			realVal := convertMySQLType(colByteVal, colType)
			tmp[cols[i]] = gItem.MakeItemElem(realVal)
		}
		ret = append(ret, tmp)
	}
	return ret, nil
}

//连接MySQL
func (my *DBase) Conn() (*DBase, error) {
	conf := my.Conf
	if len(conf.Host) <= 0 {
		return nil, fmt.Errorf("invalid mysql host")
	}
	host := conf.Host
	if len(conf.DbName) <= 0 {
		return nil, fmt.Errorf("invalid mysql database")
	}
	dname := conf.DbName
	if len(conf.Charset) <= 0 {
		return nil, fmt.Errorf("invalid mysql charset")
	}
	autoCommit := "1"
	if !conf.AutoCommit {
		autoCommit = "0"
	}
	charset := conf.Charset
	if conf.Timeout <= 0 {
		return nil, fmt.Errorf("invalid mysql timeout")
	}
	tm := strconv.FormatInt(int64(conf.Timeout), 10)
	port := strconv.FormatUint(uint64(conf.Port), 10)
	//开始连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&autocommit=%s&timeout=%ss",
		conf.User, conf.Passwd, host, port, dname, charset, autoCommit, tm)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		my.Db = nil
		return nil, fmt.Errorf("connect MySQL failed:%s", err)
	}
	//设置连接池参数
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetConnMaxLifetime(conf.ConnMaxLiftTime)
	my.Db = db
	return my, nil
}

//目前的打开连接数
func (my *DBase) OpenConnNum() int {
	return my.Db.Stats().OpenConnections
}

//提取单行的单个字段
func (my *DBase) FetchOne(sql string, args ...interface{}) (ret gItem.ItemElem, err error) {
	r, err := my.FetchCols(sql, args...)
	if err != nil {
		return
	}
	if len(r) <= 0 {
		return
	}
	return r[0], nil
}

//提取所有行的第一个字段的列表
func (my *DBase) FetchCols(sql string, args ...interface{}) ([]gItem.ItemElem, error) {
	rets, err := my.FetchRows(sql, args...)
	if err != nil {
		return nil, err
	}
	if len(rets) <= 0 {
		return nil, nil
	}
	var ret []gItem.ItemElem
	for _, r := range rets { //r为一个map
		for _, iv := range r {
			ret = append(ret, iv)
		}
	}
	return ret, nil
}

//提取一行数据
func (my *DBase) FetchRow(sql string, args ...interface{}) (map[string]gItem.ItemElem, error) {
	ret, err := my.FetchRows(sql, args...)
	if err != nil {
		return nil, err
	}
	if len(ret) <= 0 {
		return nil, nil
	}
	return ret[0], nil
}

//提取多行数据
func (my *DBase) FetchRows(sql string, args ...interface{}) ([]map[string]gItem.ItemElem, error) {
	if my.Db == nil {
		return nil, fmt.Errorf("not connect MySQL")
	}
	if my.IsDebug {
		fmt.Printf("doFetch:\n")
		fmt.Printf("\tSQL:%s\n", sql)
	}
	stmt, err := my.Db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return reFormatRowsData(rows)
}

//提取多行数据
//table为表名
//cond为查询条件，全为and
//fields为要查询的字段，为空时表示查询全部
func (my *DBase) FetchCondRows(table string, cond map[string]gItem.ItemElem, fields ...string) (ret []map[string]gItem.ItemElem, err error) {
	f := filterTableFields(fields...)
	cd, param := FormatCond(cond, "AND")
	fsql := fmt.Sprintf("SELECT %s FROM `%s`", f, table)
	execArgs := ConvertArgs(param)
	if len(cd) > 0 {
		fsql = fmt.Sprintf("%s WHERE %s", fsql, cd)
	}
	return my.FetchRows(fsql, execArgs...)
}

//执行一条insert/update/delete语句，返回影响行数
func (my *DBase) Execute(sql string, args ...interface{}) (int64, bool, error) {
	ret, err := my.doExec(sql, args...)
	if err != nil {
		return 0, false, err
	}
	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		return 0, false, err
	}
	return rowsAffected, true, nil
}

//删除一条数据，返回lastAffectedRows
func (my *DBase) DeleteData(table string, cond map[string]gItem.ItemElem) (int64, bool, error) {
	cd, param := FormatCond(cond, "AND")
	dsql := fmt.Sprintf("DELETE FROM `%s`", table)
	execArg := ConvertArgs(param)
	if len(cd) > 0 {
		dsql = fmt.Sprintf("%s WHERE %s", dsql, cd)
	}
	return my.Execute(dsql, execArg...)
}

//写入一条数据，返回lastInsertId
func (my *DBase) InsertData(table string, data map[string]gItem.ItemElem, isIgnore bool) (int64, bool, error) {
	ignore := "IGNORE"
	if !isIgnore {
		ignore = ""
	}
	cd, param := FormatCond(data, ",")
	if len(cd) == 0 {
		return 0, false, fmt.Errorf("invalid insert data")
	}
	execArgs := ConvertArgs(param)
	isql := fmt.Sprintf("INSERT %s INTO `%s` SET %s", ignore, table, cd)
	ret, err := my.doExec(isql, execArgs...)
	if err != nil {
		return 0, false, err
	}
	lastInsertId, err := ret.LastInsertId()
	if err != nil {
		return 0, false, err
	}
	return lastInsertId, true, nil
}

//执行一条：INSERT INTO table (a,b,c) VALUES (1,2,3) ON DUPLICATE KEY UPDATE c=c+1 语句
func (my *DBase) InsertUpdateData(table string, insert map[string]gItem.ItemElem, update map[string]gItem.ItemElem) (int64, bool, error) {
	icd, iparam := FormatCond(insert, ",")
	ucd, uparam := FormatCond(update, ",")
	if len(icd) == 0 || len(ucd) == 0 {
		return 0, false, fmt.Errorf("invalid insert/update data")
	}
	iparam = append(iparam, uparam...)
	execArgs := ConvertArgs(iparam)
	iusql := fmt.Sprintf("INSERT INTO `%s` SET %s ON DUPLICATE KEY UPDATE %s", table, icd, ucd)
	return my.Execute(iusql, execArgs...)
}

//更新一条数据，返回lastAffectedRows
func (my *DBase) UpdateData(table string, data map[string]gItem.ItemElem, cond map[string]gItem.ItemElem) (int64, bool, error) {
	dcd, dparam := FormatCond(data, ",")
	ccd, cparam := FormatCond(cond, "AND")
	if len(dcd) == 0 {
		return 0, false, fmt.Errorf("invalid update data")
	}
	usql := fmt.Sprintf("UPDATE `%s` SET %s", table, dcd)
	if len(ccd) > 0 {
		usql = fmt.Sprintf("%s WHERE %s", usql, ccd)
		dparam = append(dparam, cparam...)
	}
	execArgs := ConvertArgs(dparam)
	return my.Execute(usql, execArgs...)
}

//执行一条select ... for update语句
func (my *DBase) FetchForUpdate(table string, cond map[string]gItem.ItemElem) (map[string]gItem.ItemElem, error) {
	cd, param := FormatCond(cond, "AND")
	execArgs := ConvertArgs(param)
	fusql := fmt.Sprintf("SELECT * FROM `%s`", table)
	if len(cd) > 0 {
		fusql = fmt.Sprintf("%s WHERE %s", fusql, cd)
	}
	fusql = fmt.Sprintf("%s FOR UPDATE", fusql)
	return my.FetchRow(fusql, execArgs...)
}

//执行一条写语句
func (my *DBase) doExec(sql string, args ...interface{}) (sql.Result, error) {
	if my.Db == nil {
		return nil, fmt.Errorf("not connect MySQL")
	}
	if my.IsDebug {
		fmt.Printf("\ndoExec:\n")
		fmt.Printf("\tSQL:%s\n", sql)
	}
	stmt, err := my.Db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	ret, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//关闭连接
func (my *DBase) Close() error {
	if my.IsDebug {
		fmt.Printf("\nstart close mysql....\n")
	}
	if my.Db == nil {
		return fmt.Errorf("not connect mysql")
	}
	return my.Db.Close()
}

//提取原始的DB对象
func (my *DBase) GetDB() *sql.DB {
	return my.Db
}

//提取原始的DB对象
func (my *DBase) SetDebug(d bool) {
	my.IsDebug = d
}

//Ping
func (my *DBase) Ping() error {
	if my.Db == nil {
		return fmt.Errorf("Not Connect MySQL....")
	}
	return my.Db.Ping()
}

//开启事务
func (my *DBase) BeginTransaction() (*sql.Tx, error) {
	if my.IsDebug {
		fmt.Printf("\nstart BeginTransaction....\n")
	}
	return my.Db.Begin()
}

//提交事务
func (my *DBase) CommitTransaction(tx *sql.Tx) error {
	if my.IsDebug {
		fmt.Printf("\nstart CommitTransaction....\n")
	}
	return tx.Commit()
}

//回滚事务
func (my *DBase) RollBackTransaction(tx *sql.Tx) error {
	if my.IsDebug {
		fmt.Printf("\nstart RollBackTransaction....\n")
	}
	return tx.Rollback()
}
