package main

import "fmt"

func main() {
	foodList := map[string]int{"小面": 6,
		"饭团":   7,
		"香菇滑鸡": 12,
		"小炒肉":  15,
		"黄焖鸡":  16,
		"冒菜":   18}
	for i, v := range foodList {
		fmt.Printf("菜式：%v 价格：%v\n", i, v)
	}
}
