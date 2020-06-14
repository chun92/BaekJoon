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
	c, _ := strconv.Atoi(in.Text())

	var n int

	if b >= c {
		n = -1
	} else {
		n = a/(c-b) + 1
	}

	out.WriteString(strconv.Itoa(n) + "\n")
	out.Flush()
}
