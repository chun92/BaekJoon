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
	in.Scan()
	m, _ := strconv.Atoi(in.Text())

	array := make([][]bool, n)
	for i := 0; i < n; i++ {
		array[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		in.Scan()
		for j := 0; j < m; j++ {
			target := in.Text()[j]
			if target == 'W' {
				array[i][j] = true
			} else {
				array[i][j] = false
			}
		}
	}

	res := 64
	for i := 0; i <= n-8; i++ {
		for j := 0; j <= m-8; j++ {
			temp := checkNum(i, j, array)
			if temp < res {
				res = temp
			}
		}
	}
	out.WriteString(strconv.Itoa(res) + "\n")
	out.Flush()
}

func checkNum(x int, y int, array [][]bool) int {
	countA := 0
	countB := 0
	target := true

	for i := x; i < x+8; i++ {
		for j := y; j < y+8; j++ {
			if target != array[i][j] {
				countA++
			} else {
				countB++
			}
			target = !target
		}
		target = !target
	}

	if countA > countB {
		return countB
	} else {
		return countA
	}
}
