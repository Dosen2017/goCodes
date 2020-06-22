package dbops

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_,err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord err: %v", err)
		return err
	}
	return nil
}