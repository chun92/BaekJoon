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
	cnt, _ := strconv.Atoi(in.Text())

	for i := 0; i < cnt; i++ {
		in.Scan()
		n, _ := strconv.Atoi(in.Text())

		for j := n / 2; j > 0; j-- {
			if numArray[j] && numArray[n-j] {
				out.WriteString(strconv.Itoa(j) + " " + strconv.Itoa(n-j) + "\n")
				break
			}
		}
	}

	out.Flush()
}
