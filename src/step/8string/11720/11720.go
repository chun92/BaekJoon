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

	in.Scan()
	buf := in.Text()

	sum := 0
	for _, char := range buf {
		n := int(char - '0')
		sum += n
	}

	out.WriteString(strconv.Itoa(sum) + "\n")
	out.Flush()
}
