package main

import (
	"bufio"
	"os"
	"strconv"
)

const NumOfLine = 5

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	sum := 0
	for i := 0; i < NumOfLine; i++ {
		in.Scan()
		score, _ := strconv.Atoi(in.Text())
		if score < 40 {
			score = 40
		}
		sum += score
	}

	out.WriteString(strconv.Itoa(sum/NumOfLine) + "\n")
	out.Flush()
}
