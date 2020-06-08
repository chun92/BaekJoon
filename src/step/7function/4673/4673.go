package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var nums [10001]int

	for i := 1; i <= 10000; i++ {
		nums[i] = 0
	}

	for i := 1; i <= 10000; i++ {
		d := getNextD(i)

		if d <= 10000 {
			nums[d] += 1
		}
	}

	out := bufio.NewWriter(os.Stdout)

	for i := 1; i <= 10000; i++ {
		d := nums[i]
		if d == 0 {
			out.WriteString(strconv.Itoa(i) + "\n")
		}
	}
	out.Flush()
}

func getNextD(n int) int {
	sum := n
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
