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
	a, _ := strconv.Atoi(in.Text())
	in.Scan()
	b, _ := strconv.Atoi(in.Text())
	in.Scan()
	v, _ := strconv.Atoi(in.Text())

	fullDay := a - b
	lastDistance := v - a
	var carry int
	if lastDistance%fullDay == 0 {
		carry = 0
	} else {
		carry = 1
	}
	count := lastDistance/fullDay + carry + 1

	out.WriteString(strconv.Itoa(count) + "\n")
	out.Flush()
}
