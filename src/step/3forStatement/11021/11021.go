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
		sum := 0
		for _, f := range strings.Fields(in.Text()) {
			num, _ := strconv.Atoi(f)
			sum += num
		}
		out.WriteString("Case #" + strconv.Itoa(i) + ": " + strconv.Itoa(sum) + "\n")
	}
	out.Flush()
}
