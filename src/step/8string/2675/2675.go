package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanWords)
	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	for i := 0; i < n; i++ {
		in.Scan()
		numOfRepeat, _ := strconv.Atoi(in.Text())
		in.Scan()
		buf := in.Text()

		for _, char := range buf {
			for repeat := 0; repeat < numOfRepeat; repeat++ {
				_, err := out.WriteRune(char)
				if err != nil {
					out.Flush()
					out.WriteRune(char)
				}
			}
		}
		out.WriteString("\n")
	}
	out.Flush()
}
