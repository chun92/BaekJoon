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
	res := factorial(n, 1)
	out.WriteString(strconv.Itoa(res) + "\n")
	out.Flush()
}

func factorial(n int, fac int) int {
	if n == 0 {
		return fac
	} else {
		return factorial(n-1, fac*n)
	}
}
