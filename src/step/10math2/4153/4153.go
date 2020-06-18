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
	for {
		in.Scan()
		a, _ := strconv.Atoi(in.Text())
		in.Scan()
		b, _ := strconv.Atoi(in.Text())
		in.Scan()
		c, _ := strconv.Atoi(in.Text())

		if a == 0 && b == 0 && c == 0 {
			break
		}

		var x, y, z int

		if a <= b && b <= c {
			x = a
			y = b
			z = c
		} else if a <= c && c <= b {
			x = a
			y = c
			z = b
		} else if b <= a && a <= c {
			x = b
			y = a
			z = c
		} else if b <= c && c <= a {
			x = b
			y = c
			z = a
		} else if c <= a && a <= b {
			x = c
			y = a
			z = b
		} else if c <= b && b <= a {
			x = c
			y = b
			z = a
		}

		if x*x+y*y-z*z == 0 {
			out.WriteString("right\n")
		} else {
			out.WriteString("wrong\n")
		}
	}

	out.Flush()
}
