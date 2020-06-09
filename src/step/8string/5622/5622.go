package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanRunes)

	sum := 0
	for in.Scan() {
		str := in.Text()
		char := str[0]
		sum += getNum(char)
	}

	out.WriteString(strconv.Itoa(sum))
	out.Flush()
}

func getNum(a byte) int {
	if a == 'A' || a == 'B' || a == 'C' {
		return 3
	} else if a == 'D' || a == 'E' || a == 'F' {
		return 4
	} else if a == 'G' || a == 'H' || a == 'I' {
		return 5
	} else if a == 'J' || a == 'K' || a == 'L' {
		return 6
	} else if a == 'M' || a == 'N' || a == 'O' {
		return 7
	} else if a == 'P' || a == 'Q' || a == 'R' || a == 'S' {
		return 8
	} else if a == 'T' || a == 'U' || a == 'V' {
		return 9
	} else if a == 'W' || a == 'X' || a == 'Y' || a == 'Z' {
		return 10
	} else {
		return 0
	}
}
