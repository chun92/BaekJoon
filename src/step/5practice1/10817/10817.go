package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Scan()
	nums := strings.Fields(in.Text())

	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])
	c, _ := strconv.Atoi(nums[2])

	if a > b {
		swap(&a, &b)
	}
	if a > c {
		swap(&a, &c)
	}
	if b > c {
		swap(&b, &c)
	}
	out.WriteString(strconv.Itoa(b) + "\n")
	out.Flush()
}

func swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
