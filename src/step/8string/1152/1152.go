package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanBytes)

	count := 0
	var prevChar byte
	prevChar = ' '
	for in.Scan() {
		str := in.Text()
		char := str[0]

		if prevChar == ' ' {
			if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
				count++
			}
		}
		prevChar = char
	}

	out.WriteString(strconv.Itoa(count) + "\n")
	out.Flush()
}
