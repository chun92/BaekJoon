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
	n, _ := strconv.Atoi(in.Text())
	xArray := make([]int, n)
	yArray := make([]int, n)
	rankArray := make([]int, n)

	for i := 0; i < n; i++ {
		in.Scan()
		x, _ := strconv.Atoi(in.Text())
		in.Scan()
		y, _ := strconv.Atoi(in.Text())

		xArray[i] = x
		yArray[i] = y
		rankArray[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if xArray[i] > xArray[j] && yArray[i] > yArray[j] {
				rankArray[j] += 1
			} else if xArray[i] < xArray[j] && yArray[i] < yArray[j] {
				rankArray[i] += 1
			}
		}
	}

	for i := 0; i < n; i++ {
		out.WriteString(strconv.Itoa(rankArray[i]) + " ")
	}
	out.WriteString("\n")
	out.Flush()
}
