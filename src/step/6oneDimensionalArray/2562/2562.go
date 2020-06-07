package main

import (
	"bufio"
	"os"
	"strconv"
)

const TotalCount = 9

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	count := 0
	max := 0

	for i := 1; i <= TotalCount; i++ {
		in.Scan()
		num, _ := strconv.Atoi(in.Text())

		if num > max {
			count = i
			max = num
		}
	}

	out.WriteString(strconv.Itoa(max) + "\n")
	out.WriteString(strconv.Itoa(count) + "\n")
	out.Flush()
}
