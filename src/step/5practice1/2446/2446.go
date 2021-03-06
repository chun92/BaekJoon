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

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			out.WriteString(" ")
		}
		for j := 2*(n-i) - 1; j > 0; j-- {
			out.WriteString("*")
		}
		out.WriteString("\n")
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < i; j++ {
			out.WriteString(" ")
		}
		for j := 2*(n-i) - 1; j > 0; j-- {
			out.WriteString("*")
		}
		out.WriteString("\n")
	}
	out.Flush()
}
