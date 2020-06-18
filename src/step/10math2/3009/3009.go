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
	x1, _ := strconv.Atoi(in.Text())
	in.Scan()
	y1, _ := strconv.Atoi(in.Text())
	in.Scan()
	x2, _ := strconv.Atoi(in.Text())
	in.Scan()
	y2, _ := strconv.Atoi(in.Text())
	in.Scan()
	x3, _ := strconv.Atoi(in.Text())
	in.Scan()
	y3, _ := strconv.Atoi(in.Text())

	var x4, y4 int

	if x1 == x2 {
		x4 = x3
	} else if x1 == x3 {
		x4 = x2
	} else {
		x4 = x1
	}

	if y1 == y2 {
		y4 = y3
	} else if y1 == y3 {
		y4 = y2
	} else {
		y4 = y1
	}

	out.WriteString(strconv.Itoa(x4) + " " + strconv.Itoa(y4) + "\n")
	out.Flush()
}
