package session

import (
	"time"
	"sync"
	"go_dev/video_server/api/defs"
	"go_dev/video_server/api/dbops"
	"go_dev/video_server/api/utils"
	"fmt"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowIntMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSession()
	if err != nil {
		return 
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k ,ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id := utils.UniqueId()
	ct := nowIntMilli()
	ttl := ct + 30 * 60 * 1000
	ss := &defs.SimpleSession{Username:un, TTL:ttl}
	sessionMap.Store(id, ss)

	//io.WriteString(w, ubody.Username + ubody.Pwd)

fmt.Println(id, "----", ttl, "----", un)

	err := dbops.InsertSession(id, ttl, un)

	if err != nil {
		fmt.Println("insertSession err :", err)
	}

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowIntMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			//delete expired session
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}
