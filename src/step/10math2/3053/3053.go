package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanWords)
	in.Scan()
	r, _ := strconv.Atoi(in.Text())

	euclid := math.Pi * math.Pow(float64(r), 2)
	nonEuclid := 2 * math.Pow(float64(r), 2)

	out.WriteString(fmt.Sprintf("%f", euclid) + "\n")
	out.WriteString(fmt.Sprintf("%f", nonEuclid) + "\n")
	out.Flush()
}
