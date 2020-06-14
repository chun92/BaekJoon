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

	for i := 0; i < c; i++ {
		in.Scan()
		a, _ := strconv.Atoi(in.Text())
		in.Scan()
		b, _ := strconv.Atoi(in.Text())

		n := b - a

		count := 1
		max := 1
		for n > max {
			max += count/2 + 1
			count++
		}
		out.WriteString(strconv.Itoa(count) + "\n")
	}
	out.Flush()
}
