package opt

import (
	"cron/db"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

//查找多少天内的status=4(CP回调)状态的订单
const DAYS_WITH_GREOUP int64 = 365

const SUCCESS_DESC = "SUCCESS"
//const FAIL_DESC = "FAIL"

var wg sync.WaitGroup

func OrderCallBackWithGroup(){
	var Order_num string

	nowTime := time.Now().Unix()
	rows ,errQ := db.Db.Query("SELECT order_num FROM orders WHERE status = 4 and complete_time > ? order by complete_time desc limit 10000", nowTime - db.Cron.Days * 86400)

	if errQ != nil {
		fmt.Println("query rows error : ", errQ)
	}

	startTime := time.Now().Unix()

	i := 0
	for rows.Next() {
		err := rows.Scan(&Order_num)
		if err != nil{
			fmt.Println("scan err:",err)
			continue
		}

		wg.Add(1)
		//利用协程并发执行
		go handleCallBackWithGroup(Order_num)

		i++

	}

	wg.Wait()

	endTime := time.Now().Unix()

	fmt.Println("spend time :", endTime - startTime)

	fmt.Println("execute end!")

	select{}
}

func handleCallBackWithGroup(Order_num string) {

	defer wg.Done()

	time.Sleep(time.Second * 5)

	var Callback_url, Callback_data string
	//
	row := db.Db.QueryRow("SELECT callback_url, callback_data FROM error_callback WHERE order_num = ?", Order_num)
	err := row.Scan(&Callback_url, &Callback_data)
	if err != nil{
		fmt.Println("scan err:",err)
		//continue
	}

	res,_:= simplejson.NewJson([]byte(Callback_data))
	var orderData map[string]string
	orderData = make(map[string]string, 20)
	appId, _ := res.Get("app_id").Int()
	orderData["app_id"] = strconv.Itoa(appId)
	orderData["tf_trade_no"], _ = res.Get("tf_trade_no").String()
	orderData["chl_order_num"], _ = res.Get("chl_order_num").String()
	orderData["role_id"], _ = res.Get("role_id").String()
	orderData["server_id"], _ = res.Get("server_id").String()
	moneyType ,_ := res.Get("money_type").Int()
	orderData["money_type"]= strconv.Itoa(moneyType)
	orderData["total_fee"], _ = res.Get("total_fee").String()
	payType,_ := res.Get("pay_type").Int()
	orderData["pay_type"]= strconv.Itoa(payType)
	payResult,_ := res.Get("pay_result").Int()
	orderData["pay_result"] = strconv.Itoa(payResult)
	orderData["sign"], _ = res.Get("sign").String()

	var postData string = ""
	var flag string = ""
	for k, v := range orderData {
		if postData != "" {
			flag = "&"
		}
		postData += flag + k + "=" + v
	}

	fmt.Println(postData)
	//var resData io.Reader
	resp, _ := http.Post(Callback_url, "application/x-www-form-urlencoded" , strings.NewReader(postData))
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	reqRet := string(body)
	if reqRet == SUCCESS_DESC {
		//if reqRet == "FAIL" {
		//成功的话，修改订单的状态
		nowTime := time.Now().Unix()
		_, err = db.Db.Exec("update orders set status = 8, complete_time = ? where order_num=?", nowTime, Order_num)
		if err != nil {

		}
	}

	fmt.Println("wg:", Order_num)

}