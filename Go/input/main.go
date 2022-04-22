package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Scan()
		// 从stdin中取内容直到遇到换行符，停止
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}

		input = strings.Replace(input, string('\n'), " ", -1)

		numbers := strings.Split(input, " ")
		a, err := strconv.ParseInt(numbers[0], 10, 64)
		if err != nil {
			panic(err)
		}

		b, err := strconv.ParseInt(numbers[1], 10, 64)
		if err != nil {
			panic(err)
		}

		fmt.Println(a + b)
	}
}