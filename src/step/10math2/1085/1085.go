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
	x, _ := strconv.Atoi(in.Text())
	in.Scan()
	y, _ := strconv.Atoi(in.Text())
	in.Scan()
	w, _ := strconv.Atoi(in.Text())
	in.Scan()
	h, _ := strconv.Atoi(in.Text())

	var a = [4]int{x, y, w - x, h - y}

	min := x

	for _, value := range a {
		if value < min {
			min = value
		}
	}

	out.WriteString(strconv.Itoa(min) + "\n")
	out.Flush()
}
