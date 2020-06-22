package utils

import (
	"bufio"
	"cron/db"
	"fmt"
	"os"
	"time"
)

var logFileDir = db.Cron.LogDir
var logFile = db.Cron.LogFile

//定义一个用户级的日志函数
func UserLog(str string) {
	filePath := logFileDir + time.Now().Format("2006-01") + "_" + logFile
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}//及时关闭 file 句柄
	defer file.Close()
	//准备写入 5 句 "你好,尚硅谷!"
	str = time.Now().Format("2006-01-02 15:04:05") + ":" + str + "\r\n"
	// \r\n 表示换行
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	//for i := 0; i < 10; i++ {
	_, _ = writer.WriteString(str)
	//}
	//因为 writer 是带缓存，因此在调用 WriterString 方法时，其实 //内容是先写入到缓存的,所以需要调用 Flush 方法，将缓冲的数据 //真正写入到文件中， 否则文件中会没有数据!!!
	_ = writer.Flush()
}