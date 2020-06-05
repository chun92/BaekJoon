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

	for {
		in.Scan()
		nums := strings.Fields(in.Text())
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		if a == 0 && b == 0 {
			break
		}

		out.WriteString(strconv.Itoa(a+b) + "\n")
	}
	out.Flush()
}
