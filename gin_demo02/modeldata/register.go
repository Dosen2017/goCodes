package modeldata

import (
	//"fmt"
	"gin_demo02/db"
	"time"
	"encoding/json"
)

type StreamRegister struct {

}

func (this *StreamRegister)Add (dataMap map[string]string) bool {

	flag := true

	dataSlice, _ := json.Marshal(dataMap)
	data := string(dataSlice)

	nowTime := time.Now().Unix()


	_, err := db.Db.Exec("insert into stream_register(`app_id`, `account_id`, `server_id`, `role_id`,`create_time`,`data`, `time`) values (?, ?, ?, ?, ?, ?, ?)" , dataMap["app_id"], dataMap["account_id"], dataMap["server_id"], dataMap["role_id"], dataMap["create_time"], data, nowTime)
	if  err != nil {
		flag = false
	}

	return flag
}