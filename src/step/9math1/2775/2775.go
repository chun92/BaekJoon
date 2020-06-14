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
		n, _ := strconv.Atoi(in.Text())
		in.Scan()
		k, _ := strconv.Atoi(in.Text())

		n = n + 1
		k = k - 1

		res := combination(n+k, k)
		out.WriteString(strconv.Itoa(res) + "\n")
	}
	out.Flush()
}

func combination(n, r int) int {
	if n == r || r == 0 {
		return 1
	} else {
		return combination(n-1, r-1) + combination(n-1, r)
	}
}
