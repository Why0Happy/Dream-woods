package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var per string
	stu0 := make([]string, 50)
	stu := make([]string, 50)
	go Read(per)
	Cut(per, stu0, "\"}")
	Cut("", stu, ",")
	num := make([]int, 50)
	go SelectNo(stu[:], num)
	SelectName(stu[:], num)
}

//读取信息
func Read(arr string) {
	f, err := os.Open("./Students.txt")
	if err != nil {
		fmt.Println("read Students.txt", err)
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fail", err)
	}
	arr = string(fd)
}

func Cut(str string, arr []string, a string) {
	for _, v := range strings.SplitAfter(str, a) {
		arr = append(arr, v)
	}
}

//检索学号
func SelectNo(arr []string, sum []int) {
	no := make([]int, 50)
	for i, _ := range arr {
		n := strings.Index(arr[i], "xh") //检索含学号的键值
		no = append(no, n)
	}
	for j, _ := range arr {
		n := strings.Index(arr[j], "573\"") //检索以573结尾的学号
		sum = append(sum, n)
	}
}

//检索姓名
func SelectName(arr []string, sum []int) {
	no := make([]int, 50)
	for i, _ := range arr {
		n := strings.Index(arr[i], "xmEn") //检索含姓名的键值
		no = append(no, n)
	}
	for n, _ := range sum {
		i := strings.Count(arr[n], " ") //是否含两个空格（字数）
		if i == 2 {
			fmt.Printf("%v", arr[n])
		}
	}
}
