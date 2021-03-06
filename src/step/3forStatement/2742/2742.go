package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	for i := n; i > 0; i-- {
		out.WriteString(strconv.Itoa(i) + "\n")
	}
	out.Flush()
}
