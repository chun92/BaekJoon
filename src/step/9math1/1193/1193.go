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

	line := 1
	ran := 1
	for n > ran {
		line++
		ran = getRangeMax(line)
	}

	index := n - getRangeMax(line-1)

	var a, b int
	if line%2 == 0 {
		a = index
		b = line - index + 1
	} else {
		a = line - index + 1
		b = index
	}

	out.WriteString(strconv.Itoa(a) + "/" + strconv.Itoa(b) + "\n")
	out.Flush()
}

func getRangeMax(n int) int {
	return n * (n + 1) / 2
}
