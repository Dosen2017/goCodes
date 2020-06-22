package dbops

import (
	"testing"
	"fmt"
	"strconv"
	"time"
)

// init(dblogin, truncate tables)->run tests->clear data(truncate tables)

var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

// func TestUserWorkFlow(t *testing.T) {
// 	t.Run("Add", testAddUser)
// 	t.Run("Get", testGetUser)
// 	t.Run("Del", testDeleteUser)
// 	t.Run("Reget", testRegetUser)
// }

func testAddUser(t *testing.T) {
	err := AddUserCredential("sudesheng", "123")
	if err != nil {
		t.Errorf("Error of AddUser:%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("sudesheng")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("sudesheng", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("sudesheng")
	if err != nil {
		t.Errorf("Error of RegetUser:%v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user failed ")
	}
}

// func TestVideoWorkFlow(t *testing.T) {
// 	clearTables()
// 	t.Run("PrepareUser", testAddUser)
// 	t.Run("AddVideo", testAddVideoInfo)
// 	t.Run("GetVideo", testGetUserVideoInfo)
// 	t.Run("DelVideo", testDeleteUserVideoInfo)
// 	t.Run("RegetVideo", testRegetUserVideoInfo)
// }

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo:%v", err)
	}
	//fmt.Printf("%v", vi)
	tempvid = vi.Id
}

func testGetUserVideoInfo(t *testing.T) {

	// fmt.Println(tempvid)

	vi, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo")
	}

	if vi == nil {
		t.Errorf("Error of not GetVideoInfo")
	}
}

func testDeleteUserVideoInfo(t *testing.T) {
	err := DeleteVideoInfo("tempvid")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v", err)
	}
}

func testRegetUserVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)

	// fmt.Printf("%v", vi)

	if err != nil {
		t.Errorf("Error of GetVideoInfo")
	}

	if vi == nil {
		t.Errorf("Error of not GetVideoInfo")
	}
}

// func TestVideoWorkFlow(t *testing.T) {
// 	clearTables()
// 	t.Run("AddUser", testAddUser)
// 	t.Run("AddComments", testAddComments)
// 	// t.Run("AddComments2", testAddComments2)
// 	t.Run("ListComments", testListComments)
// }

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"
	err := AddNewComments(vid , aid , content )
	if err != nil {
		t.Errorf("Error of AddComments:%v", err)
	}
}

// func testAddComments2(t *testing.T) {
// 	vid := "12345"
// 	aid := 1
// 	content := "I like this video 2"
// 	err := AddNewComments(vid , aid , content )
// 	if err != nil {
// 		t.Errorf("Error of AddComments:%v", err)
// 	}
// }


func testListComments(t *testing.T) {

	vid := "12345"
	from := 1569484074
	to , _:= strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res, err := ListComments(vid , from , to )
	if err != nil {

		fmt.Printf("%v", err)

		t.Errorf("Error of ListComments")
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}


func TestSessionsWorkFlow(t *testing.T) {
	clearTables()
	t.Run("InsertSession", testInsertSession)
	t.Run("InsertSession2", testInsertSession2)
	t.Run("RetrieveSession", testRetrieveSession)
	t.Run("RetrieveAllSession", testRetrieveAllSession)
}

func testInsertSession(t *testing.T)  {
	sid := "session12345"
	ttl := int64(1515426665)
	uname := "I like this video"
	err := InsertSession(sid , ttl , uname )

	if  err != nil {
		t.Errorf("Error of InsertSession")
	}
}

func testInsertSession2(t *testing.T)  {
	sid := "session1234578"
	ttl := int64(1515426665)
	uname := "I like this video78"
	err := InsertSession(sid , ttl , uname )

	if  err != nil {
		t.Errorf("Error of InsertSession")
	}
}

func testRetrieveSession(t *testing.T)  {
	sid := "session12345"
	_, err := RetrieveSession(sid)
	if err != nil {
		t.Errorf("Error of RetrieveSession")
	}

	// fmt.Printf("retrieveSession %v\n", simpleSession)
	
}

func testRetrieveAllSession(t *testing.T)  {
	sessions , err := RetrieveAllSession()
	if err != nil {
		t.Errorf("testRetrieveAllSession err : %v", err)
	}
	// fmt.Printf("testRetrieveAllSession %v\n", sessions)
	//测试遍历session的方法，只能打印出来
	sessions.Range(func(k, v interface{}) bool {
        fmt.Println("iterate:", k, v)
        return true
    })
}