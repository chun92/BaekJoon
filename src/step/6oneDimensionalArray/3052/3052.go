package main

import (
	"bufio"
	"os"
	"strconv"
)

const Inputs = 10
const Divisor = 42

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var remainders [Divisor]int

	for i := 0; i < Divisor; i++ {
		remainders[i] = 0
	}

	for i := 0; i < Inputs; i++ {
		in.Scan()
		num, _ := strconv.Atoi(in.Text())
		remainders[num%Divisor] = 1
	}

	count := 0
	for i := 0; i < Divisor; i++ {
		if remainders[i] > 0 {
			count++
		}
	}

	out.WriteString(strconv.Itoa(count) + "\n")
	out.Flush()
}
