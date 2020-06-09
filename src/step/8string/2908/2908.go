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

	reverseA := reverse(a)
	reverseB := reverse(b)

	if reverseA > reverseB {
		out.WriteString(strconv.Itoa(reverseA))
	} else {
		out.WriteString(strconv.Itoa(reverseB))
	}
	out.Flush()
}

func reverse(n int) int {
	res := 0
	for n > 0 {
		res = res*10 + (n % 10)
		n = n / 10
	}
	return res
}
