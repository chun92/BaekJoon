package main

import (
	"bufio"
	"os"
	"strconv"
)

const Numbers = 10

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Scan()
	a, _ := strconv.Atoi(in.Text())
	in.Scan()
	b, _ := strconv.Atoi(in.Text())
	in.Scan()
	c, _ := strconv.Atoi(in.Text())
	res := a * b * c

	var numbers [Numbers]int
	for i := 0; i < Numbers; i++ {
		numbers[i] = 0
	}

	for res > 0 {
		num := res % 10
		numbers[num] += 1
		res = res / 10
	}

	for i := 0; i < Numbers; i++ {
		out.WriteString(strconv.Itoa(numbers[i]) + "\n")
	}
	out.Flush()
}
