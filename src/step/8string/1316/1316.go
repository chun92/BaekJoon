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
	n, _ := strconv.Atoi(in.Text())

	count := 0
	for i := 0; i < n; i++ {
		in.Scan()
		if isGroupWord(in.Text()) {
			count++
		}
	}

	out.WriteString(strconv.Itoa(count))
	out.Flush()
}

func isGroupWord(a string) bool {
	var index [26]int
	for i := 0; i < 26; i++ {
		index[i] = 0
	}

	prevNum := -1
	for _, c := range a {
		num := int(c - 'a')
		if index[num] != 0 && prevNum != num {
			return false
		}

		index[num] = 1
		prevNum = num
	}
	return true
}
