package common

import (
	"crypto/md5"
	"fmt"
	"sort"
)

func CheckParamIsEmpty(dataMap map[string]string) bool {
	//判断是否为空
	flag := true
	for _, v := range dataMap {
		if v == "" {
			flag = false
		}
	}
	return flag
}

func CheckSignIsRight(dataMap map[string]string) bool {
	flag := true

	kSlice := make([]string, 0, 20)
	for k, _ := range dataMap {
		kSlice = append(kSlice, k)
	}
	sort.Strings(kSlice)
	signStr := ""
	for _, v := range kSlice {
		if v == "sign" {
			continue
		}
		signStr += v + "=" + dataMap[v] + "&"
	}

	appKey := "abc123"  //理论上是从数据库里面去取
	md5Sign := fmt.Sprintf("%x", md5.Sum([]byte(signStr + appKey)))

	fmt.Println(md5Sign)

	if dataMap["sign"] != md5Sign {
		flag = false
	}

	return flag
}