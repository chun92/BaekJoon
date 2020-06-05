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

	if n == 1 {
		out.WriteString("*\n")
		out.Flush()
		return
	}

	for line := 0; line < n*2; line++ {
		if line%2 == 1 {
			for col := 0; col < n; col++ {
				if col%2 == 0 {
					out.WriteString(" ")
				} else {
					out.WriteString("*")
				}
			}
		} else {
			for col := 0; col < n; col++ {
				if col%2 == 1 {
					out.WriteString(" ")
				} else {
					out.WriteString("*")
				}
			}
		}
		out.WriteString("\n")
	}
	out.Flush()
}
