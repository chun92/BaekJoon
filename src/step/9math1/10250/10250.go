package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanWords)
	in.Scan()
	c, _ := strconv.Atoi(in.Text())

	for i := 0; i < c; i++ {
		in.Scan()
		h, _ := strconv.Atoi(in.Text())
		in.Scan()
		in.Scan()
		n, _ := strconv.Atoi(in.Text())

		x := n/h + 1
		y := n % h
		if y == 0 {
			y = h
			x -= 1
		}
		target := y*100 + x
		out.WriteString(strconv.Itoa(target) + "\n")
	}
	out.Flush()
}
