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
	n, _ := strconv.Atoi(in.Text())

	for i := 1; i <= n; i++ {
		in.Scan()
		nums := strings.Fields(in.Text())
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		sum := a + b

		out.WriteString("Case #" + strconv.Itoa(i) + ": " +
			strconv.Itoa(a) + " + " +
			strconv.Itoa(b) + " = " +
			strconv.Itoa(sum) + "\n")
	}
	out.Flush()
}
