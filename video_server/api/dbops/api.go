package dbops

import (
	// "fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_dev/video_server/api/defs"
	"go_dev/video_server/api/utils"
	"time"
)

// func init() {
// 	fmt.Println("api.go")
// }

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	defer stmtOut.Close()

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? and pwd =?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	defer stmtDel.Close()

	_, err = stmtDel.Exec(loginName,pwd)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(loginName string) (*defs.User, error) {
	stmtOut, err := dbConn.Prepare("SELECT id, pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var id int
	var pwd string

	err = stmtOut.QueryRow(loginName).Scan(&id, &pwd)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.User{Id: id, LoginName: loginName, Pwd: pwd}

	defer stmtOut.Close()

	return res, nil
}

func AddNewVideo(aid int, name string)(*defs.VideoInfo, error) {
	vid := utils.UniqueId()

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info
		(id,author_id, name, display_ctime) VALUES(?, ?, ?, ?)`)

	if err != nil {
		return nil ,err
	}

	defer stmtIns.Close()
	
	_, err = stmtIns.Exec(vid, aid, name, ctime)

	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtime:ctime}

	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmtOut.Close()
	var aid int
	var dct string
	var name string
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}
	
	res := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtime:dct}

	return res, nil
}

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info 
		INNER JOIN users ON video_info.author_id = users.id
		WHERE users.login_name = ? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time <= FROM_UNIXTIME(?) 
		ORDER BY video_info.create_time DESC`)


	var res []*defs.VideoInfo

	if err != nil {
		return res, err
	}
	
	rows, err := stmtOut.Query(uname, from, to)
	if err != nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}

		vi := &defs.VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCtime: ctime}
		res = append(res, vi)
	}

	defer stmtOut.Close()

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	defer stmtDel.Close()

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	return nil
}

func AddNewComments(vid string, aid int, content string) error{
	id := utils.UniqueId()

	stmtIns, err := dbConn.Prepare(`INSERT INTO comments
		(id,video_id, author_id, content) VALUES(?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	defer stmtIns.Close()
	
	_, err = stmtIns.Exec(id, vid, aid, content)

	if err != nil {
		return  err
	}

	return nil
}

func ListComments(vid string, from , to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments 
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) and comments.time <= FROM_UNIXTIME(?)`)

	if err != nil {
		return nil, err
	}

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from , to)
	if err != nil {

		// fmt.Printf("%v", err)

		return res, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id:id, VideoId:vid, Author:name, Content:content}
		res = append(res, c)
	}

	// fmt.Printf("%v", res)

	return res, nil
}