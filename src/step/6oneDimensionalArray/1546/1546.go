package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanWords)
	in.Scan()
	num, _ := strconv.Atoi(in.Text())

	max := 0
	sum := 0
	for i := 0; i < num; i++ {
		in.Scan()
		n, _ := strconv.Atoi(in.Text())
		sum += n
		if n > max {
			max = n
		}
	}

	res := float64(sum) / float64(max) / float64(num) * 100
	out.WriteString(fmt.Sprintf("%.2f", res))
	out.Flush()
}
