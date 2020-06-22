package opt

import (
	"fmt"
	"gin_demo02/common"
	"gin_demo02/modeldata"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Req struct {

}

func (this *Req)AddRegisterData(c *gin.Context) {

	//_ = ioutil.WriteFile("./8080.txt", []byte("8080 data write"), 0655)

	dataMap := make(map[string]string)

	dataMap["app_id"] = c.Query("app_id")
	dataMap["account_id"] = c.Query("account_id")
	dataMap["server_id"] = c.Query("server_id")
	dataMap["role_id"] = c.Query("role_id")
	dataMap["create_time"] = c.Query("create_time")
	dataMap["sign"] = c.Query("sign")

	fmt.Println(dataMap)

	flag := true

	//判断是否为空
	flag = common.CheckParamIsEmpty(dataMap)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"code" : 1,
			"message": "param is null",
		})
		return
	}

	//判断sign是否正确
	flag = common.CheckSignIsRight(dataMap)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"code" : 2,
			"message": "sign is error",
		})
		return
	}

	streamRegister := &modeldata.StreamRegister{}
	flag = streamRegister.Add(dataMap)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"code" : 3,
			"message": "add fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : 0,
		"message": "add success",
	})

}

