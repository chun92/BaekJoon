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

	var res int
	if n == 4 || n == 7 {
		res = -1
	} else if n%5 == 0 {
		res = n / 5
	} else if n%5 == 1 {
		res = n/5 + 1
	} else if n%5 == 2 {
		res = n/5 + 2
	} else if n%5 == 3 {
		res = n/5 + 1
	} else if n%5 == 4 {
		res = n/5 + 2
	}

	out.WriteString(strconv.Itoa(res) + "\n")
	out.Flush()
}
