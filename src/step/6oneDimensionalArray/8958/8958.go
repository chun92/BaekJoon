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

	for i := 0; i < n; i++ {
		in.Scan()
		printScore(in.Text(), out)
	}
	out.Flush()
}

func printScore(str string, out *bufio.Writer) {
	NumOfConsecutiveO := 0
	score := 0
	for _, char := range str {
		if char == 'O' {
			NumOfConsecutiveO++
			score += NumOfConsecutiveO
		} else {
			NumOfConsecutiveO = 0
		}
	}
	out.WriteString(strconv.Itoa(score) + "\n")
}
