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

	count := 0
	num := 666

	for {
		if checkNum(num) {
			count++
			if count == n {
				break
			}
		}
		num++
	}

	out.WriteString(strconv.Itoa(num) + "\n")
	out.Flush()
}

func checkNum(num int) bool {
	temp := num
	count := 0

	for temp > 0 {
		if temp%10 == 6 {
			count++
			if count == 3 {
				return true
			}
		} else {
			count = 0
		}
		temp = temp / 10
	}
	return false
}
