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

	var numArray [1001]bool

	for i := 1; i < 1001; i++ {
		numArray[i] = true
	}

	numArray[0] = false
	numArray[1] = false

	for i := 2; i < 32; i++ {
		if numArray[i] {
			for j := i * 2; j <= 1000; j += i {
				numArray[j] = false
			}
		}
	}

	count := 0
	for i := 0; i < c; i++ {
		in.Scan()

		n, _ := strconv.Atoi(in.Text())
		if numArray[n] {
			count++
		}

	}
	out.WriteString(strconv.Itoa(count) + "\n")
	out.Flush()
}
