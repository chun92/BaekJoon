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
	array := make([]bool, n*n)
	for i := 0; i < n*n; i++ {
		array[i] = false
	}
	printStar(n, n, 0, 0, array)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if array[j*n+i] {
				out.WriteRune('*')
			} else {
				out.WriteRune(' ')
			}
		}
		out.WriteRune('\n')
	}
	out.Flush()
}

func printStar(n int, fixed int, x int, y int, array []bool) {
	if n == 1 {
		array[y*fixed+x] = true
	} else {
		unit := n / 3
		printStar(unit, fixed, x, y, array)
		printStar(unit, fixed, x+unit, y, array)
		printStar(unit, fixed, x+unit*2, y, array)
		printStar(unit, fixed, x, y+unit, array)
		printStar(unit, fixed, x+unit*2, y+unit, array)
		printStar(unit, fixed, x, y+unit*2, array)
		printStar(unit, fixed, x+unit, y+unit*2, array)
		printStar(unit, fixed, x+unit*2, y+unit*2, array)
	}
}
