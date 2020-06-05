package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	out := bufio.NewWriter(os.Stdout)
	n, _ := strconv.Atoi(in.Text())

	for i := 0; i < n; i++ {
		in.Scan()
		sum := 0
		for _, f := range strings.Fields(in.Text()) {
			num, _ := strconv.Atoi(f)
			sum += num
		}
		out.WriteString(strconv.Itoa(sum) + "\n")
	}
	out.Flush()
}
