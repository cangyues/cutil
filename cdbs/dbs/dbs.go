package dbs

import (
	json "cjson/qson"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	WriteDB *sql.DB
	ReadDB  *sql.DB
}

func getDB(dbName string, url string) *sql.DB {
	db, err := sql.Open(dbName, url)
	if err != nil {
		fmt.Println("数据库连接异常...", err)
	}
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	return db
}

func (db *DB) Init(dbType string, url string) {
	db.ReadDB = getDB(dbType, url)
	db.WriteDB = db.ReadDB
}

func (db *DB) ReadDb(dbName string, url string) {
	db.ReadDB = getDB(dbName, url)
}

func (db *DB) WriteDb(dbName string, url string) {
	db.WriteDB = getDB(dbName, url)
}

func (db *DB) _Query(_sql string, array json.JSONArray) *json.JSONArray {
	return db.Query(_sql, array.ToArray()...)
}

func (db *DB) Query(_sql string, param ...interface{}) *json.JSONArray {
	var row *sql.Rows
	var err error
	if param == nil {
		row, err = db.ReadDB.Query(_sql)
	} else {
		row, err = db.ReadDB.Query(_sql, param)
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
		r, e = db.WriteDB.Exec(_sql)
	} else {
		r, e = db.WriteDB.Exec(_sql, param)
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
	tx, err := db.WriteDB.Begin()
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
