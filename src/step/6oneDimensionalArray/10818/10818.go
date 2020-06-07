package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanWords)

	in.Scan()

	min := 1000000
	max := -1000000

	for in.Scan() {
		num, _ := strconv.Atoi(in.Text())
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	out.WriteString(strconv.Itoa(min) + " " + strconv.Itoa(max) + "\n")
	out.Flush()
}
