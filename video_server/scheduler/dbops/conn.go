package dbops

import (
	// "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:@/video_server?charset=utf8")
	if err != nil {
		//fmt.Println("conn error")
		panic(err.Error())
	}

	// fmt.Println("conn.go")
	// defer dbConn.Close()
}