package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Scan()
	firstLine := strings.Fields(in.Text())
	n, _ := strconv.Atoi(firstLine[0])
	x, _ := strconv.Atoi(firstLine[1])

	in.Scan()
	secondLine := strings.Fields(in.Text())
	for i := 0; i < n; i++ {
		k, _ := strconv.Atoi(secondLine[i])
		if k < x {
			out.WriteString(strconv.Itoa(k) + " ")
		}
	}
	out.Flush()
}
