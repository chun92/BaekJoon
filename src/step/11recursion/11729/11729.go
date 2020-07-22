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
	out.WriteString(strconv.Itoa(numHanoi(n)) + "\n")
	moveHanoi(n, 1, 3, out)
	out.Flush()
}

func numHanoi(n int) int {
	if n == 1 {
		return 1
	} else {
		return 1 + numHanoi(n-1)*2
	}
}

func moveHanoi(n int, from int, to int, out *bufio.Writer) {
	if n == 1 {
		out.WriteString(strconv.Itoa(from) + " " + strconv.Itoa(to) + "\n")
	} else {
		other := 6 - from - to
		moveHanoi(n-1, from, other, out)
		out.WriteString(strconv.Itoa(from) + " " + strconv.Itoa(to) + "\n")
		moveHanoi(n-1, other, to, out)
	}
}
