package main

import (
	"fmt"
)

var (
	myres = make(map[int]int, 20)
)

func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	channel <- res
	myres[n] = <-channel

}

var channel = make(chan int, 20)

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}

	for i := 1; i <= 20; i++ {
		fmt.Printf("myres[%d] = %d\n", i, myres[i])
	}

}
