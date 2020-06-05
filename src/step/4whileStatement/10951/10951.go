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

	for in.Scan() {
		line := in.Text()
		if len(line) == 0 {
			break
		}
		nums := strings.Fields(line)
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		out.WriteString(strconv.Itoa(a+b) + "\n")
	}
	out.Flush()
}
