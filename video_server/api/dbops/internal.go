package dbops

import (
	"log"
	"strconv"
	"sync"
	"go_dev/video_server/api/defs"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL,login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmtIns.Close()

	// fmt.Println("InsertSession---" ,sid, "----", ttlstr, "----", uname)

	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error){

	//fmt.Println(".......RetrieveSession")

	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name from sessions where session_id = ?")
	if err != nil {
		// fmt.Println("RetrieveSession err:", err)
		return nil, err
	}


	// fmt.Println("RetrieveSession")
	defer stmtOut.Close()

	var ttl string
	var uname string

	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)

	// fmt.Println("RetrieveSession---", ttl, "----", uname)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err	
	}

	// fmt.Println("222RetrieveSession---", ss.TTL, "----", ss.Username)

	return ss, nil
}

func RetrieveAllSession() (*sync.Map, error){

	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("select * from sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
// fmt.Printf("%v\n", rows.Next())
// fmt.Printf("%v\n", rows.Next())
// fmt.Printf("%v\n", rows.Next())
	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("retrieve sessions error:%s", err)
			return nil, err
		}
		// fmt.Println("RetrieveAllSession ", id, "--" , ttlstr, "---", login_name)
		
		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64);err1 == nil {
			ss := &defs.SimpleSession{Username:login_name, TTL:ttl}
			m.Store(id, ss)
			log.Printf("session id:%s, ttl:%d", id, ss.TTL)
		}
	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions where session_id = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _,err := stmtOut.Query(sid); err != nil {
		return err
	}
	return nil
}