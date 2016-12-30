/*
***********************************************************
// Describe : golang 操作数据库
// date  : 2016.12.30
// Author :
*********************************************************
*/
package mysql

import (
	"database/sql"
	//"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenMysql() {

	//连接数据库
	//DSN 格式为：username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", "root:mumusan@tcp(localhost:3306)/test")
	if err != nil {
		panic("open database fail")
	}
	err = db.Ping()

	if err != nil {
		panic("linked database fail")
	}
}
