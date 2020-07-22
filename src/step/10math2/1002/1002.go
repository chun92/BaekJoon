package main

import (
	"bufio"
	"math"
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
		x1, _ := strconv.Atoi(in.Text())
		in.Scan()
		y1, _ := strconv.Atoi(in.Text())
		in.Scan()
		r1, _ := strconv.Atoi(in.Text())
		in.Scan()
		x2, _ := strconv.Atoi(in.Text())
		in.Scan()
		y2, _ := strconv.Atoi(in.Text())
		in.Scan()
		r2, _ := strconv.Atoi(in.Text())

		dist := math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
		sum := float64(r1 + r2)
		sub := math.Abs(float64(r1 - r2))

		if dist == 0 {
			if r1 == r2 {
				out.WriteString("-1\n")
			} else {
				out.WriteString("0\n")
			}
		} else if sum < dist {
			out.WriteString("0\n")
		} else if sum == dist {
			out.WriteString("1\n")
		} else {
			if sub < dist {
				out.WriteString("2\n")
			} else if sub == dist {
				out.WriteString("1\n")
			} else {
				out.WriteString("0\n")
			}
		}
	}
	out.Flush()
}
