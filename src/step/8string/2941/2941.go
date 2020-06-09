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
	var prev1Char byte
	var prev2Char byte
	prev1Char = ' '
	prev2Char = ' '
	for in.Scan() {
		str := in.Text()
		char := str[0]

		cond := false
		if 'a' <= char && char <= 'z' {
			cond = true
		} else if char == '-' || char == '=' {
			cond = true
		}

		if !cond {
			continue
		}

		if prev2Char == 'z' {
			if char == '=' {
				count++
				prev1Char = ' '
				prev2Char = ' '
				continue
			} else {
				count += 2
				prev1Char = ' '
				prev2Char = ' '
			}
		}

		if prev1Char != ' ' {
			if prev1Char == 'c' {
				if char == '=' || char == '-' {
					count++
					prev1Char = ' '
					continue
				} else {
					count++
					prev1Char = ' '
				}
			} else if prev1Char == 'd' {
				if char == '-' {
					count++
					prev1Char = ' '
					continue
				} else if char == 'z' {
					prev2Char = char
					continue
				} else {
					count++
					prev1Char = ' '
				}
			} else if prev1Char == 'l' || prev1Char == 'n' {
				if char == 'j' {
					count++
					prev1Char = ' '
					continue
				} else {
					count++
					prev1Char = ' '
				}
			} else if prev1Char == 's' || prev1Char == 'z' {
				if char == '=' {
					count++
					prev1Char = ' '
					continue
				} else {
					count++
					prev1Char = ' '
				}
			}
		}

		if char == 'c' || char == 'd' || char == 'l' || char == 'n' || char == 's' || char == 'z' {
			prev1Char = char
		} else {
			count++
		}
	}

	if prev1Char != ' ' {
		count++
	}

	if prev2Char != ' ' {
		count++
	}

	out.WriteString(strconv.Itoa(count) + "\n")
	out.Flush()
}
