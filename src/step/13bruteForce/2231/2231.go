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

	res := 0

	for i := getLeastRange(n); i < n; i++ {
		if getRes(i) == n {
			res = i
			break
		}
	}

	out.WriteString(strconv.Itoa(res) + "\n")
	out.Flush()
}

func getLeastRange(n int) int {
	temp := n
	sum := 0
	for temp > 0 {
		sum = sum + 9
		temp = temp / 10
	}
	return n - sum
}

func getRes(n int) int {
	temp := n
	sum := n
	for temp > 0 {
		sum = sum + temp%10
		temp = temp / 10
	}
	return sum
}
