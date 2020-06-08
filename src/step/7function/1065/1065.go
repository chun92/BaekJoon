package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	var res int
	if n < 100 {
		res = n
	} else {
		res = 99
		for i := 100; i <= n; i++ {
			if getOneNum(i) {
				res++
			}
		}
	}

	out.WriteString(strconv.Itoa(res) + "\n")
	out.Flush()
}

func getOneNum(n int) bool {
	dist := 10
	prev := n % 10
	n = n / 10
	for n > 0 {
		cur := n % 10
		curDist := cur - prev
		if dist == 10 {
			dist = curDist
		} else {
			if dist != curDist {
				return false
			}
		}
		prev = cur
		dist = curDist
		n = n / 10
	}

	return true
}
