package main

import (
	"bufio"
	"os"
	"strconv"
)

const MAX = 10000

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var numArray [MAX + 1]bool

	for i := 1; i < MAX+1; i++ {
		numArray[i] = true
	}

	numArray[0] = false
	numArray[1] = false

	for i := 2; i <= 100; i++ {
		if numArray[i] {
			for j := i * 2; j <= MAX; j += i {
				numArray[j] = false
			}
		}
	}

	in.Split(bufio.ScanWords)
	in.Scan()
	m, _ := strconv.Atoi(in.Text())
	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	sum := 0
	min := 0
	for i := m; i <= n; i++ {
		if numArray[i] {
			if min == 0 {
				min = i
			}
			sum += i
		}
	}
	if min > 0 {
		out.WriteString(strconv.Itoa(sum) + "\n")
		out.WriteString(strconv.Itoa(min) + "\n")
	} else {
		out.WriteString("-1\n")
	}
	out.Flush()
}
