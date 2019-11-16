package main

import (
	"bufio"
	"os"
)
import (
	"strconv"
	"time"
)
import (
	"fmt"
)

func main() {
	counts1 := make(map[string]int)
	counts2 := make(map[int]int64)
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please type in something:\n")
	for input.Scan() {
		line := input.Text()
		if line == "result" {
			break
		}
		counts1[line]++
	}

	for line, n := range counts1 {
		a, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Println("input ok!")
			counts2[n] = a
		}
		fmt.Printf("%v\n", time.Unix(a, 0))
	}
}
