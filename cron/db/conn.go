package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"encoding/json"
)

var Db *sql.DB

var Cron struct {
	User       string
	Passwd     string
	Net        string
	Addr       string
	DBName     string

	//日志的配置参数
	LogFile    string
	LogDir     string

	//执行频率
	Spec     string

	//获取多少天内的数据
	Days     int64
}

func init()  {

	data, err1 := ioutil.ReadFile("cron_cfg/cron.json")
	if err1 != nil {
		//log.Fatal("%v", err1)
	}
	err2 := json.Unmarshal(data, &Cron)
	if err2 != nil {
		//log.Fatal("%v", err2)
	}

	cfg := mysql.NewConfig()
	cfg.User = Cron.User
	cfg.Passwd = Cron.Passwd
	cfg.Net = Cron.Net
	cfg.Addr = Cron.Addr
	cfg.DBName = Cron.DBName
	dsn := cfg.FormatDSN()


	var err error
	Db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Println("connect mysql error:", err)
	}
}
