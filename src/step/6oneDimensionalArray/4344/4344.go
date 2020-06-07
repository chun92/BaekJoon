package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Scan()
	cases, _ := strconv.Atoi(in.Text())

	for i := 0; i < cases; i++ {
		printPercentageOverAverage(in, out)
	}
	out.Flush()
}

func printPercentageOverAverage(in *bufio.Scanner, out *bufio.Writer) {
	in.Scan()
	nums := strings.Fields(in.Text())
	count := 0
	sum := 0
	conditionalCount := 0
	for i, n := range nums {
		num, _ := strconv.Atoi(n)
		if i < 1 {
			count = num
		} else {
			sum += num
		}
	}

	average := float64(sum) / float64(count)

	for i, n := range nums {
		num, _ := strconv.Atoi(n)
		if i < 1 {
			continue
		} else {
			if float64(num) > average {
				conditionalCount++
			}
		}
	}

	res := float64(conditionalCount) / float64(count) * 100

	out.WriteString(fmt.Sprintf("%.3f", res) + "%\n")
}
