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

	res := 1
	path := 1
	for path = 1; res < n; {
		path++
		res = getHexnum(path)
	}

	out.WriteString(strconv.Itoa(path) + "\n")
	out.Flush()
}

func getHexnum(n int) int {
	return 3*n*(n-1) + 1
}
