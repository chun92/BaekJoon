package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var index [26]int

	for i := 0; i < 26; i++ {
		index[i] = -1
	}

	in.Scan()
	buf := in.Text()

	for i, char := range buf {
		n := int(char - 'a')
		if index[n] == -1 {
			index[n] = i
		}
	}

	for i := 0; i < 26; i++ {
		out.WriteString(strconv.Itoa(index[i]) + " ")
	}
	out.WriteString("\n")
	out.Flush()
}
