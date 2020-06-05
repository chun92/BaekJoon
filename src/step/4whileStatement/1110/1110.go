package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	count := 1
	for i := getNext(n); i != n; count++ {
		i = getNext(i)
	}
	out.WriteString(strconv.Itoa(count) + "\n")
	out.Flush()
}

func getNext(n int) int {
	first := 0
	second := 0

	if n < 10 {
		second = n
	} else {
		first = n / 10
		second = n % 10
	}

	newNum := first + second
	return second*10 + newNum%10
}
