package main

import (
	"bufio"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanRunes)

	var chars [26]int

	for i := 0; i < 26; i++ {
		chars[i] = 0
	}

	for in.Scan() {
		str := in.Text()
		char := str[0]
		index := int(char)
		if index >= 97 {
			index = index - 97
		} else {
			index = index - 65
		}

		if index >= 0 && index < 26 {
			chars[index]++
		}
	}

	maxCount := 0
	var maxChar rune

	for i := 0; i < 26; i++ {
		if chars[i] > maxCount {
			maxCount = chars[i]
			maxChar = rune(int('A') + i)
		} else if chars[i] == maxCount {
			maxChar = '?'
		}
	}

	out.WriteRune(maxChar)
	out.Flush()
}
