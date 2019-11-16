package main

import (
	"fmt"
	"time"
)

func timestampTurn(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func main() {
	now := time.Now()
	timestamp := now.Unix()
	fmt.Printf("now:%v\ntimestamp:%v\n", now, timestamp)
	var a int64 = 978282061
	fmt.Printf("time:%v\n", time.Unix(a, 0))
}
