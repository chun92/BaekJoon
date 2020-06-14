package main

import (
	"bufio"
	"os"
	"strconv"
)

const MAX = 246912

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var numArray [MAX + 1]bool

	for i := 1; i < MAX+1; i++ {
		numArray[i] = true
	}

	numArray[0] = false
	numArray[1] = false

	for i := 2; i <= 496; i++ {
		if numArray[i] {
			for j := i * 2; j <= MAX; j += i {
				numArray[j] = false
			}
		}
	}

	in.Split(bufio.ScanWords)
	for {
		in.Scan()
		n, _ := strconv.Atoi(in.Text())
		if n == 0 {
			break
		}

		count := 0
		for i := n + 1; i <= 2*n; i++ {
			if numArray[i] {
				count++
			}
		}
		out.WriteString(strconv.Itoa(count) + "\n")
	}

	out.Flush()
}
