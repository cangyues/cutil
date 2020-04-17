package dbs

import (
	json "cjson/qson"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type DB struct {
	WriteDB     []*sql.DB
	ReadDB      []*sql.DB
	WriteIndex  int
	ReadIndex   int
	WriteLength int
	ReadLength  int
	Mutex       sync.Mutex
}

func (db *DB) getWrite() *sql.DB {
	wdb := db.WriteDB[db.WriteIndex]
	if db.ReadLength == 1 { //单例模式直接返回
		return wdb
	}
	db.Mutex.Lock()         //上锁
	defer db.Mutex.Unlock() //解锁

	if db.WriteIndex+1 > db.WriteLength {
		db.WriteIndex = 0
	} else {
		db.WriteIndex++
	}
	return wdb
}

func (db *DB) getRead() *sql.DB {
	rdb := db.ReadDB[db.ReadIndex]
	if db.ReadLength == 1 { //单例模式直接返回
		return rdb
	}
	db.Mutex.Lock()         //上锁
	defer db.Mutex.Unlock() //解锁

	if db.ReadIndex+1 > db.ReadLength {
		db.ReadIndex = 0
	} else {
		db.ReadIndex++
	}
	return rdb
}

func getDB(dbName string, url string) *sql.DB {
	db, err := sql.Open(dbName, url)
	if err != nil {
		fmt.Println("数据库连接异常...", err)
	}
	db.SetMaxIdleConns(40)
	db.SetMaxOpenConns(50)
	return db
}

//单例模式使用该方法初始化
func (db *DB) Init(dbType string, url string) {
	db.WriteLength = 1
	db.ReadLength = 1
	da := make([]*sql.DB, 1)
	da[0] = getDB(dbType, url)
	db.ReadDB = da
	db.WriteDB = db.ReadDB
}

func (db *DB) ReadDb(dbName string, url []string) {
	db.ReadLength = len(url)
	da := make([]*sql.DB, db.ReadLength)
	for i, v := range url {
		da[i] = getDB(dbName, v)
	}
	db.ReadDB = da
}

func (db *DB) WriteDb(dbName string, url []string) {
	db.WriteLength = len(url)
	da := make([]*sql.DB, db.WriteLength)
	for i, v := range url {
		da[i] = getDB(dbName, v)
	}
	db.WriteDB = da
}

func (db *DB) _Query(_sql string, array json.JSONArray) *json.JSONArray {
	return db.Query(_sql, array.ToArray()...)
}

func (db *DB) Query(_sql string, param ...interface{}) *json.JSONArray {
	var row *sql.Rows
	var err error
	if param == nil {
		row, err = db.getRead().Query(_sql)
	} else {
		row, err = db.getRead().Query(_sql, param)
	}
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()
	cols, _ := row.Columns()
	len := len(cols)
	val := make([]interface{}, len)
	scans := make([]interface{}, len)
	for k, _ := range val {
		scans[k] = &val[k]
	}
	result := json.NewJSONArray()
	for row.Next() {
		row.Scan(scans...)
		_r := json.NewJSONObject()
		for k, v := range val {
			key := cols[k]
			b, _ := v.([]byte)
			_r.PutString(key, string(b))
		}
		result.Add(_r)
	}
	return result
}

func (db *DB) Exec(_sql string, param ...interface{}) (sql.Result, error) {
	var (
		r sql.Result
		e error
	)
	if param == nil {
		r, e = db.getWrite().Exec(_sql)
	} else {
		r, e = db.getWrite().Exec(_sql, param)
	}
	return r, e
}

func (db *DB) _Exec(sql string, array json.JSONArray) (sql.Result, error) {
	return db.Exec(sql, array.ToArray()...)
}

func (db *DB) Insert(_sql string, array json.JSONArray) (int64, error) {
	var (
		r sql.Result
	)
	if array.IsEmpty() {
		r, _ = db.Exec(_sql, nil)
	} else {
		r, _ = db.Exec(_sql, array.ToArray()...)
	}
	return r.LastInsertId()
}

func (db *DB) BatchExec(array []string) bool {
	tx, err := db.getWrite().Begin()
	if err != nil {
		fmt.Println("执行批处理异常！", err)
		return false
	}
	for _, v := range array {
		_, e := tx.Exec(v)
		if e != nil {
			tx.Rollback()
			return false
		}
	}
	tx.Commit()
	return true
}
